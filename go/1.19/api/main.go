package main

import (
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to de-serialize json body
	Name string   `json:"name"`
	Args []string `json:"args"`
}

func respond(c *gin.Context) {

	body := Body{}
	// using BindJson method to serialize body with struct
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "Note": "Please check you have provided the job name in the correct format {Name:Job Name}."})
		return
	}
	//fn := filepath.Join("/mnt/ns01/servers/factotum/01/api/apps", body.Name)
	fn := filepath.Join("apps", body.Name)
	cmd := exec.Command(fn, body.Args...)

	stdout, err := cmd.Output()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err, "Note": "Job not found", "File Name": fn})
		return
	}
	// return the body as a response to call
	c.Data(http.StatusOK, "application/json", stdout)
}

func main() {
	engine := gin.Default()
	engine.POST("/api", respond)
	engine.Run(":3000")
}
