package model

type Email struct {
	ReceiverName      string
	ReceiverEmail     string
	Sender            string
	Subject           string
	Template          string
	Token             string
	ExpirationTime    string
	VerificationURL   string
	Title             string
	TitleType         string
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
}
