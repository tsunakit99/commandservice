package products

import (
	"github.com/google/uuid"
	"github.com/tsunakit99/commandservice/domain/models/categories"
	"github.com/tsunakit99/commandservice/errs"
)

// 商品エンティティ
type Product struct {
	id       *ProductId
	name     *ProductName
	price    *ProductPrice
	category *categories.Category
}

// 商品IDを取得する
func (ins *Product) Id() *ProductId {
	return ins.id
}

// 商品名を取得する
func (ins *Product) Name() *ProductName {
	return ins.name
}

// 商品価格を取得する
func (ins *Product) Price() *ProductPrice {
	return ins.price
}

// 商品カテゴリを取得する
func (ins *Product) Category() *categories.Category {
	return ins.category
}

// 値の変更
func (ins *Product) ChangeProductName(name *ProductName) {
	ins.name = name
}

func (ins *Product) ChangeProductPrice(price *ProductPrice) {
	ins.price = price
}

func (ins *Product) ChangeCategory(category *categories.Category) {
	ins.category = category
}

// 同一性の比較
func (ins *Product) Equals(obj *Product) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

// コンストラクタ
func NewProduct(name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewProductId(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &Product{
				id:       id,
				name:     name,
				price:    price,
				category: category,
			}, nil
		}
	}
}

// Productエンティティを再構築する
func BuildProduct(id *ProductId, name *ProductName, price *ProductPrice, category *categories.Category) *Product {
	product := Product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}
	return &product
}
