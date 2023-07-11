package outfit

import (
	"net/http"

	"strings"

	"github.com/TobiasNickel/outfit_creator/internal/service/outfit"
	"github.com/gin-gonic/gin"
)

func RegisterOutfitHandler(r *gin.RouterGroup) {
	r.GET("/random", func(c *gin.Context) {
		gender := c.Query("gender")
		if gender == "" {
			gender = "FEMALE"
		}
		language := c.Query("language")
		if language == "" {
			language = "de"
		}
		accessory, upper, under, err := outfit.GetRandomOutfit(strings.ToUpper(gender), strings.ToLower(language))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(200, gin.H{
			"message":   "outfit",
			"accessory": accessory,
			"upper":     upper,
			"under":     under,
		})
	})
}
