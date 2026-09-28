package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var embeddedTemplates embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <project-dir-name> [project-title]\n", filepath.Base(os.Args[0]))
		fmt.Printf("Example: %s cybernetics-monograph \"Cybernetics and Cognition\"\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	projectDir := os.Args[1]
	projectTitle := projectDir
	if len(os.Args) >= 3 {
		projectTitle = os.Args[2]
	}

	projectSlug := strings.ReplaceAll(projectDir, "-", "_")

	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "[-] Error: Directory '%s' already exists.\n", projectDir)
		os.Exit(1)
	}

	fmt.Printf("[*] Initializing LuaLaTeX project in: %s\n", projectDir)

	// 1. 骨格ディレクトリの作成
	scaffoldDirs := []string{
		".ctags.d",
		"build",
		"db/seed",
		"src/chapters/00_intro",
		"src/chapters/01_foundations",
		"src/chapters/99_colophon",
		"src/dot",
		"src/epigraph",
		"src/figures",
		"src/fonts",
		"src/images",
		"src/lua",
		"src/math",
		"src/old",
		"src/pikchr",
		"src/plots",
		"src/preamble",
		"src/tables",
		"src/tikz",
	}

	for _, d := range scaffoldDirs {
		target := filepath.Join(projectDir, d)
		if err := os.MkdirAll(target, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "[-] Failed to create dir %s: %v\n", target, err)
			os.Exit(1)
		}
	}

	// 2. 埋め込みファイルの展開マップ (Embed Path -> Dest Relative Path)
	fileMapping := map[string]string{
		"templates/gitignore":                ".gitignore",
		"templates/latexmkrc":                ".latexmkrc",
		"templates/ctags.d/latex.ctags":      ".ctags.d/latex.ctags",
		"templates/Makefile":                 "Makefile",
		"templates/CODING_CONVENTIONS.md":    "CODING_CONVENTIONS.md",
		"templates/preamble/packages.tex":    "src/preamble/packages.tex",
		"templates/preamble/init.tex":        "src/preamble/init.tex",
		"templates/preamble/macros.tex":      "src/preamble/macros.tex",
		"templates/preamble/math_commands.tex": "src/preamble/math_commands.tex",
		"templates/preamble/environments.tex": "src/preamble/environments.tex",
		"templates/preamble/terms.tex":       "src/preamble/terms.tex",
		"templates/chapters/intro.tex":       "src/chapters/00_intro/main.tex",
		"templates/chapters/foundations.tex": "src/chapters/01_foundations/main.tex",
		"templates/chapters/colophon.tex":    "src/chapters/99_colophon/main.tex",
		"templates/main.tex":                 fmt.Sprintf("src/%s.tex", projectSlug),
	}

	for srcPath, destRel := range fileMapping {
		raw, err := fs.ReadFile(embeddedTemplates, srcPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[-] Read template error (%s): %v\n", srcPath, err)
			os.Exit(1)
		}

		content := string(raw)
		content = strings.ReplaceAll(content, "{{PROJECT_SLUG}}", projectSlug)
		content = strings.ReplaceAll(content, "{{PROJECT_TITLE}}", projectTitle)

		destPath := filepath.Join(projectDir, destRel)
		if err := os.WriteFile(destPath, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "[-] Write error (%s): %v\n", destPath, err)
			os.Exit(1)
		}
	}

	// 3. ctags の実行（環境にあれば）
	if _, err := exec.LookPath("ctags"); err == nil {
		cmd := exec.Command("ctags", "-R", "src/")
		cmd.Dir = projectDir
		_ = cmd.Run()
	}

	fmt.Printf("[✔] Successfully initialized %s.\n", projectDir)
	fmt.Printf("    cd %s && make\n", projectDir)
}
