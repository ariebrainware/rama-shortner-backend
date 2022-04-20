package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ariebrainware/rama-shortner-backend/endpoint"
	"github.com/ariebrainware/rama-shortner-backend/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var success string = fmt.Sprintf("Server listening with version %s", os.Getenv("VERSION"))
		c.JSON(http.StatusOK, &model.Response{
			Success: true,
			Error:   nil,
			Msg:     success,
			Data:    nil,
		})
	})
	r.POST("/url/short", endpoint.ShortURL)
	r.GET("/url/:key", endpoint.GetURL)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	log.Infof("Service version: %s", os.Getenv("VERSION"))
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error(err)
	}
}
