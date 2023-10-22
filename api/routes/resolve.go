package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rejrak/url-shortener/database"
)

func ResolveURL(c *gin.Context) {
	url := c.Param("url")

	rdb := database.CreateClient(0)
	defer rdb.Close()
	value, err := rdb.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "short not found in the database"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot connect do DB"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	c.Redirect(http.StatusPermanentRedirect, value)
}
