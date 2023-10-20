package usecase

import(
	"github.com/f-godoy/produtos-go/domain/model"
)

type ProductUseCase struct{
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) SaveOneProduct(name string, description string, price float32) (*model.Product, error){
	newProduct, err := model.NewProduct(name,description,price)
	
	if err != nil{
		return nil,err
	}

	productInserted, err := p.ProductRepository.CreateProduct(newProduct)

	if err != nil{
		return nil,err
	}

	return productInserted, nil
}

func (p *ProductUseCase) SearchAllProducts() (*[]model.Product, error){
	products, err := p.ProductRepository.FetchProducts()

	if err != nil{
		return nil, err
	}

	return products, nil
}