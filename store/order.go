// Mock database. DO NOT EDIT.

package store

import (
	"sort"
	"sync"

	"github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"
	"github.com/pkg/errors"
)

// Storage interface for orders.
type OrderStore interface {
	// Retrieve an order by ID from the store.
	Order(id string) (*orderrpc.Order, error)
	// Retrieve all orders in the store, sorted by id asc.
	Orders() []*orderrpc.Order
	// Upsert an order in the store.
	SetOrder(*orderrpc.Order)
}

type orderStore struct {
	lock sync.RWMutex
	m    map[string]*orderrpc.Order
}

// Create a new order store.
func NewOrderStore() OrderStore {
	return &orderStore{
		m: make(map[string]*orderrpc.Order),
	}
}

// Retrieve all orders in the store, sorted by id asc.
func (s *orderStore) Orders() []*orderrpc.Order {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ods := make([]*orderrpc.Order, 0, len(s.m))
	for _, o := range s.m {
		ods = append(ods, o)
	}

	sort.Slice(ods, func(i, j int) bool {
		return ods[i].GetId() < ods[j].GetId()
	})

	return ods
}

// Retrieve an order by ID from the store.
func (s *orderStore) Order(id string) (*orderrpc.Order, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	o, ok := s.m[id]
	if !ok {
		return nil, errors.New("order not found")
	}

	return o, nil
}

// Upsert an order in the store.
func (s *orderStore) SetOrder(o *orderrpc.Order) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.m[o.GetId()] = o
}
