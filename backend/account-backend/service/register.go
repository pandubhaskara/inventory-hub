package service

import (
	"account-backend/config"
	"account-backend/database"
	"account-backend/helper"
	"account-backend/model"
	"account-backend/rabbitmq"
	"encoding/json"
)

func RegisterSubscriber() {
	// declare queue
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

	// queue bind
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
		helper.Logger.Error("regina-register consume error: ", err)
		return
	}

	type user struct {
		Name                   string
		Email                  string
		MobileNumber           string
		NationalIdentityNumber string
		Applications           []int
	}

	go func() {
		for d := range msgs {
			var u user

			if err := json.Unmarshal(d.Body, &u); err != nil {
				helper.Logger.Error(err)
				return
			}

			apps := []model.Application{}

			if len(u.Applications) > 0 {
				for _, appID := range u.Applications {
					apps = append(
						apps,
						model.Application{
							ID: uint(appID),
						})

				}
			}

			newAccount := model.Account{
				Name:                   u.Name,
				Email:                  u.Email,
				MobileNumber:           &u.MobileNumber,
				NationalIdentityNumber: &u.NationalIdentityNumber,
				Applications:           apps,
			}

			result := database.Db.Where(model.Account{Email: u.Email}).FirstOrCreate(&newAccount)
			if result.Error != nil {
				helper.Logger.Error(result.Error)
				return
			}

			d.Ack(false)
		}
	}()

	helper.Logger.Info(" [register-subscriber] Waiting for messages. To exit press CTRL+C")
}
