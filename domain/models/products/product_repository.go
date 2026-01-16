package products

import (
	"context"
	"database/sql"
)

type ProductRepository interface {
	// 同名の商品が存在するか確認する
	Exists(ctx context.Context, tran *sql.Tx, product *Product) error
	// 新しい商品を永続化する
	Create(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を変更する
	UpdateById(ctx context.Context, tran *sql.Tx, product *Product) error
	// 商品を削除する
	DeleteById(ctx context.Context, tran *sql.Tx, product *Product) error
}
