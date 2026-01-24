package impl

import (
	"context"

	"github.com/tsunakit99/commandservice/command/application/service"
	"github.com/tsunakit99/commandservice/command/domain/models/categories"
)

// CategoryServiceインターフェイスの実装
type categoryServiceImpl struct {
	rep         categories.CategoryRepository
	transaction // transaction構造体のエンベデッド
}

// コンストラクタ
func NewcategoryServiceImpl(rep categories.CategoryRepository) service.CategoryService {
	return &categoryServiceImpl{
		rep: rep,
	}
}

// カテゴリを登録する
func (ins *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {

	// トランザクションを開始して処理を実行
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションをコミットまたはロールバック
	defer func() {
		err = ins.complete(tran, err)
	}()

	// 既に同じ名前のカテゴリが存在するか確認
	if err = ins.rep.Exists(ctx, tran, category); err != nil {
		return err
	}

	// カテゴリを登録
	if err = ins.rep.Create(ctx, tran, category); err != nil {
		return err
	}
	return err
}

// カテゴリを更新する
func (ins *categoryServiceImpl) Update(ctx context.Context, category *categories.Category) error {
	// トランザクションを開始して処理を実行
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションをコミットまたはロールバック
	defer func() {
		err = ins.complete(tran, err)
	}()
	// カテゴリを更新
	if err = ins.rep.UpdateById(ctx, tran, category); err != nil {
		return err
	}
	return err
}

// カテゴリを削除する
func (ins *categoryServiceImpl) Delete(ctx context.Context, category *categories.Category) error {
	// トランザクションを開始して処理を実行
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションをコミットまたはロールバック
	defer func() {
		err = ins.complete(tran, err)
	}()
	// カテゴリを削除
	if err = ins.rep.DeleteById(ctx, tran, category); err != nil {
		return err
	}
	return err
}
