package service

import (
	"credential-auth/config"
	"credential-auth/database"
	"credential-auth/helper"
	"credential-auth/model"
	"credential-auth/rabbitmq"
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type user struct {
	Name              string
	Email             string
	Sender            string
	ClientUrl         string
	Subject           string
	Template          string
	Applications      []int
	OrderID           string
	TransactionDate   string
	TransactionStatus string
	CustomerName      string
	CustomerEmail     string
	PaymentMethod     string
	PaymentStatus     string
	Type              string
	Quota             string
	Price             int
	Tax               int
	TotalAmount       int
	ContactEmail      string
}

func AccountSubscriber() {
	// Declare queue
	q, err := rabbitmq.ChannelSub.QueueDeclare(
		config.Amqp.QueueRegister, // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		helper.Logger.Error(err)
		return
	}

	// Queue bind
	err = rabbitmq.ChannelSub.QueueBind(
		q.Name,                         // queue name
		config.Amqp.RoutingKeyRegister, // routing key
		config.Amqp.Exchange,           // exchange
		false,
		nil)
	if err != nil {
		helper.Logger.Error(err)
		return
	}

	msgs, err := rabbitmq.ChannelSub.Consume(
		q.Name,          // queue
		config.App.Name, // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)

	if err != nil {
		helper.Logger.Error(err)
		return
	}

	go func() {
		for d := range msgs {
			var user user

			if err := json.Unmarshal(d.Body, &user); err != nil {
				helper.Logger.Error(err)
				return
			}

			database.Db.Transaction(func(tx *gorm.DB) error {
				var u model.User

				result := tx.FirstOrCreate(&u, model.User{
					Email: user.Email,
				})
				if result.Error != nil {
					helper.Logger.Error(result.Error)
					return result.Error
				}

				if result.RowsAffected == 1 {
					// Verification token
					verificationToken := uuid.New()
					verificationTokenDuration, _ := strconv.ParseInt(config.RegistrationMailer.ExpirationTime, 10, 64)

					// Save verification token
					newVerificationToken := model.VerificationToken{
						UserID:    int(u.ID),
						Token:     verificationToken.String(),
						IssuedAt:  time.Now(),
						ExpiredAt: time.Now().Add((time.Second * time.Duration(verificationTokenDuration))),
						IsActive:  true,
					}

					result = tx.Create(&newVerificationToken)
					if result.Error != nil {
						helper.Logger.Error(result.Error)
						return result.Error
					}

					template := "verification-email.html"
					subject := "Account Verification"

					if user.Template != "" {
						template = user.Template
					}

					if user.Subject != "" {
						subject = user.Subject
					}

					// Email format
					email := model.Email{
						ReceiverName:      user.Name,
						ReceiverEmail:     user.Email,
						Sender:            config.RegistrationMailer.SenderName,
						Subject:           subject,
						Template:          template,
						Token:             verificationToken.String(),
						ExpirationTime:    config.RegistrationMailer.ExpirationTime,
						VerificationURL:   config.App.Client + config.RegistrationMailer.VerificationUrl,
						OrderID:           user.OrderID,
						TransactionDate:   user.TransactionDate,
						TransactionStatus: user.TransactionStatus,
						CustomerName:      user.CustomerName,
						CustomerEmail:     user.CustomerEmail,
						PaymentMethod:     user.PaymentMethod,
						PaymentStatus:     user.PaymentStatus,
						Type:              user.Type,
						Quota:             user.Quota,
						Price:             user.Price,
						Tax:               user.Tax,
						TotalAmount:       user.TotalAmount,
					}

					// Send email module
					SendEmail(&email)
				}

				d.Ack(false)

				return nil
			})
		}
	}()

	helper.Logger.Info("[register-subscriber] Waiting for messages. To exit press CTRL+C")
}
