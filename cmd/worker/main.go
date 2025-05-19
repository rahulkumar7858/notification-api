package main

import (
	"notification-system/internal/db"
	"notification-system/internal/messaging"
	"notification-system/internal/services"
)

func main() {
	db := db.InitDB()
	conn := messaging.ConnectRabbitMQ()
	defer conn.Close()

	msgs := messaging.ConsumeNotifications(conn)

	for msg := range msgs {
		notification := parseMessage(msg.Body)

		switch notification.Type {
		case "email":
			services.SendEmail(notification.To, notification.Message)
		case "sms":
			services.SendSMS(notification.To, notification.Message)
		case "inapp":
			services.SaveInAppNotification(db, notification)
		}

		msg.Ack(false)
	}
}
