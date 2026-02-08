package presen

import (
	"github.com/tsunakit99/commandservice/command/application"
	"github.com/tsunakit99/commandservice/command/presen/adapter"
	"github.com/tsunakit99/commandservice/command/presen/prepare"
	"github.com/tsunakit99/commandservice/command/presen/server"
	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend,
	fx.Provide(
		adapter.NewcategoryAdapterImpl,
		adapter.NewproductAdapterImpl,
		server.NewcategoryServer,
		server.NewproductServer,
		prepare.NewCommandServer,
	),
	fx.Invoke(prepare.CommandServiceLifecycle),
)
