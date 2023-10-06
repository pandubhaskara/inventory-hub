package service

import (
	"bytes"
	"context"
	"credential-auth/config"
	"credential-auth/helper"
	"credential-auth/model"
	"credential-auth/rabbitmq"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendEmail(email *model.Email) error {
	type Message struct {
		From       string
		To         string
		Cc         string
		Subject    string
		Message    string
		Attachment string
	}

	p := message.NewPrinter(language.Indonesian)

	templateData := struct {
		Name              string
		URL               string
		OrderID           string
		TransactionDate   string
		TransactionStatus string
		CustomerName      string
		CustomerEmail     string
		PaymentMethod     string
		PaymentStatus     string
		Type              string
		Quota             string
		Price             string
		Tax               string
		TotalAmount       string
	}{
		Name:              email.ReceiverName,
		URL:               fmt.Sprintf("%s/%s", email.VerificationURL, email.Token),
		OrderID:           email.OrderID,
		TransactionDate:   email.TransactionDate,
		TransactionStatus: email.TransactionStatus,
		Type:              email.Type,
		Quota:             email.Quota,
		TotalAmount:       p.Sprintf("%d", email.TotalAmount),
		Price:             p.Sprintf("%d", email.Price),
		Tax:               p.Sprintf("%d", email.Tax),
		PaymentMethod:     email.PaymentMethod,
	}

	html, err := parseTemplate("template/"+email.Template, templateData)
	if err != nil {
		helper.Logger.Error(err)
		return err
	}

	m := Message{
		From:    email.Sender,
		To:      email.ReceiverEmail,
		Subject: email.Subject,
		Message: html,
	}

	s, err := json.Marshal(m)
	if err != nil {
		helper.Logger.Error(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//send to rabbitMQ
	err = rabbitmq.ChannelPub.PublishWithContext(
		ctx,                        // context
		config.Amqp.ExchangeMail,   // exchange
		config.Amqp.RoutingKeyMail, // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         s,
		})
	if err != nil {
		helper.Logger.Error(err)
		return err
	}
	return nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		helper.Logger.Error(err)
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		helper.Logger.Error(err)
		return "", err
	}
	return buf.String(), nil
}
