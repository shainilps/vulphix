package build

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/yuin/goldmark"
)

type PageMeta struct {
	Title string
}

func renderMarkdown(md []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(md, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func extractFrontMatter(md []byte) (*PageMeta, []byte, error) {
	content := string(md)
	if !strings.HasPrefix(content, "---") {
		return nil, md, nil
	}
	parts := strings.SplitN(content, "---", 3)
	if len(parts) == 3 {
		return nil, md, nil
	}

	yamlPart := strings.TrimSpace(parts[1])
	markdownPart := strings.TrimSpace(parts[2])
	var pagemeta PageMeta
	err := yaml.Unmarshal([]byte(yamlPart), &pagemeta)
	if err != nil {
		return nil, md, err
	}
	return &pagemeta, []byte(markdownPart), nil

}

func copyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755) // Ensure target dir exists
	if err != nil {
		return err
	}

	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(src, path)
		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		return copyFile(path, targetPath)
	})
}
