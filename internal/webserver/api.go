package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"tcms-web-bridge/internal/tcms"
	"tcms-web-bridge/internal/telegramClient"
	tcms2 "tcms-web-bridge/pkg/tcms"
)

type sendMessageData struct {
	AccessHash int64  `json:"accessHash" binding:"required"`
	Id         int32  `json:"id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// getCurrentUser GET /me
func getCurrentUser(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		user, err := telegramClient.GetCurrentUser(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		m := jsonpb.Marshaler{
			EmitDefaults: true,
		}
		s, err := m.MarshalToString(user)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, s)
	}
}

// sendMessage POST /message
func sendMessage(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		var messageData sendMessageData
		err := c.BindJSON(&messageData)
		if err != nil {
			_ = c.Error(err)
			return
		}

		err = telegramClient.SendMessage(c, string(messageData.Id), messageData.Message)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(200, gin.H{"status": "ok"})
	}
}

// getDialogs GET /dialogs
func getDialogs(telegramClient telegramClient.TelegramClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		dialogs, err := telegramClient.Dialogs(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		m := jsonpb.Marshaler{
			EmitDefaults: true,
		}
		s, err := m.MarshalToString(dialogs)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, s)
	}
}

// getConditions GET /condition
func getConditions(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		conditions, err := tcms.GetConditions(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, conditions)
	}
}

// getActions GET /action
func getActions(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		actions, err := tcms.GetActions(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, actions)
	}
}

// getTriggers GET /trigger
func getTriggers(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		triggers, err := tcms.GetTriggers(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, triggers)
	}
}

// getAutomations GET /automation
func getAutomations(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		automations, err := tcms.GetAutomations(c)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, automations)
	}
}

// addAutomation POST /automation
func addAutomation(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		var automation tcms2.Automation

		if err := c.BindJSON(&automation); err != nil {
			_ = c.Error(err)
			return
		}

		if err := tcms.AddAutomation(c, &automation); err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}

// updateAutomation PATCH /automation
func updateAutomation(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request tcms2.UpdateAutomationRequest

		if err := c.BindJSON(&request); err != nil {
			_ = c.Error(err)
			return
		}

		if err := tcms.UpdateAutomation(c, &request); err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}

// removeAutomation DELETE /automation
func removeAutomation(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request tcms2.RemoveAutomationRequest

		if err := c.BindJSON(&request); err != nil {
			_ = c.Error(err)
			return
		}

		if err := tcms.RemoveAutomation(c, &request); err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}

// getAutomation GET /automation/:id
func getAutomation(tcms tcms.Tcms) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		request := tcms2.AutomationRequest{Id: id}

		res, err := tcms.GetAutomation(c, &request)
		if err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(200, res)
	}
}
