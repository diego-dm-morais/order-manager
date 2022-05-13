package presenter

import (
	order "github.com/diego-dm-morais/order-manager/usecase/order"
)

func CreateOrderPresenter() order.IOrderPresenter {
	return &orderPresenter{}
}
