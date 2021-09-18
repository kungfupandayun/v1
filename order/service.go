package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	orderrpc "github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"

	"github.com/bigbluedisco/tech-challenge/backend/v1/store"
)

// Service holds RPC handlers for the order service. It implements the orderrpc.ServiceServer interface.
type service struct {
	orderrpc.UnimplementedServiceServer
	s store.OrderStore
}

func NewService(s store.OrderStore) *service {
	return &service{s: s}
}

// Fetch all existing orders in the system.
func (s *service) ListOrders(ctx context.Context, r *orderrpc.ListOrdersRequest) (*orderrpc.ListOrdersResponse, error) {
	return &orderrpc.ListOrdersResponse{Orders: s.s.Orders()}, nil
}

// Quest to website (return addr, postcode,city )
func Quest(path string) (string, string, string, error) {
	res, err := http.Get(path)
	if err != nil {
		return "", "", "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "", "", "", err
	}

	var result ModelAdr
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return "", "", "", err
	}

	// Check if City and PostCode not found
	if len(result.Features) == 0 || err != nil {
		return "", "", "", errors.New("address not found")
	}

	postcode := result.Features[0].Properties.PostCode
	city := result.Features[0].Properties.City
	adrname := result.Features[0].Properties.Name
	return adrname, postcode, city, nil
}

// Verify Address
func VerifyAddr(order *orderrpc.Order) error {

	// 1) Get Country Correct
	country := order.GetAddr().GetCountry()
	if !(strings.EqualFold(country, "france") || strings.EqualFold(country, "fr") || strings.EqualFold(country, "")) {
		return errors.New("send in France only")
	}
	order.Addr.Country = "France"

	// 2) Get Address correct
	city := order.Addr.GetCity()
	postcode := order.Addr.GetPostalCode()
	addr := order.Addr.GetAddress()

	query := "q=" + url.QueryEscape(addr) + "&" +
		"city=" + url.QueryEscape(city) + "||" +
		"postcode=" + url.QueryEscape(postcode)
	path := fmt.Sprintf("https://api-adresse.data.gouv.fr/search/?" + query)

	// Query to website https://geo.api.gouv.fr/adresse
	addrq, codeq, cityq, err := Quest(path)
	if err != nil {
		return err
	}

	// Rename Order Address ,City, PostCode in Order
	order.Addr.Address = addrq
	order.Addr.City = cityq
	order.Addr.PostalCode = codeq

	return nil
}

// Create an order
func (s *service) CreateOrder(ctx context.Context, order *orderrpc.Order) (*orderrpc.CreateOrderResponse, error) {

	// Check customer first and last name not empty
	if order.GetC().FirstName == "" || order.GetC().LastName == "" {
		return nil, errors.New("customer name not completed")
	}

	// Verify all the product exists in the product store
	pdstor := store.NewProductStore()
	for i := 0; i < len(order.ProdQuant); i++ {
		if item, err := pdstor.Product(order.ProdQuant[i].Pid); item == nil || err != nil {
			return nil, errors.New("product (" + order.ProdQuant[i].Pid + ") not found")
		}
	}

	// Verify the Shipping Address is valid , and do autocorrect if minor error
	err := VerifyAddr(order)
	if err != nil {
		return nil, err
	}

	// Upsert Order to OrderStore
	s.s.SetOrder(order)

	return &orderrpc.CreateOrderResponse{}, nil
}
