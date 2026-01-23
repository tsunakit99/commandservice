package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/tsunakit99/commandservice/command/domain/models/categories"
	"github.com/tsunakit99/commandservice/command/errs"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/handler"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/models"
)

// カテゴリリポジトリインターフェースの実装
type categoryRepositorySQLBoiler struct{}

// コンストラクタ
func NewcategpryRepositorySQLBoiler() categories.CategoryRepository {
	// フック関数の登録
	models.AddCategoryHook(boil.AfterInsertHook, CategoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, CategoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, CategoryAfterDeleteHook)

	return &categoryRepositorySQLBoiler{}
}

func (rep *categoryRepositorySQLBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// レコードの存在確認条件を作成
	condition := models.CategoryWhere.Name.EQ(category.Name().Value())
	// レコードの存在確認
	if exists, err := models.Categories(condition).Exists(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	} else if !exists {
		return nil
	} else {
		return errs.NewCRUDError(fmt.Sprintf("%sは既に登録されています。", category.Name().Value()))
	}
}

func (rep *categoryRepositorySQLBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// SQLBoilerのモデルに変換
	new_category := models.Category{
		ID:    0,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}
	if err := new_category.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *categoryRepositorySQLBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// 更新するレコードを取得
	up_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// 更新内容を設定
	up_model.ObjID = category.Id().Value()
	up_model.Name = category.Name().Value()
	// レコード更新
	if _, err := up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

func (rep *categoryRepositorySQLBoiler) DeleteById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// 削除するレコードを取得
	del_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return errs.NewCRUDError(fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。", category.Id().Value()))
	}
	if err != nil {
		return handler.DBErrHandler(err)
	}
	// レコード削除
	if _, err := del_model.Delete(ctx, tran); err != nil {
		return handler.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック関数
func CategoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを登録しました。\n", category.ObjID, category.Name)
	return nil
}

// 更新処理後に実行されるフック関数
func CategoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを更新しました。\n", category.ObjID, category.Name)
	return nil
}

// 削除処理後に実行されるフック関数
func CategoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを削除しました。\n", category.ObjID, category.Name)
	return nil
}
