package order

import (
	"context"
	"log"
	"net"
	"testing"

	orderrpc "github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"
	"github.com/bigbluedisco/tech-challenge/backend/v1/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	ords := store.NewOrderStore()
	ord := NewService(ords)
	orderrpc.RegisterServiceServer(s, ord)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestOrder_OK(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var addr = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75007", City: "Paris", Country: "France"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "1", C: customer, Addr: addr, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	assert.NoError(t, err)

}

func TestOrder_NameAbsence(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var c_err = &orderrpc.Order_Customer{LastName: "John"}
	var addr = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75007", City: "Paris", Country: "France"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "2", C: c_err, Addr: addr, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	require.Error(t, err)
}

func TestOrder_PostcodeError(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var addr_err_postcode = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75000", City: "Paris", Country: "France"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "3", C: customer, Addr: addr_err_postcode, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	assert.NoError(t, err)

}

func TestOrder_CityError(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var addr_err_city = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75007", City: "Pari", Country: "France"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "4", C: customer, Addr: addr_err_city, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	assert.NoError(t, err)

}

func TestOrder_AddrError(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var addr_err_addr = &orderrpc.Order_ShippingAddress{Address: "20 av de Ségu", PostalCode: "75007", City: "Paris", Country: "France"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "5", C: customer, Addr: addr_err_addr, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	assert.NoError(t, err)

}

func TestOrder_ProductError(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var addr = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75007", City: "Paris", Country: "France"}
	var pd_q_err = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-JACKET", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "6", C: customer, Addr: addr, ProdQuant: pd_q_err})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	require.Error(t, err)
}

func TestOrder_CountryErr(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	var customer = &orderrpc.Order_Customer{FirstName: "Joe", LastName: "John"}
	var country_err = &orderrpc.Order_ShippingAddress{Address: "20 avenue de Ségur", PostalCode: "75007", City: "Paris", Country: "Espagne"}
	var pd_q = [](*orderrpc.Order_ProductQuantity){{Pid: "PIPR-JACKET-SIZM", Quantity: 5}, {Pid: "PIPR-MOSPAD-0000", Quantity: 5}, {Pid: "PIPR-JOGCAS-SIZL", Quantity: 5}}

	resp, err := client.CreateOrder(ctx, &orderrpc.Order{Id: "7", C: customer, Addr: country_err, ProdQuant: pd_q})
	log.Printf("Response: %+v", resp)
	// Test for output here.
	require.Error(t, err)
}

func Test_ListOrders(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	require.NoError(t, err)

	defer conn.Close()
	client := orderrpc.NewServiceClient(conn)
	_, err = client.ListOrders(ctx, &orderrpc.ListOrdersRequest{})
	assert.NoError(t, err)

}

// Test to check Retrieve all orders and verify length
// Test to check product validity
// Test to check address validity
// Test multiple items
