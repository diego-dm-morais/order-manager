package order

import (
	"github.com/diego-dm-morais/order-manager/entities/cliente"
	"github.com/diego-dm-morais/order-manager/entities/endereco"
	"github.com/diego-dm-morais/order-manager/entities/item"
	"github.com/diego-dm-morais/order-manager/entities/pedido"
	"github.com/diego-dm-morais/order-manager/entities/produto"
	"github.com/diego-dm-morais/order-manager/usecase/address"
	"github.com/diego-dm-morais/order-manager/usecase/customer"
	itemUsecase "github.com/diego-dm-morais/order-manager/usecase/item"
	"github.com/diego-dm-morais/order-manager/usecase/product"
)

type orderInteractor struct {
	orderOutputBoundary IOrderOutputBoundary
	productGateway      product.IProductGateway
	customerGateway     customer.ICustomerGateway
	orderGateway        IOrderGateway
}

func (o *orderInteractor) Salvar(order OrderRequest) (*OrderResponse, error) {
	itens := *o._GetItens(order.Items)
	cliente := *o._GetCustomer(order.CustomerID)
	endereco := *o._GetAddress(order.ShippingAddress)
	var pedido pedido.IPedido = pedido.New().SetItens(itens).SetCliente(cliente).SetEnderecoEntrega(endereco).SetFrete(order.Freight).Build()
	_, erro := pedido.EValido()
	if erro != nil {
		return nil, erro
	}

	var orderDataRequest OrderDataRequest = OrderDataRequest{
		IdentificationDocument: cliente.GetDocumentoIdentificacao(),
		Freight:                pedido.GetFrete(),
		Total:                  pedido.GetTotal(),
		Itens:                  o._GetItensData(itens),
		ShippingAddress: address.ShippingAddressDataRequest{
			Street:  endereco.GetRua(),
			Number:  endereco.GetNumero(),
			Zipcode: endereco.GetCep(),
			City:    endereco.GetCidade(),
		},
	}

	orderID, erroData := o.orderGateway.save(orderDataRequest)
	if erroData != nil {
		return nil, erroData
	}
	return o.orderOutputBoundary.success(OrderInputData{OrderID: orderID, CustomerName: cliente.GetNome()}), nil
}

func (o *orderInteractor) _GetItens(itensRequest []itemUsecase.ItemRequest) *[]item.IItem {
	var itens []item.IItem
	for _, it := range itensRequest {
		var product product.ProductResponse = o.productGateway.FindByProduct(it.ProductID)
		var produto produto.IProduto = produto.New().SetNome(product.Nome).SetPreco(product.Price).SetEstoqueEstaDisponivel(product.EstoqueEstaDisponivel).Build()
		var item item.IItem = item.New().SetProduto(produto).SetQuantidade(it.Amount).Build()
		itens = append(itens, item)
	}
	return &itens
}

func (o *orderInteractor) _GetCustomer(customerID string) *cliente.ICliente {
	var customer customer.CustomerResponse = o.customerGateway.FindByCustomer(customerID)
	var cliente cliente.ICliente = cliente.New().SetNome(customer.Name).SetDocumentoIdentificacao(customer.IdentificationDocument).SetTelefone(customer.Telephone).Build()
	return &cliente
}

func (o *orderInteractor) _GetAddress(address address.ShippingAddressRequest) *endereco.IEndereco {
	var endereco endereco.IEndereco = endereco.New().SetCep(address.Zipcode).SetCidade(address.City).SetNumero(address.Number).SetRua(address.Street).Build()
	return &endereco
}

func (o *orderInteractor) _GetItensData(itens []item.IItem) []itemUsecase.ItemDataRequest {
	var itensData []itemUsecase.ItemDataRequest
	for _, it := range itens {
		itensData = append(itensData, itemUsecase.ItemDataRequest{
			ProductName: it.GetProduto().GetNome(),
			Price:       it.GetProduto().GetPreco(),
			Amount:      it.GetQuantidade(),
		})
	}
	return itensData
}