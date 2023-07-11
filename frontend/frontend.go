package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed public
var BuildFS embed.FS

func BuildHTTPFS() http.FileSystem {
	buildDir, err := fs.Sub(BuildFS, "public")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FS(buildDir)

	return fs
}

func BuildHTTPDevelopmentFS() http.FileSystem {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cwd)
	buildDir := os.DirFS(filepath.Join(cwd, "frontend", "public"))
	fs := http.FS(buildDir)

	return fs
}
