# To execute this script on Windows, you can either:
# - typing each command on the cmd
# - saving this file as a .cmd and running it as .\generate.cmd
protoc --go_out=product --go-grpc_out=product product/rpc/product.proto
protoc --go_out=order --go-grpc_out=order order/rpc/order.proto