package framework

import (
	orderUseCase "github.com/diego-dm-morais/order-manager/application_business_rules/usecase/order"
	customerApi "github.com/diego-dm-morais/order-manager/framework_drivers/apiclient/customer"
	productApi "github.com/diego-dm-morais/order-manager/framework_drivers/apiclient/product"

	database "github.com/diego-dm-morais/order-manager/framework_drivers/repository/database/mongo"
	//database "github.com/diego-dm-morais/order-manager/framework_drivers/repository/database/postgres"
	orderRepository "github.com/diego-dm-morais/order-manager/framework_drivers/repository/order"
	orderController "github.com/diego-dm-morais/order-manager/interface_adapters/controller"
	customerGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/customer"
	orderGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/order"
	productGateway "github.com/diego-dm-morais/order-manager/interface_adapters/gateway/product"
	presenter "github.com/diego-dm-morais/order-manager/interface_adapters/presenter"
)

func CreateOrderRest() IOrderRest {
	//connectorDataSource := database.CreateConnectorPostgresDataSource()
	//orderRepository := orderRepository.CreateOrderRepository(connectorDataSource)

	connectorMongoDataSource := database.CreateConnectorMongoDataSource()
	orderRepository := orderRepository.CreateOrderRepository(connectorMongoDataSource)
	orderGateway := orderGateway.CreateOrderGateway(orderRepository)
	customerApi := customerApi.CreateCustomerApi()
	customerGateway := customerGateway.CreateCustomerGateway(customerApi)
	productApi := productApi.CreateProductApi()
	productGateway := productGateway.CreateProductGateway(productApi)
	orderPresenter := presenter.CreateOrderPresenter()
	orderUseCase := orderUseCase.CreateOrderUseCase(orderPresenter, productGateway, customerGateway, orderGateway)
	orderController := orderController.CreateOrderController(orderUseCase)

	return &orderRest{orderController: orderController}
}
