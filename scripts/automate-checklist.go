package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	THIS_FILE = "automate-checklist.go"
)

func main() {
	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	if owner == "" {
		owner = getGitOwner()
	}

	repo := ""
	githubRepo := os.Getenv("GITHUB_REPOSITORY") // e.g., "owner/repo"
	if githubRepo != "" {
		parts := strings.Split(githubRepo, "/")
		if len(parts) == 2 {
			repo = parts[1]
		}
	}
	if repo == "" {
		repo = getGitRepoName()
	}

	if owner == "" || repo == "" {
		fmt.Println("Could not determine owner or repo. Exiting.")
		os.Exit(1)
	}

	fmt.Printf("---> 😎 OWNER: %s\n", owner)
	fmt.Printf("---> 🧰 REPO: %s\n\n", repo)

	lcRepo := strings.ToLower(repo)
	if len(repo) == 0 {
		fmt.Println("Repo name is empty. Exiting.")
		os.Exit(1)
	}
	tcRepo := strings.ToUpper(repo[:1]) + strings.ToLower(repo[1:])
	ucRepo := strings.ToUpper(repo)

	replacements := map[string]string{
		"snivilised/arcadia": fmt.Sprintf("%s/%s", owner, lcRepo),
		"snivilised/Arcadia": fmt.Sprintf("%s/%s", owner, tcRepo),
		"snivilised/ARCADIA": fmt.Sprintf("%s/%s", owner, ucRepo),
		"arcadia":            lcRepo,
		"Arcadia":            tcRepo,
		"ARCADIA":            ucRepo,
	}

	// 1. & 2. Global replace
	doGlobalReplace(replacements)

	// 3. Rename files
	renameTargetFiles("arcadia", lcRepo)
	renameTargetFiles("Arcadia", tcRepo)
	renameTargetFiles("ARCADIA", ucRepo)

	// 4. Reset version
	err := os.WriteFile("VERSION", []byte("v0.1.0\n"), 0644) //nolint:gosec // ok, not sensitive
	if err != nil {
		fmt.Printf("Error writing VERSION: %v\n", err)
	}

	// 5. Create .env
	f, err := os.OpenFile(".env", os.O_CREATE|os.O_RDWR, 0644) //nolint:gosec // ok, not sensitive
	if err == nil {
		_ = f.Close()
	}

	fmt.Println("✔️ done")
}

func getGitOwner() string {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	url := strings.TrimSpace(string(out))
	// Parse URL: https://github.com/owner/repo.git or git@github.com:owner/repo.git
	if strings.Contains(url, "git@") {
		parts := strings.Split(url, ":")
		if len(parts) > 1 {
			path := parts[1]
			pathParts := strings.Split(path, "/")
			if len(pathParts) > 0 {
				return pathParts[0]
			}
		}
	} else if strings.Contains(url, "https://") {
		parts := strings.Split(url, "/")
		if len(parts) >= 4 {
			return parts[3]
		}
	}
	return ""
}

func getGitRepoName() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return filepath.Base(strings.TrimSpace(string(out)))
}

func isTargetFile(path string) bool {
	if strings.Contains(path, "/.git/") || strings.Contains(path, "/vendor/") {
		return false
	}
	ext := filepath.Ext(path)
	base := filepath.Base(path)
	if base == THIS_FILE {
		return false
	}
	switch ext {
	case ".go", ".md", ".yml", ".yaml", ".json":
		return true
	}
	if base == "go.mod" || strings.HasPrefix(base, "Taskfile") {
		return true
	}
	return false
}

func doGlobalReplace(replacements map[string]string) {
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == "vendor" {
				return fs.SkipDir
			}
			return nil
		}
		if !isTargetFile(path) {
			return nil
		}

		content, err := os.ReadFile(path) //nolint:gosec // ok, not user-provided content
		if err != nil {
			return nil
		}

		modified := false
		newContent := content

		// Order of replacements matters
		order := []string{
			"snivilised/arcadia",
			"snivilised/Arcadia",
			"snivilised/ARCADIA",
			"arcadia",
			"Arcadia",
			"ARCADIA",
		}

		for _, k := range order {
			v := replacements[k]
			if bytes.Contains(newContent, []byte(k)) {
				newContent = bytes.ReplaceAll(newContent, []byte(k), []byte(v))
				modified = true
			}
		}

		if modified {
			// Don't change file permissions
			info, err := d.Info()
			if err == nil {
				_ = os.WriteFile(path, newContent, info.Mode())
			}
		}
		return nil
	})
}

func renameTargetFiles(target, replacement string) {
	var pathsToRename []string
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if d.Name() == ".git" || d.Name() == "vendor" {
				return fs.SkipDir
			}
			return nil
		}
		base := filepath.Base(path)
		if strings.Contains(base, target) {
			pathsToRename = append(pathsToRename, path)
		}
		return nil
	})

	for _, path := range pathsToRename {
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		newBase := strings.ReplaceAll(base, target, replacement)
		newPath := filepath.Join(dir, newBase)
		_ = os.Rename(path, newPath) //nolint:gosec // ok, not user-provided content
	}
}
