package categories

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	// 同名のカテゴリが存在するか確認する
	Exists(ctx context.Context, tran *sql.Tx, category *Category) error
	// 新しいカテゴリを永続化する
	Create(ctx context.Context, tran *sql.Tx, category *Category) error
	// カテゴリを変更する
	UpedateById(ctx context.Context, tran *sql.Tx, category *Category) error
	// カテゴリを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, category *Category) error
}
