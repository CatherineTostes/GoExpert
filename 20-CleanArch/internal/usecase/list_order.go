package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type (
	OrderResponseDTO struct {
		ID         string  `json:"id"`
		Price      float64 `json:"price"`
		Tax        float64 `json:"tax"`
		FinalPrice float64 `json:"final_price"`
	}

	ListOrderUseCase struct {
		OrderRepository entity.OrderRepositoryInterface
	}
)

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) ListOrders() ([]OrderResponseDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []OrderResponseDTO
	for _, o := range orders {
		response = append(response, OrderResponseDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return response, nil
}
