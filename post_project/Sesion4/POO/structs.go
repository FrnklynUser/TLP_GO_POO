package main

import "fmt"

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println("LOG:", message)
}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct {
	Logger
	EmailAddress string
}

func (e EmailNotifier) Send(message string) {
	e.Log("Sending email...")
	fmt.Printf("📧 Enviado a %s: %s\n", e.EmailAddress, message)
}

type SMSNotifier struct {
	Logger
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) {
	s.Log("Sending SMS...")
	fmt.Printf("📱 Enviado a %s: %s\n", s.PhoneNumber, message)
}

func SendNotification(n Notifier, message string) {
	n.Send(message)
}

func main() {
	email := EmailNotifier{EmailAddress: "admin@example.com"}
	sms := SMSNotifier{PhoneNumber: "+51 920123568"}

	SendNotification(email, "Bienvenido al sistema")
	SendNotification(sms, "Tu pedido está listo")
}
