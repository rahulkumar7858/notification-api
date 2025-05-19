package main

import (
	"metaverse-2d/OneDrive/Desktop/notification/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB and RabbitMQ
	db := db.InitDB()
	conn := messaging.ConnectRabbitMQ()
	defer conn.Close()

	// Create Gin router
	r := gin.Default()

	r.POST("/notifications", func(c *gin.Context) {
		var notification models.Notification
		c.BindJSON(&notification)

		// Save to DB
		db.Create(&notification)

		// Queue in RabbitMQ
		messaging.PublishNotification(conn, notification)

		c.JSON(200, gin.H{"status": "queued"})
	})

	r.Run(":8080")
}
