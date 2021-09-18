# Software Engineer Challenge (Backend)
##  Introduction
In logistics, inaccurate addresses are the primary reason why orders do not arrive on time
([Magento, 2018](https://magento.com/sites/default/files8/fixing-failed-deliveries-community-insight.pdf)).


At Bigblue, our mission is to create the ultimate delivery experience for brands. We sync e-commerce orders from our merchants' stores in real time, and have to ensure that their addresses are valid to ensure a frictionless experience.

**Your challenge is to design and implement address validation for orders shipped to France.**


## Guidelines

- We value a clean, simple working solution. Some code is already provided, so that you just have to write over it.
- The project must be submitted as a git repository (github, bitbucket, gitlab). Repository must avoid containing the words `bigblue` and `challenge`.
- Having unit tests is a strong bonus.

## The project
The current codebase sets up a simple [gRPC](https://grpc.io/) API written in Golang. It's composed of:

- a `product` service that exposes a fixed list of products
- an `order` service to be implemented to manage orders
- a `generate.sh` script to be executed in order to generate client/server code based on the proto files.
- a store used to mock persistent storage for orders and products. Read/write operations must only be done through the transactor interfaces.
> ⚠️ The store code must not be edited.
- a `server.go` entrypoint to initialize services and launch the API.

### Useful resources

- Golang tutorial: [https://tour.golang.org/welcome/1](https://tour.golang.org/welcome/1)
- Introduction to gRPC: [https://grpc.io/docs/what-is-grpc/introduction/](https://grpc.io/docs/what-is-grpc/introduction/)
- gRPC with Golang: [https://grpc.io/docs/languages/go/](https://grpc.io/docs/languages/go/)
- Protocol Buffers: [https://developers.google.com/protocol-buffers/docs/proto3](https://developers.google.com/protocol-buffers/docs/proto3)

### Setup
1. Install Golang: [https://golang.org/doc/install#download](https://golang.org/doc/install#download)
2. Install Protoc to perform code generation: [protoc](./doc/protoc.md)
3. Install Golang-specific code generation plugins: [https://grpc.io/docs/languages/go/quickstart/](https://grpc.io/docs/languages/go/quickstart/)
4. Install the Go packages of the project: run `go install` from `go.mod` directory level
5. When editing proto files, code generation can be triggered by running `generate.sh`
6. Start the API: `go run server.go`
 
> The API can be manually tested using [Insomnia](https://insomnia.rest/download): [https://support.insomnia.rest/article/188-grpc](https://support.insomnia.rest/article/188-grpc)

## Your missions
You will improve the gRPC `order` API to allow order management, as well as add an address validation system to validate orders destination before creating or updating them.

### I - Order service
The current system exposes a product service that provides a fixed list of products through an RPC.

The goal here is to improve the order service in order to create orders based on the products:

1. Complete the proto of the `product` service to implement a RPC to retrieve a single product by its ID.
2. Complete the proto of the `order` service and implement a RPC to create a new order. Order must have the following fields:
    - customer firstname
    - customer lastname
    - line items (products & quantities)
    - shipping address (destination)
        - address line (45 Rue des Petites Ecuries)
        - postal code (75010)
        - city (Paris)
        - country (FR)

### II - Address validation
As discussed previously, you should propose and implement a solution to validate the shipping address of an order before creating it:

- If the address contains some slight errors and the correct data can be identified with certainty by the system, the address will be automatically fixed and the order is created. Some examples:
    - 45 Rue des Pet**is** Ecuries → 45 Rue des Pet**ites** Ecuries
    - 1 Square Emile Z**i**la → 1 Square Emile Z**o**la
    - Par**i** → Par**is**
    - Aubervi**l**iers → Aubervi**ll**iers

-  Otherwise, if some parts of the address cannot be recognised and the system fails to validate it, the order is not created and a response with an error code is returned.


The system will be restricted to the validation of French addresses.

For this task, the French Government Address API could be helpful: [https://geo.api.gouv.fr/adresse](https://geo.api.gouv.fr/adresse).


*Good luck 🚀*
