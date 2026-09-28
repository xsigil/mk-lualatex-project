PREFIX    ?= /usr/local
BUILD_DIR := build
TARGET    := $(BUILD_DIR)/mk-lualatex-project
SRCS      := src/main.go $(shell find src/templates -type f)

GOFLAGS   ?= -trimpath -ldflags="-s -w"

.PHONY: all clean install test

all: $(TARGET)

$(TARGET): $(SRCS) | $(BUILD_DIR)
	go build $(GOFLAGS) -o $(TARGET) src/main.go

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

install: all
	install -Dm755 $(TARGET) $(PREFIX)/bin/mk-lualatex-project

test: all
	@rm -rf /tmp/test-tex-proj
	@$(TARGET) /tmp/test-tex-proj "Integration Test Project"
	@test -f /tmp/test-tex-proj/Makefile
	@test -f /tmp/test-tex-proj/src/test_tex_proj.tex
	@rm -rf /tmp/test-tex-proj
	@echo "[+] All integration tests passed."
