package preview

import (
	"fmt"
	"net/http"

	"github.com/codeshaine/vulpix/cmd/build"
)

const previewPort = 8080 // you can change this

func PreviewBuild() int {
	const dir = build.DEFAULT_BUILD // or "dist" if you renamed DEFAULT_BUILD

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Printf("Preview server running at http://localhost:%d/\n", previewPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", previewPort), nil)
	if err != nil {
		fmt.Printf("Failed to start preview server: %v\n", err)
		return 1
	}
	return 0
}
