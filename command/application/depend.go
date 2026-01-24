package application

import (
	"github.com/tsunakit99/commandservice/command/application/impl"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler"
	"go.uber.org/fx"
)

// アプリケーション層の依存定義
var SrvDepend = fx.Options(
	sqlboiler.RepDepend, // SQLBoilderを利用したリポジトリインターフェイス実装
	fx.Provide(
		// サービスインターフェイス実装のコンストラクタ
		impl.NewcategoryServiceImpl,
		impl.NewproductServiceImpl,
	),
)
