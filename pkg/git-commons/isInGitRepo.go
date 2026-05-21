package git_commons

import (
	"context"
	"os"
	"path/filepath"
	"sync"
)

func IsGitRepo(basePath string) bool {
	dirs := parentDirs(basePath)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	checkChan := make(chan bool, 1)
	var wg sync.WaitGroup

	for _, dir := range dirs {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			isAGitRepo(ctx, checkChan, d, cancel)
		}(dir)
	}

	go func() {
		wg.Wait()
		close(checkChan)
	}()

	select {
	case <-checkChan:
		return true
	case <-ctx.Done():
		return false
	}
}

func parentDirs(path string) []string {
	var parents []string

	for {
		parent := filepath.Dir(path)

		// stop when we can’t go higher
		if parent == path {
			break
		}

		parents = append(parents, parent)
		path = parent
	}

	return parents
}

func isAGitRepo(ctx context.Context, resultCh chan<- bool, basePath string, cancel context.CancelFunc) {
	if _, err := os.Stat(filepath.Join(basePath, ".git")); err == nil {
		select {
		case <-ctx.Done():
			return
		case resultCh <- true:
			cancel()
		}
	}
}
