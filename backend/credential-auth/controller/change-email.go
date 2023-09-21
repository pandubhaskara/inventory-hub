package controller

import (
	"context"
	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/model"
	"credential-auth/rabbitmq"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChangeEmailController struct{}

func (ctrl *ChangeEmailController) Index(c *gin.Context) {
	user := c.MustGet("account").(model.User)

	// Get the data off req body
	type bodyStruct struct {
		Email string `json:"email" binding:"required,email" `
	}

	var body bodyStruct

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		c.Abort()

		return
	}

	type message struct {
		Email    string
		NewEmail string
	}

	var m message
	m.Email = user.Email
	m.NewEmail = body.Email

	// Update it
	update := database.Db.Where("email = ?", user.Email).Model(&model.User{}).Update("email", body.Email)
	if update.Error != nil {
		helper.Logger.Error(update.Error)
		c.JSON(500, gin.H{
			"type":    "change_email",
			"message": update.Error,
		})
		c.Abort()
		return
	}

	// RABBIT MQ

	s, err := json.Marshal(m)
	if err != nil {
		helper.Logger.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// send to rabbitMQ
	err = rabbitmq.ChannelPub.PublishWithContext(
		ctx,                    // context
		config.Amqp.Exchange,   // exchange
		config.Amqp.RoutingKey, // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         s,
		})
	if err != nil {
		helper.Logger.Error(err)
	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"message": "Email has been updated successfully!",
	})
}
