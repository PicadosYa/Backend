package api

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func (a *API) PaymentPrincipal(c echo.Context) error {
	// Configuración de Mercado Pago
	cfg, err := config.New("TEST-1442054152662695-111115-391930c18f616b088fa289b70e9ed314-29578319")
	if err != nil {
		log.Printf("Error creando la configuración de Mercado Pago: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error al configurar Mercado Pago",
		})
	}

	prefClient := preference.NewClient(cfg)

	// Estructura para recibir los datos de la solicitud
	var reqBody struct {
		Title    string  `json:"title"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
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
				Title:      reqBody.Title,
				Quantity:   reqBody.Quantity,
				UnitPrice:  reqBody.Price,
				CurrencyID: "UYU",
			},
		},
		BackURLs: &preference.BackURLsRequest{
			Success: "https://www.google.com.uy",
			Failure: "https://www.youtube.com",
			Pending: "https://www.elpais.com.uy",
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