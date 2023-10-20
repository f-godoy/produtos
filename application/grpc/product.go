package grpc

import(
	"github.com/f-godoy/produtos-go/application/grpc/pb"
	"github.com/f-godoy/produtos-go/application/usecase"
	"context"
)

type ProductGrpcService struct{
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService{
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error){
	product, err := p.ProductUseCase.SaveOneProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{
			Product: nil,
		}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id: product.Id,
			Name: product.Name,
			Description: product.Description,
			Price: product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error){
	products, err := p.ProductUseCase.SearchAllProducts()

	if err != nil {
		return &pb.FindProductsResponse{
			Products: nil,
		}, err
	}

	var productsResponse []*pb.Product
	var itemProductResponse *pb.Product

	for _, v := range *products {
		itemProductResponse = &pb.Product{
			Id: v.Id,
			Name: v.Name,
			Description: v.Description,
			Price: v.Price,
		}
		productsResponse = append(productsResponse,itemProductResponse)
	}

	return &pb.FindProductsResponse{
		Products: productsResponse,		
	}, nil
}