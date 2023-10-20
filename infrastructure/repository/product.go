package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/f-godoy/produtos-go/domain/model"
)

type ProductRepositoryDb struct{
	Db *gorm.DB
}

func (r ProductRepositoryDb) CreateProduct(product *model.Product) (*model.Product, error){
	err := r.Db.Create(product).Error

	if err != nil{
		return nil, err
	}
	return product, nil
}

func (r ProductRepositoryDb) FetchProducts() (*[]model.Product, error){
	var products []model.Product
	err := r.Db.Find(&products).Error

	if err !=nil {
		return nil, err
	}

	return &products,nil
}
