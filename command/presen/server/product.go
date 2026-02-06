package server

import (
	"context"

	"github.com/tsunakit99/commandservice/command/application/service"
	"github.com/tsunakit99/commandservice/command/presen/adapter"
	"github.com/tsunakit99/samplepb/pb"
)

type productServer struct {
	adapter adapter.ProductAdapter
	service service.ProductService
	pb.UnimplementedProductCommandServer
}

// コンストラクタ
func NewProductServer(adapter adapter.ProductAdapter, service service.ProductService) pb.ProductCommandServer {
	return &productServer{adapter: adapter, service: service}
}

// 商品の追加 pb.ProductCommandServerインターフェイスメソッドの実装
func (ins *productServer) Create(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	// DTO→エンティティ変換
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのAdd()メソッドを実行
		if err := ins.service.Add(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}

// 商品の更新 pb.ProductCommandServerインターフェイスメソッドの実装
func (ins *productServer) Update(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	// DTO→エンティティ変換
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのUpdate()メソッドを実行
		if err := ins.service.Update(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}

// 商品の削除 pb.ProductCommandServerインターフェイスメソッドの実装
func (ins *productServer) Delete(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	// DTO→エンティティ変換
	if product, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのDelete()メソッドを実行
		if err := ins.service.Delete(ctx, product); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(product), nil
	}
}
