package webserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tcms-web-bridge/internal/config"
	"tcms-web-bridge/internal/tcms"
	"tcms-web-bridge/internal/telegramClient"
)

type loginData struct {
	Phone string `json:"phone" binding:"required"`
}

type signData struct {
	Code string `json:"code" binding:"required"`
}

// StartWebServer start web server
func StartWebServer(config config.Config, telegramClient telegramClient.TelegramClient, tcms tcms.Tcms, addConsumer chan chan []uint8) {
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
		if err != nil {
			_ = c.Error(err)
			return
		}

		err = telegramClient.Authorization(c, loginData.Phone)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/sign", func(c *gin.Context) {
		var signData signData
		err := c.BindJSON(&signData)
		if err != nil {
			_ = c.Error(err)
			return
		}

		err = telegramClient.AuthSignIn(c, signData.Code)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	})

	// GET
	router.GET("/me", getCurrentUser(telegramClient))
	router.GET("/dialogs", getDialogs(telegramClient))
	router.GET("/condition", getConditions(tcms))
	router.GET("/action", getActions(tcms))
	router.GET("/automation", getAutomations(tcms))
	router.GET("/trigger", getTriggers(tcms))
	router.GET("/automation/:id", getAutomation(tcms))

	// POST
	router.POST("/message", sendMessage(telegramClient))
	router.POST("/automation", addAutomation(tcms))

	// PATCH
	router.PATCH("/automation", updateAutomation(tcms))

	// DELETE
	router.DELETE("/automation", removeAutomation(tcms))

	// websockets
	router.GET("/ws", getWcHandler(addConsumer))

	if err := router.Run(config.ApiHost); err != nil {
		panic(err)
	}
}
