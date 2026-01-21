package categories

import (
	"fmt"
	"unicode/utf8"

	"github.com/tsunakit99/commandservice/command/errs"
)

// カテゴリIDを保持する値オブジェクト
type CategoryName struct {
	value string
}

// カテゴリ名を取得する
func (ins *CategoryName) Value() string {
	return ins.value
}

// コンストラクタ
func NewCategoryName(value string) (*CategoryName, *errs.DomainError) {
	const MIN_LENGTH int = 2  // 2文字以上
	const MAX_LENGTH int = 20 // 20文字以下
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("カテゴリ名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &CategoryName{value: value}, nil
}
