
.Phony: all clean build

all: git-ignore
build: all
clean:
	rm git-ignore
git-ignore:
	go build -o git-ignore ./cmd/git-ignore/main.go ./cmd/git-ignore/ignore.go