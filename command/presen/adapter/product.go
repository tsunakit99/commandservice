package adapter

import (
	"github.com/tsunakit99/commandservice/command/domain/models/products"
	"github.com/tsunakit99/samplepb/pb"
)

type ProductAdapter interface {
	// ProductUpParamからProductに変換
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	// 実行結果からProductUpResultに変換
	ToResult(result any) *pb.ProductUpResult
}
