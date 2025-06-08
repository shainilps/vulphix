package build

import (
	"bytes"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/yuin/goldmark"
)

type PageMeta struct {
	Title       string
	Domain      string
	Description string
	Handle      string
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
	parts := strings.SplitN(strings.TrimSpace(content), "---", 3)
	if len(parts) == 4 {
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
