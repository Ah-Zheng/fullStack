package api

import (
	"full-stack/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SetupApi() *gin.Engine {
	r := gin.New()
	r.Use(static.Serve("/", static.LocalFile("./public/", false)))

	api := r.Group("/api")
	{
		api.GET("/test", testHandler)
	}

	r.GET("/ws", socketHandler)

	return r
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func socketHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal("Error during connection upgradation: ", err)
		return
	}

	for {
		msgType, msg, err := conn.ReadMessage()

		if err != nil {
			log.Fatal("Error during message reading: ", err)
			break
		}

		log.Printf("Received: %s", msg)

		// sendText := string(msg) + " from backend"
		err = conn.WriteMessage(msgType, []byte(utils.CaptureImg()))

		if err != nil {
			log.Println("Error during message writing: ", err)
			break
		}
	}
}
