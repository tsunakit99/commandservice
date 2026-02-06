package server

import (
	"context"

	"github.com/tsunakit99/commandservice/command/application/service"
	"github.com/tsunakit99/commandservice/command/presen/adapter"
	"github.com/tsunakit99/samplepb/pb"
)

// カテゴリ更新サーバの実装
type categoryServer struct {
	adapter adapter.CategoryAdapter
	service service.CategoryService
	pb.UnimplementedCategoryCommandServer
}

// コンストラクタ
func NewcategoryServer(adapter adapter.CategoryAdapter, service service.CategoryService) pb.CategoryCommandServer {
	return &categoryServer{adapter: adapter, service: service}
}

// カテゴリの追加 pb.CategoryCommandServerインターフェイスメソッドの実装
func (ins *categoryServer) Create(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// DTO→エンティティ変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのAdd()メソッドを実行
		if err := ins.service.Add(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

// カテゴリの更新 pb.CategoryCommandServerインターフェイスメソッドの実装
func (ins *categoryServer) Update(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// DTO→エンティティ変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのUpdate()メソッドを実行
		if err := ins.service.Update(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}

// カテゴリの削除 pb.CategoryCommandServerインターフェイスメソッドの実装
func (ins *categoryServer) Delete(ctx context.Context, param *pb.CategoryUpParam) (*pb.CategoryUpResult, error) {
	// DTO→エンティティ変換
	if category, err := ins.adapter.ToEntity(param); err != nil {
		return ins.adapter.ToResult(err), nil
	} else {
		// サービスのDelete()メソッドを実行
		if err := ins.service.Delete(ctx, category); err != nil {
			return ins.adapter.ToResult(err), nil
		}
		return ins.adapter.ToResult(category), nil
	}
}
