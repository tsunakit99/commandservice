package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/tsunakit99/commandservice/command/domain/models/products"
	"github.com/tsunakit99/commandservice/command/errs"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/handler"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/models"
)

type productRepositorySQLBoiler struct{}

// コンストラクタ
func NewProductRepositorySQLBoiler() products.ProductRepository {
	// フック関数の登録
	models.AddProductHook(boil.AfterInsertHook, ProductAfterInsertHook)
	models.AddProductHook(boil.AfterUpdateHook, ProductAfterUpdateHook)
	models.AddProductHook(boil.AfterDeleteHook, ProductAfterDeleteHook)

	return &productRepositorySQLBoiler{}
}

func (rep *productRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// レコードの存在確認条件を作成
	condition := models.ProductWhere.Name.EQ(product.Name().Value())
	// レコードの存在確認
	if exists, err := models.Products(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", product.Name().Value()))
	}
}

func (rep *productRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// SQLBoilerのモデルに変換
	new_product := models.Product{
		ID:         0,
		ObjID:      product.Id().Value(),
		Name:       product.Name().Value(),
		Price:      int(product.Price().Value()),
		CategoryID: product.Category().Id().Value(),
	}

	if err := new_product.Insert(ctx, tran, boil.Whitelist("obj_id", "name", "price", "category_id")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *productRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// 更新するレコードを取得
	up_model, err := models.Products(qm.Where("obj_id = ?", product.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("商品番号:%sは存在しないため、更新できませんでした。", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 更新内容を設定
	up_model.Name = product.Name().Value()
	up_model.Price = int(product.Price().Value())
	// レコード更新
	if _, err := up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name", "price")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *productRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, product *products.Product) error {
	// 削除対象を取得する
	del_model, err := models.Products(qm.Where("obj_id = ?", product.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("商品番号:%sは存在しないため、削除できませんでした。", product.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック
func ProductAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を登録しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}

// 変更処理後に実行されるフック
func ProductAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を変更しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}

// 削除処理後に実行されるフック
func ProductAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, product *models.Product) error {
	log.Printf("商品ID:%s 商品名:%s 単価:%d カテゴリ番号: %s を削除しました。\n",
		product.ObjID, product.Name, product.Price, product.CategoryID)
	return nil
}
