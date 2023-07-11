package server

import (
	"net/http"

	"github.com/TobiasNickel/outfit_creator/internal/handler/outfit"
	"github.com/TobiasNickel/outfit_creator/internal/utils/staticMiddleware"
	"github.com/gin-gonic/gin"
)

func New(staticfs http.FileSystem) *gin.Engine {
	r := gin.Default()

	outfit.RegisterOutfitHandler(r.Group("/api/outfit"))
	r.Use(staticMiddleware.CreateStaticMiddleware(staticfs))
	r.SetTrustedProxies([]string{"127.0.0.1"})

	return r
}
