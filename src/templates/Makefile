PROJECT = {{PROJECT_SLUG}}
MAIN    = src/$(PROJECT).tex
BUILD   = build

.PHONY: all pdf clean tags watch

all: pdf

pdf: tags
	latexmk $(MAIN)

watch:
	latexmk -pvc $(MAIN)

tags:
	@if command -v ctags >/dev/null 2>&1; then \
		ctags -R src/ ; \
		echo "[+] ctags updated" ; \
	fi

clean:
	latexmk -C $(MAIN)
	rm -rf $(BUILD) tags
