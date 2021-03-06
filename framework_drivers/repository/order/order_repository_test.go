package repository

import (
	"testing"

	order "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSalverPedido(t *testing.T) {
	connectorMongoDataSource := new(MockConnectorMongoDataSource)
	orderRepository := CreateOrderRepository(connectorMongoDataSource)

	connectorMongoDataSource.On("Connect").Return(&mongo.Client{}, nil)
	connectorMongoDataSource.On("DataSource").Return(&mongo.Collection{}, nil)
	connectorMongoDataSource.On("Save").Return("627d7a330ce9a74dc8c6b1d7", nil)
	connectorMongoDataSource.On("Disconnect").Return(true, nil)

	inputMapper := order.OrderInputMapper{
		IdentificationDocument: "999.999.999-00",
		Freight:                55.0,
		Total:                  17555.0,
		Itens: []order.ItemInputMapper{
			{
				ProductName: "Macbook pro 15",
				Price:       17500.0,
				Amount:      1,
			},
		},
		ShippingAddress: order.AddressInputMapper{
			Street:  "Rua new street",
			Number:  "8987",
			Zipcode: "89898-09",
			City:    "São Paulo - SP",
		},
	}

	id, _ := orderRepository.Save(inputMapper)

	assert.NotNil(t, id)
	assert.NotEmpty(t, id)
	assert.Equal(t, "627d7a330ce9a74dc8c6b1d7", id)
}
