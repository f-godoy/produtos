package model

import(
    "github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface{
    CreateProduct(product *Product) (*Product, error)
    FetchProducts() (*[]Product, error)
}

type Product struct{
	Id int32 `json:"id" gorm:"primaryKey;autoIncrement:true" valid:"notnull"`
    Name string `json:"name" gorm:"type:varchar(20)" valid:"notnull"`
    Description string `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
    Price float32	`json:"price" gorm:"type:float" valid:"notnull"`
}

func NewProduct(name string, description string, price float32) (*Product, error){
	product := Product{
        Name: name,
        Description: description,
        Price: price,
    }

    err := product.isValid()
    if err != nil{
        return nil, err
    }

    return &product, nil
}

func (p *Product) isValid() error {
    _, err := govalidator.ValidateStruct(p)

    if err == nil{
        return err
    }

    return nil
}