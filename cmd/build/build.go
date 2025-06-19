package build

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

const (
	CONFIG_FILE    = "vulphix.config.yaml"
	DEFAULT_SOURCE = "src"
	DEFAULT_BUILD  = "dist"
)

type SidebarItem struct {
	Title string      `yaml:"title"`
	Pages [][2]string `yaml:"pages"`
}

type Config struct {
	Title       string        `yaml:"title"`
	Domain      string        `yaml:"domain"`
	Description string        `yaml:"description"`
	Handle      string        `yaml:"handle"`
	Source      string        `yaml:"source"` //optional
	Build       string        `yaml:"build"`  //optional
	Sidebar     []SidebarItem `yaml:"sidebar"`
}

type Page struct {
	Title       string
	Meta        PageMeta
	LeftSideBar []SidebarItem
	Content     template.HTML
	StyleSheets []string
}

//go:embed assets/*
var assets embed.FS

func readEmbedFiles(filename string) ([]byte, error) {
	return assets.ReadFile(filename)
}

func ReadConfig() (*Config, error) {
	byte, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(byte, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func Build() int {
	config, err := ReadConfig()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || config == nil {
			fmt.Printf("Missing '%v'\n", CONFIG_FILE)
		}
		return 1
	}

	if config.Source == "" {
		config.Source = DEFAULT_SOURCE
	}
	if config.Build == "" {
		config.Build = DEFAULT_BUILD
	}
	if config.Title == "" {
		fmt.Println("provide valid title")
		return 1
	}
	if config.Sidebar == nil {
		fmt.Println("provide valid sidebar")
		return 1
	}

	//create build file
	os.RemoveAll(config.Build)
	os.Mkdir(config.Build, 0755)

	//walk throuh directory and do logic
	filepath.WalkDir(config.Source, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		md, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		meta, emd, err := extractFrontMatter(md)
		if err != nil {
			return err
		}
		htmlBody, err := renderMarkdown(emd)
		if err != nil {
			return err
		}
		if meta == nil {
			meta = &PageMeta{}
		}
		meta.Description = config.Description
		meta.Domain = config.Domain
		meta.Handle = config.Handle

		page := Page{
			Meta:        *meta,
			Title:       config.Title,
			Content:     template.HTML(string(htmlBody)),
			LeftSideBar: config.Sidebar,
			StyleSheets: []string{"/style.css"},
		}

		templ := template.Must(template.ParseFS(assets, "assets/templates.html"))
		var out bytes.Buffer
		err = templ.Execute(&out, page)
		if err != nil {
			fmt.Println(err)
			return err
		}

		relPath, _ := filepath.Rel(config.Source, path)
		outPath := filepath.Join(config.Build, strings.TrimSuffix(relPath, ".md")+".html")
		os.MkdirAll(filepath.Dir(outPath), 0755)

		err = os.WriteFile(outPath, out.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})

	//---------copy assets into the build (like custom) ----------
	styleCSS, err := readEmbedFiles("assets/style.css")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	err = os.WriteFile(filepath.Join(config.Build, "style.css"), styleCSS, 0644)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	notFoundPage, err := readEmbedFiles("assets/404.html")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	err = os.WriteFile(filepath.Join(config.Build, "404.html"), notFoundPage, 0644)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	favicon, err := os.ReadFile(filepath.Join(config.Source, "favicon.ico"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("favicon not found: `favicon.ico`")
		}
	} else {
		err = os.WriteFile(filepath.Join(config.Build, "favicon.ico"), favicon, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

	return 0
}
