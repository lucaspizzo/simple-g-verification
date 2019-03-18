package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		resultText := "<html><head><meta name=\"google-site-verification\" "
		resultText += fmt.Sprintf("content=\"%s\" /></head><body><code>", os.Getenv("KEY_CODE"))
		resultText += "Working...</code></body></html>"
		c.Writer.WriteHeader(200)
		c.Writer.Write([]byte(resultText))
	})
	r.Run()
}
