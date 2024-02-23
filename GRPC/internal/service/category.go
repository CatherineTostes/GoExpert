package service

import (
	"context"
	"io"

	"github.com/devfullcycle/grpc/internal/database"
	"github.com/devfullcycle/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (s *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	categoryIn, err := s.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	category := &pb.Category{
		Id:          categoryIn.ID,
		Name:        categoryIn.Name,
		Description: categoryIn.Description,
	}

	return category, nil
}

func (s *CategoryService) ListCategory(ctx context.Context, blank *pb.Blank) (*pb.CategoryList, error) {
	categoriesIn, err := s.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category
	for _, categoryIn := range categoriesIn {
		category := &pb.Category{
			Id:          categoryIn.ID,
			Name:        categoryIn.Name,
			Description: categoryIn.Description,
		}
		categories = append(categories, category)
	}

	return &pb.CategoryList{Categories: categories}, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	categoryIn, err := s.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	category := &pb.Category{
		Id:          categoryIn.ID,
		Name:        categoryIn.Name,
		Description: categoryIn.Description,
	}
	return category, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		categoryResult, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (s *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryResult, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
