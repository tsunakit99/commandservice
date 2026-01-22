package products

import (
	"fmt"

	"github.com/tsunakit99/commandservice/command/errs"
)

// 商品価格を保持する値オブジェクト
type ProductPrice struct {
	value uint32
}

// 商品価格を取得する
func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

// コンストラクタ
func NewProductPrice(value uint32) (*ProductPrice, *errs.DomainError) {
	const MIN_VALUE = 50    // 50円以上
	const MAX_VALUE = 10000 // 10,000円以下
	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(fmt.Sprintf("単価は%d以上、%d以下です。", MIN_VALUE, MAX_VALUE))
	}
	return &ProductPrice{value: value}, nil
}
