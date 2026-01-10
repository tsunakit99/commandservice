package categories

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/tsunakit99/commandservice/errs"
)

// カテゴリIDを保持する値オブジェクト(UUID)
type CategoryId struct {
	value string
}

// カテゴリIDを取得する
func (ins *CategoryId) Value() string {
	return ins.value
}

// 同一性の比較
func (ins *CategoryId) Equals(value *CategoryId) bool {
	if ins == value {
		return true
	}
	return ins.value == value.Value()
}

// カテゴリIDを生成するコンストラクタ
func NewCategoryId(value string) (*CategoryId, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("カテゴリIDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &CategoryId{value: value}, nil
}
