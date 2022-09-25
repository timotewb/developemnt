package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
	ID   int16  `json:"id"`
}

func respond(c *gin.Context) {
	fmt.Println(c.Request.Body)
	body := Body{}
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Error: c.BindJSON()")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(filepath.Join("apps", body.Name))
	cmd := exec.Command(filepath.Join("apps", body.Name))
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("cmd.Output()")
		fmt.Println(err.Error())
		return
	}
	// return the body as a response to call
	c.Data(http.StatusOK, "application/json", stdout)
}

func main() {
	engine := gin.Default()
	engine.POST("/test", respond)
	engine.Run(":3000")
}
