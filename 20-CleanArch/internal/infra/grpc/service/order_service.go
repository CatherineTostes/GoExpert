package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, blank *pb.Blank) (*pb.ListedOrderResponse, error) {
	ordersOut, err := s.ListOrdersUseCase.ListOrders()
	if err != nil {
		return nil, err
	}
	var orders []*pb.OrderResponse
	for _, orderOut := range ordersOut {
		order := &pb.OrderResponse{
			Id:         orderOut.ID,
			Price:      float32(orderOut.Price),
			Tax:        float32(orderOut.Tax),
			FinalPrice: float32(orderOut.FinalPrice),
		}
		orders = append(orders, order)
	}
	return &pb.ListedOrderResponse{Orders: orders}, nil
}
