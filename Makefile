PKG    = github.com/talal/mimir
PREFIX := /usr/local

all: build/mimir

GO            := GOBIN=$(CURDIR)/build go
GO_BUILDFLAGS :=
GO_LDFLAGS    := -s -w


build/mimir: FORCE
	$(GO) install $(GO_BUILDFLAGS) -ldflags '$(GO_LDFLAGS)' '$(PKG)'

install: FORCE all
	install -d -m 0755 "$(DESTDIR)$(PREFIX)/bin"
	install -m 0755 build/mimir "$(DESTDIR)$(PREFIX)/bin/mimir"

.PHONY: FORCE
