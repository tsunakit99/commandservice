package adapter

import (
	"github.com/tsunakit99/commandservice/command/domain/models/categories"
	"github.com/tsunakit99/commandservice/command/errs"
	"github.com/tsunakit99/samplepb/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryAdapterImpl struct{}

func NewCategoryAdapterImpl() CategoryAdapter {
	return &categoryAdapterImpl{}
}

// CategoryUpParam からCategoryに変換
func (ins *categoryAdapterImpl) ToEntity(param *pb.CategoryUpParam) (*categories.Category, error) {
	// コマンド種別別のエンティティ生成
	switch param.GetCrud() {
	case pb.CRUD_INSERT:
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		category, err := categories.NewCategory(name)
		if err != nil {
			return nil, err
		}
		return category, nil
	case pb.CRUD_UPDATE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		name, err := categories.NewCategoryName(param.GetName())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, name), nil
	case pb.CRUD_DELETE:
		id, err := categories.NewCategoryId(param.GetId())
		if err != nil {
			return nil, err
		}
		return categories.BuildCategory(id, nil), nil
	default:
		return nil, errs.NewDomainError("不明な操作を受信しました。")
	}
}

// 実行結果からCategoryUpResultに変換
func (ins *categoryAdapterImpl) ToResult(result any) *pb.CategoryUpResult {
	var up_category *pb.Category
	var up_err *pb.Error
	switch v := result.(type) {
	case *categories.Category: // 実行結果がCategoryエンティティの場合
		if v.Name() == nil {
			up_category = &pb.Category{Id: v.Id().Value(), Name: ""}
		} else {
			up_category = &pb.Category{Id: v.Id().Value(), Name: v.Name().Value()}
		}
	case *errs.DomainError: // 実行結果がDomainErrorの場合
		up_err = &pb.Error{Type: "Domain Error", Message: v.Error()}
	case *errs.CRUDError: // 実行結果がCRUDErrorの場合
		up_err = &pb.Error{Type: "CRUD Error", Message: v.Error()}
	case *errs.InternalError: // 実行結果がInternalErrorの場合
		up_err = &pb.Error{Type: "Internal Error", Message: "内部エラーが発生しました。"}
	}
	return &pb.CategoryUpResult{Category: up_category, Error: up_err, Timestamp: timestamppb.Now()}
}
