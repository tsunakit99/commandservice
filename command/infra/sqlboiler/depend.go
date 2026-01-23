package sqlboiler

import (
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/repository"
	"go.uber.org/fx"
)

// SQLBoilerを使ったリポジトリ実装を依存注入するためのパッケージ
var RepDepend = fx.Options(
	fx.Provide(
		repository.NewproductRepositorySQLBoiler,
		repository.NewcategpryRepositorySQLBoiler,
	),
)
