# Install prefix (override with: make PREFIX=/opt install)
PREFIX  ?= /usr/local
BINDIR  := $(PREFIX)/bin

# Project files
TARGET := git-ignore-tool

.Phony: all clean build

all: $(TARGET)
build: all
clean:
	rm $(TARGET)
$(TARGET):
	go build -o $(TARGET) ./cmd/git-ignore/main.go ./cmd/git-ignore/ignore.go ./cmd/git-ignore/unignore.go ./cmd/git-ignore/untrack-ignored.go


install: $(TARGET)
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 $(TARGET) $(DESTDIR)$(BINDIR)/
	ln -s $(DESTDIR)$(BINDIR)/$(TARGET) $(DESTDIR)$(BINDIR)/git-ignore
	ln -s $(DESTDIR)$(BINDIR)/$(TARGET) $(DESTDIR)$(BINDIR)/git-unignore
	ln -s $(DESTDIR)$(BINDIR)/$(TARGET) $(DESTDIR)$(BINDIR)/git-untrack-ignored

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/$(TARGET)
	rm -f $(DESTDIR)$(BINDIR)/git-ignore
	rm -f $(DESTDIR)$(BINDIR)/git-unignore
	rm -f $(DESTDIR)$(BINDIR)/git-untrack-ignored
