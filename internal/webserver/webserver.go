package webserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tcms-web-bridge/internal/dry"
	"tcms-web-bridge/internal/tcms"
	"tcms-web-bridge/internal/telegramClient"
)

type loginData struct {
	Phone string `json:"phone" binding:"required"`
}

type signData struct {
	Code string `json:"code" binding:"required"`
}

func StartWebServer(telegramClient telegramClient.TelegramClient, tcms tcms.Tcms, addConsumer chan chan []uint8) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/login", func(c *gin.Context) {
		var loginData loginData
		err := c.BindJSON(&loginData)
		dry.HandleError(err)

		err = telegramClient.Authorization(loginData.Phone)
		dry.HandleError(err)
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/sign", func(c *gin.Context) {
		var signData signData
		err := c.BindJSON(&signData)
		dry.HandleError(err)

		err = telegramClient.AuthSignIn(signData.Code)
		dry.HandleError(err)

		c.JSON(200, gin.H{"status": "ok"})
	})

	// GET
	router.GET("/me", getCurrentUser(telegramClient))
	router.GET("/dialogs", getDialogs(telegramClient))
	router.GET("/conditions", getConditions(tcms))

	// POST
	router.POST("/message", sendMessage(telegramClient))

	// websockets
	router.GET("/ws", getWcHandler(addConsumer))

	host, err := getApiHost()
	dry.HandleErrorPanic(err)
	dry.HandleErrorPanic(router.Run(host))
}
