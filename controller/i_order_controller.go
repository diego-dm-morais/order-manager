package controller

import order "github.com/diego-dm-morais/order-manager/usecase/order"

type IOrderController interface {
	Save(OrderRequest order.OrderRequest) (*order.OrderResponse, error)
}