package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"github.com/plutov/paypal/v4"
)

func (a *API) PaymentPrincipal(c echo.Context) error {
	// Configuración de Mercado Pago
	cfg, err := config.New(os.Getenv("MERCADO_PAGO_ACCESS_TOKEN"))
	if err != nil {
		log.Printf("Error creando la configuración de Mercado Pago: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al configurar Mercado Pago",
		})
	}

	prefClient := preference.NewClient(cfg)

	// Estructura para recibir los datos de la solicitud
	var reqBody struct {
		ID       string  `json:"id"`
		Title    string  `json:"title"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
		Email    string  `json:"email"`
		UserID   string  `json:"user_id"`
	}

	// Parsear el cuerpo de la solicitud
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Datos inválidos",
		})
	}

	// Crear la preferencia
	preferenceRequest := preference.Request{
		Items: []preference.ItemRequest{
			{
				ID:         reqBody.ID,
				Title:      reqBody.Title,
				Quantity:   reqBody.Quantity,
				UnitPrice:  reqBody.Price,
				CurrencyID: "UYU",
			},
		},
		Payer: &preference.PayerRequest{
			Email: reqBody.Email,
		},
		BackURLs: &preference.BackURLsRequest{
			Success: os.Getenv("MERCADO_PAGO_SUCCESS_URL") + "/field/" + reqBody.ID,
			Failure: os.Getenv("MERCADO_PAGO_FAILURE_URL") + "/field/" + reqBody.ID,
			Pending: os.Getenv("MERCADO_PAGO_PENDING_URL") + "/field/" + reqBody.ID,
		},
		AutoReturn: "approved",
		PaymentMethods: &preference.PaymentMethodsRequest{
			ExcludedPaymentTypes: []preference.ExcludedPaymentTypeRequest{
				{ID: "ticket"},
			},
			Installments: 1,
		},
	}

	response, err := prefClient.Create(context.Background(), preferenceRequest)
	if err != nil {
		log.Printf("Error creando preferencia: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al crear la preferencia",
		})
	}

	// Responder con el ID de la preferencia creada
	return c.JSON(http.StatusOK, map[string]string{
		"id": response.ID,
	})
}

type PaymentRequest struct {
	FieldID   string  `json:"id"`
	FieldName string  `json:"name"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

func initPayPalClient() *paypal.Client {
	client, err := paypal.NewClient(
		os.Getenv("PAYPAL_CLIENT_ID"),
		os.Getenv("PAYPAL_CLIENT_SECRET"),
		paypal.APIBaseSandBox, // Cambia a paypal.Live para producción
	)
	if err != nil {
		panic(err)
	}
	return client
}

func (a *API) createPayPalOrder(c echo.Context) error {
	ctx := c.Request().Context()
	var req PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	client := initPayPalClient()

	order, err := client.CreateOrder(ctx, paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{{
		Amount: &paypal.PurchaseUnitAmount{
			Currency: "USD",
			Value:    fmt.Sprintf("%.2f", req.Amount),
			Breakdown: &paypal.PurchaseUnitAmountBreakdown{
				ItemTotal: &paypal.Money{
					Currency: "USD",
					Value:    fmt.Sprintf("%.2f", req.Amount),
				},
			},
		},
		Items: []paypal.Item{
			{
				Name: req.FieldName,
				UnitAmount: &paypal.Money{
					Currency: "USD",
					Value:    fmt.Sprintf("%.2f", req.Amount),
				},
				Quantity: "1",
			},
		},
	}},
		nil,
		&paypal.ApplicationContext{
			ReturnURL: os.Getenv("PAYPAL_RETURN_SUCCESS_URL") + "/field/" + req.FieldID,
			CancelURL: os.Getenv("PAYPAL_RETURN_CANCEL_URL") + "/field/" + req.FieldID,
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"orderID": order.ID,
	})
}

func (a *API) capturePayPalOrder(c echo.Context) error {
	ctx := c.Request().Context()
	orderID := c.Param("orderID")
	client := initPayPalClient()

	_, err := client.CaptureOrder(ctx, orderID, paypal.CaptureOrderRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "Payment successful"})
}
