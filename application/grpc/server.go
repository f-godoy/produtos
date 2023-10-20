package grpc

import(
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/f-godoy/produtos-go/infrastructure/repository"
	"github.com/f-godoy/produtos-go/application/usecase"
	"github.com/f-godoy/produtos-go/application/grpc/pb"
)

func StartGrpcServer(database *gorm.DB, port int){
	grpcServer := grpc.NewServer()
	// Para debugar
	reflection.Register(grpcServer)

	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}
	productGrpcService := NewProductGrpcService(productUseCase)
	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d",port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	} 

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil{
		log.Fatal("cannot start grpc server", err)
	}
}