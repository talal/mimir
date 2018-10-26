ifeq ($(shell uname -s),Darwin)
	PREFIX  := /usr/local
else
	PREFIX  := /usr
endif
PKG = github.com/talal/mimir

GO          := GOPATH=$(CURDIR)/.gopath GOBIN=$(CURDIR)/build go
BUILD_FLAGS :=
LD_FLAGS    := -s -w

ifndef GOOS
GOOS := $(word 1, $(subst /, " ", $(word 4, $(shell go version))))
endif

BINARY64  := mimir-$(GOOS)_amd64
RELEASE64 := mimir-$(GOOS)_amd64

################################################################################

all: build/mimir

# This target uses the incremental rebuild capabilities of the Go compiler to speed things up.
# If no source files have changed, `go install` exits quickly without doing anything.
build/mimir: FORCE
	$(GO) install $(BUILD_FLAGS) -ldflags '$(LD_FLAGS)' '$(PKG)'

install: FORCE all
	install -d -m 0755 "$(DESTDIR)$(PREFIX)/bin"
	install -m 0755 build/mimir "$(DESTDIR)$(PREFIX)/bin/mimir"

release: release/$(BINARY64)
	cd release && cp -f $(BINARY64) mimir && tar -czf $(RELEASE64).tar.gz mimir
	cd release && rm -f mimir

release/$(BINARY64): FORCE
	GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $@ -ldflags '$(LD_FLAGS)' '$(PKG)'

clean: FORCE
	rm -rf build release

.PHONY: FORCE
