// Mock database. DO NOT EDIT.

package store

import (
	"sort"
	"sync"

	"github.com/bigbluedisco/tech-challenge/backend/v1/product/rpc"
	"github.com/pkg/errors"
)

// Storage interface for products.
type ProductStore interface {
	// Retrieve a product by ID from the store.
	Product(id string) (*productrpc.Product, error)
	// Retrieve all products in the store, sorted by id asc.
	Products() []*productrpc.Product
}

type productStore struct {
	lock sync.RWMutex
	m    map[string]*productrpc.Product
}

func NewProductStore() ProductStore {
	return &productStore{
		m: map[string]*productrpc.Product{
			"PIPR-JACKET-SIZM": {
				Id:    "PIPR-JACKET-SIZM",
				Name:  "Pied Piper Jacket - Size M",
				Price: 25,
			},
			"PIPR-MOSPAD-0000": {
				Id:    "PIPR-MOSPAD-0000",
				Name:  "Silicon Valley Mousepad",
				Price: 10.5,
			},
			"PIPR-JOGCAS-SIZL": {
				Id:    "PIPR-JOGCAS-SIZL",
				Name:  "Jogging Casual - Size L",
				Price: 25,
			},
			"PIPR-PULT-SIZS": {
				Id:    "PIPR-PULT-SIZS",
				Name:  "Pull Tee-Shirt Light - Size S",
				Price: 29.99,
			},
			"PIPR-CRMSOL-50ML": {
				Id:    "PIPR-CRMSOL-50ML",
				Name:  "Crème Solaire 50ml",
				Price: 19.99,
			},
		},
	}
}

// Retrieve all products in the store, sorted by id asc.
func (s *productStore) Products() []*productrpc.Product {
	s.lock.RLock()
	defer s.lock.RUnlock()

	pts := make([]*productrpc.Product, 0, len(s.m))
	for _, o := range s.m {
		pts = append(pts, o)
	}

	sort.Slice(pts, func(i, j int) bool {
		return pts[i].GetId() < pts[j].GetId()
	})

	return pts
}

// Retrieve a product by ID from the store.
func (s *productStore) Product(id string) (*productrpc.Product, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	p, ok := s.m[id]
	if !ok {
		return nil, errors.New("product not found")
	}

	return p, nil
}
