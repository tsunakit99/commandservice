package impl

import (
	"context"

	"github.com/tsunakit99/commandservice/command/application/service"
	"github.com/tsunakit99/commandservice/command/domain/models/products"
)

type productServiceImpl struct {
	rep         products.ProductRepository
	transaction // transaction構造体のエンベデッド
}

func NewproductServiceImpl(rep products.ProductRepository) service.ProductService {
	return &productServiceImpl{
		rep: rep,
	}
}

func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
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
	if err = ins.rep.Exists(ctx, tran, product); err != nil {
		return err
	}
	// カテゴリを登録
	if err = ins.rep.Create(ctx, tran, product); err != nil {
		return err
	}
	return err
}

func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	// トランザクションを開始して処理を実行
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションをコミットまたはロールバック
	defer func() {
		err = ins.complete(tran, err)
	}()

	// 商品を更新
	if err = ins.rep.UpdateById(ctx, tran, product); err != nil {
		return err
	}
	return err
}

// 商品を削除
func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	// トランザクションを開始して処理を実行
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションをコミットまたはロールバック
	defer func() {
		err = ins.complete(tran, err)
	}()

	// 商品を削除
	if err = ins.rep.DeleteById(ctx, tran, product); err != nil {
		return err
	}
	return err
}
