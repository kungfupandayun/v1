// Mock database. DO NOT EDIT.

package store

import (
	"testing"

	"github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"
	"github.com/stretchr/testify/suite"
)

type orderStoreTestSuite struct {
	suite.Suite
	s *orderStore
}

func (s *orderStoreTestSuite) SetupTest() {
	s.s = &orderStore{
		m: map[string]*orderrpc.Order{
			"1": {Id: "1"},
			"2": {Id: "2"},
			"3": {Id: "3"},
		},
	}
}

func TestOrderStoreTestSuite(t *testing.T) {
	suite.Run(t, new(orderStoreTestSuite))
}

func (s *orderStoreTestSuite) TestOrders_OK() {
	s.Len(s.s.Orders(), 3)
	s.Equal("1", s.s.Orders()[0].GetId())
}

func (s *orderStoreTestSuite) TestOrder_OK() {
	o, err := s.s.Order("2")
	s.Require().NoError(err)
	s.Equal("2", o.GetId())
}

func (s *orderStoreTestSuite) TestOrder_Err() {
	unknownID := "5"
	od, err := s.s.Order(unknownID)
	s.Nil(od)
	s.Error(err)
	s.EqualError(err, "order not found")
}

func (s *orderStoreTestSuite) TestSetOrder_OK() {
	od := &orderrpc.Order{Id: "4"}
	s.s.SetOrder(od)
	s.Len(s.s.Orders(), 4)
	o, err := s.s.Order("4")
	s.Require().NoError(err)
	s.Equal(od, o)
}
