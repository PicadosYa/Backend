package service

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type IPaymentService interface {
	CreatePayment(ctx context.Context, payment *payment.Request) error
	GetPayment(ctx context.Context, id int) (*payment.Response, error)
}

type paymentService struct {
	client payment.Client
}

func NewPaymentService() IPaymentService {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	client := payment.NewClient(cfg)
	return &paymentService{
		client: client,
	}
}

func (s *paymentService) CreatePayment(ctx context.Context, payment *payment.Request) error {
	// TODO: implement me
	return nil
}

func (s *paymentService) GetPayment(ctx context.Context, id int) (*payment.Response, error) {
	resource, err := s.client.Get(context.Background(), id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resource, nil
}
