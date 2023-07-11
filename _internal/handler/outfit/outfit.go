package outfit

import (
	"net/http"

	"strings"

	"github.com/TobiasNickel/outfit_creator/_internal/service/outfit"
	"github.com/gin-gonic/gin"
)

func RegisterOutfitHandler(r *gin.RouterGroup) {
	r.GET("/random", func(c *gin.Context) {
		gender := c.Query("gender")
		if gender == "" {
			gender = "FEMALE"
		}
		country := c.Query("country")
		if country == "" {
			country = "de"
		}
		accessory, upper, under, err := outfit.GetRandomOutfit(strings.ToUpper(gender), strings.ToLower(country))
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
