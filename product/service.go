package product

import (
	"context"

	productrpc "github.com/bigbluedisco/tech-challenge/backend/v1/product/rpc"
	"github.com/bigbluedisco/tech-challenge/backend/v1/store"
)

// Service holds RPC handlers for the product service. It implements the product.ServiceServer interface.
type service struct {
	productrpc.UnimplementedServiceServer
	s store.ProductStore
}

func NewService(s store.ProductStore) *service {
	return &service{s: s}
}

// Fetch all existing products in the system.
func (s *service) ListProducts(ctx context.Context, r *productrpc.ListProductsRequest) (*productrpc.ListProductsResponse, error) {
	return &productrpc.ListProductsResponse{Products: s.s.Products()}, nil
}

// Fetch one product through its id
func (s *service) FetchProduct(ctx context.Context, r *productrpc.FetchProductsRequest) (*productrpc.Product, error) {
	return s.s.Product(r.Id)
}
