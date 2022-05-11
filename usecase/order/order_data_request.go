package order

import (
	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/item"
)

type OrderDataRequest struct {
	IdentificationDocument string
	Freight                float32
	Itens                  []item.ItemDataRequest
	Total                  float32
	ShippingAddress        address.ShippingAddressDataRequest
}