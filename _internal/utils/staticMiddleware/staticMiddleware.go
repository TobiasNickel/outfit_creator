package staticMiddleware

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * a gin middleware to serve static content, it is extra implemented to always fall back to delivering the
 * index.html, so it is suited for SPAs.
 */
func CreateStaticMiddleware(staticfs http.FileSystem) func(c *gin.Context) {
	indexFile, _ := staticfs.Open("index.html")
	indexData, err := io.ReadAll(indexFile)
	if err != nil {
		panic(err)
	}
	err = indexFile.Close()
	if err != nil {
		panic(err)

	}

	return func(c *gin.Context) {
		if has(staticfs, c.Request.URL.Path) {
			c.FileFromFS(c.Request.URL.Path, staticfs)
		} else {
			c.Data(200, "text/html", indexData)
		}
	}
}

func has(staticFileSystem http.FileSystem, name string) bool {
	file, err := staticFileSystem.Open(name)
	if err == nil {
		defer file.Close()
		return true
	} else {
		return false
	}
}
