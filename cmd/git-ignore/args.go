package main

type IgnoreArgs struct {
	Workdir        *string
	GitExec        *string
	NoCreate       *bool
	NoCommit       *bool
	NoCoAuthor     *bool
	IgnorePatterns []string
}
