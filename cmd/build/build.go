package build

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

// TODO: change this in prod
const (
	CONFIG_FILE    = "example.config.yaml"
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
	Meta         PageMeta
	LeftSideBar  template.HTML
	Content      template.HTML
	RightSideBar template.HTML
}

func Build() int {
	//load config
	byte, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Print("Missing '%v'", CONFIG_FILE)
			return 1
		}
		fmt.Println(err)
		return 1
	}
	var config Config
	err = yaml.Unmarshal(byte, &config)
	if err != nil {
		fmt.Println(err)
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

	//copy the asset files
	err = copyDir("cmd/build/assets", filepath.Join(config.Build, "assets"))
	if err != nil {
		fmt.Println(err)
		return 1
	}

	stpl := template.Must(template.ParseFiles("cmd/build/templates/sidebar.html"))
	var sidebar bytes.Buffer
	err = stpl.Execute(&sidebar, config.Sidebar)
	if err != nil {
		fmt.Println(err)
		return 1
	}
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
		page := Page{
			Meta:         *meta,
			Content:      template.HTML(string(htmlBody)),
			LeftSideBar:  template.HTML(sidebar.String()),
			RightSideBar: template.HTML(""),
		}
		tpl := template.Must(template.ParseFiles("cmd/build/templates/layout.html"))
		var out bytes.Buffer
		err = tpl.Execute(&out, page)
		if err != nil {
			log.Fatal(err)
		}

		relPath, _ := filepath.Rel(config.Source, path)
		outPath := filepath.Join(config.Build, strings.TrimSuffix(relPath, ".md")+".html")
		os.MkdirAll(filepath.Dir(outPath), 0755)

		err = os.WriteFile(outPath, out.Bytes(), 0644)
		if err != nil {
			return err
		}
		err = os.WriteFile(outPath, out.Bytes(), 0644)
		if err != nil {
			return err
		}

		return nil
	})

	return 0
}
