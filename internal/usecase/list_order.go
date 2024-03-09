package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrderListOutPutDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{OrderRepository: orderRepository}
}

func (c *ListOrderUseCase) Execute() ([]OrderListOutPutDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersDTO []OrderListOutPutDTO
	for _, order := range orders {
		order := OrderListOutPutDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Tax,
		}

		ordersDTO = append(ordersDTO, order)
	}

	return ordersDTO, nil
}
