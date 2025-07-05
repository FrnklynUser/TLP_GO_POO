package main

import "fmt"

type PaymentProcessor interface {
	Process(amount float64) error
	GetFee() float64
}

type CreditCardProcessor struct {
	CardNumber string
	FeeRate    float64
}

func (cc CreditCardProcessor) Process(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("monto inválido")
	}
	fmt.Printf("💳 Procesando $%.2f con tarjeta ****%s\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
	return nil
}

func (cc CreditCardProcessor) GetFee() float64 {
	return cc.FeeRate
}

type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) Process(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("monto inválido")
	}
	fmt.Printf("🅿 Procesando $%.2f via PayPal (%s)\n", amount, pp.Email)
	return nil
}

func (pp PayPalProcessor) GetFee() float64 {
	return 0.025 // 2.5%
}

func ProcessOrder(processor PaymentProcessor, amount float64) error {
	fee := processor.GetFee()
	total := amount + (amount * fee)
	fmt.Printf("🧮 Total con comisión: $%.2f (comisión: %.1f%%)\n", total, fee*100)
	return processor.Process(amount)
}

func main() {
	creditCard := CreditCardProcessor{
		CardNumber: "1234567890123456",
		FeeRate:    0.035,
	}

	paypal := PayPalProcessor{
		Email: "rortizjohnfrank@example.com",
	}

	ProcessOrder(creditCard, 100)
	ProcessOrder(paypal, 100)
}
