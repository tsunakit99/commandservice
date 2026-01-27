package adapter

import (
	"github.com/tsunakit99/commandservice/command/domain/models/categories"
	"github.com/tsunakit99/samplepb/pb"
)

// パラメータと実行結果の変換インターフェース
type CategoryAdapter interface {
	// CategoryUpParamからCategoryに変換
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error)
	// 実行結果からCategoryUpResultに変換
	ToResult(result any) *pb.CategoryUpResult
}
