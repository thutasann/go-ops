package howtousecomposeinterfaces

import "fmt"

type Payer interface {
	Pay(amount float64) error
}

type Refunder interface {
	Refund(amount float64) error
}

type TransactionLogger interface {
	LogTransaction(msg string)
}

type PaymentProcessor interface {
	Payer
	Refunder
	TransactionLogger
}

// Stripe Implementation
type Stripe struct{}

func (s *Stripe) Pay(amount float64) error {
	fmt.Println("Stripe: Paying", amount)
	return nil
}

func (s *Stripe) Refund(amount float64) error {
	fmt.Println("Stripe: Refunding", amount)
	return nil
}

func (s *Stripe) LogTransaction(msg string) {
	fmt.Println("Stripe log:", msg)
}

func Payment_Processor_Sample() {
	fmt.Println("\n===> Payment Processor Sample")

	var processor PaymentProcessor = &Stripe{}

	processor.Pay(100.50)
	processor.LogTransaction("Payment of 100.50 done")
	processor.Refund(50.25)
}
