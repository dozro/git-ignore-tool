# Install prefix (override with: make PREFIX=/opt install)
PREFIX  ?= /usr/local
BINDIR  := $(PREFIX)/bin

# Project files
TARGETGITITGNORE  := git-ignore

.Phony: all clean build

all: $(TARGETGITITGNORE)
build: all
clean:
	rm git-ignore
$(TARGETGITITGNORE):
	go build -o git-ignore ./cmd/git-ignore/main.go ./cmd/git-ignore/ignore.go

install: $(TARGET)
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 $(TARGETGITITGNORE) $(DESTDIR)$(BINDIR)/

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/$(TARGETGITITGNORE)
