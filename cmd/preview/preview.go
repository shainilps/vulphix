package preview

import (
	"errors"
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/codeshaine/vulpix/cmd/build"
)

const previewPort = 8080

func PreviewBuild() int {

	config, err := build.ReadConfig()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) || config == nil {
			fmt.Printf("Missing '%v'\n", build.CONFIG_FILE)
		}
		return 1
	}

	dir := build.DEFAULT_BUILD
	if config.Build != "" {
		dir = config.Build
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		extension := filepath.Ext(req.URL.Path)
		if extension != "" {
			data, err := os.ReadFile(path.Join(dir, req.URL.Path))
			if err != nil {
				w.WriteHeader(404)
				w.Write([]byte("404 - Not found"))
				return
			}
			w.Header().Set("Content-Type", mime.TypeByExtension(extension))
			w.Write(data)
			return
		}
		html, err := resolveHTMLRequest(req.URL.Path)
		if errors.Is(err, fs.ErrNotExist) {
			html, _ = os.ReadFile(dir + "/404.html")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(404)
			w.Write(html)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(html)
	})

	fmt.Printf("Preview server running at http://localhost:%d/\n", previewPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", previewPort), nil)
	if err != nil {
		fmt.Printf("Failed to start preview server: %v\n", err)
		return 1
	}
	return 0
}

func resolveHTMLRequest(requestPath string) ([]byte, error) {
	html, err := os.ReadFile(path.Join(build.DEFAULT_BUILD, requestPath+".html"))
	if err == nil {
		return html, nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	html, err = os.ReadFile(path.Join("dist", requestPath, "index.html"))
	return html, err
}
