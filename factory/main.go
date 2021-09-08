package main

import "fmt"

// Abstract Factory
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS Sender
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	println("Sending SMS Notification")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Email Sender
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	println("Sending Email Notification")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "Email":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("notification type invalid")
	}
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
	sender := f.GetSender()
	fmt.Printf("Sender: %s, Channel: %s\n", sender.GetSenderMethod(), sender.GetSenderChannel())
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

	getChannel(smsFactory)
	getChannel(emailFactory)
}
