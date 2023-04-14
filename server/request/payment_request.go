package request

type CreatePaymentRequest struct {
	PaymentType string `json:"paymentType" validate:"required"`
}
