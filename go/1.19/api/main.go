package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
	ID   int16  `json:"id"`
}

func respond(c *gin.Context) {
	body := Body{}
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// return the body as a response to call
	c.JSON(http.StatusAccepted, &body)
}

func main() {
	engine := gin.New()
	engine.POST("/test", respond)
	engine.Run(":3000")
}
