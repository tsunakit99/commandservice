package service

import (
	"context"

	"github.com/tsunakit99/commandservice/command/domain/models/categories"
)

// カテゴリ更新サービスインターフェース
type CategoryService interface {
	Add(ctx context.Context, category *categories.Category) error
	Update(ctx context.Context, category *categories.Category) error
	Delete(ctx context.Context, category *categories.Category) error
}
