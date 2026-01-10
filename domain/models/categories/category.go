package categories

import (
	"github.com/google/uuid"
	"github.com/tsunakit99/commandservice/errs"
)

// 商品カテゴリを保持するエンティティ
type Category struct {
	id   *CategoryId
	name *CategoryName
}

// カテゴリIDを取得する
func (ins *Category) Id() *CategoryId {
	return ins.id
}

// カテゴリ名を取得する
func (ins *Category) Name() *CategoryName {
	return ins.name
}

// 値の変更
func (ins *Category) ChangeCategoryName(name *CategoryName) {
	ins.name = name
}

// 同一性の比較
func (ins *Category) Equals(obj *Category) (bool, *errs.DomainError) {
	if obj == nil {
		return false, errs.NewDomainError("引数でnilが指定されました。")
	}
	result := ins.id.Equals(obj.Id())
	return result, nil
}

// コンストラクタ
func NewCategory(name *CategoryName) (*Category, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError(err.Error())
	} else {
		if id, err := NewCategoryId(uid.String()); err != nil {
			return nil, errs.NewDomainError(err.Error())
		} else {
			return &Category{
				id:   id,
				name: name,
			}, nil
		}
	}
}

// Categoryエンティティを再構築する
func BuildCategory(id *CategoryId, name *CategoryName) *Category {
	category := Category{
		id:   id,
		name: name,
	}
	return &category
}
