// Mock database. DO NOT EDIT.

package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type productStoreTestSuite struct {
	suite.Suite
	s ProductStore
}

func (s *productStoreTestSuite) SetupTest() {
	s.s = NewProductStore()
}

func TestProductStoreTestSuite(t *testing.T) {
	suite.Run(t, new(productStoreTestSuite))
}

func (s *productStoreTestSuite) TestProducts_OK() {
	s.Len(s.s.Products(), 5)
}

func (s *productStoreTestSuite) TestProduct_OK() {
	pdt, err := s.s.Product("PIPR-MOSPAD-0000")
	s.Require().NoError(err)
	s.Equal("PIPR-MOSPAD-0000", pdt.GetId())
}

func (s *productStoreTestSuite) TestProduct_Err() {
	unknownID := "PIPR-MOSPAD-0001"
	pdt, err := s.s.Product(unknownID)
	s.Nil(pdt)
	s.Error(err)
	s.EqualError(err, "product not found")
}
