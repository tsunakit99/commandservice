package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type DBConfig struct {
	Dbname string `toml:"dbname"`
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

// database.tomlから接続情報を取得してDbConfig型で返す
func tomlRead() (*DBConfig, error) {
	// 環境変数から設定ファイルのパスを取得
	path := os.Getenv("DATABASE_TOML_PATH")
	if path == "" {
		path = "infra/sqlboiler/config/database.toml"
	}
	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}
	config := m["mysql"]
	return &config, nil
}

// データベース接続
func DBConncet() error {
	config, err := tomlRead()
	if err != nil {
		return DBErrHandler(err)
	}

	// 接続文字列生成
	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)

	// DB接続
	conn, err := sql.Open(rdbms, connect_str)
	if err != nil {
		return DBErrHandler(err)
	}

	// 接続確認
	if err = conn.Ping(); err != nil {
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10                   // 初期接続数
	MAX_OPEN_CONNS := 100                  // 最大接続数
	CONN_MAX_LIFETIME := 300 * time.Second // 接続の最大寿命

	// コネクションプールの設定
	conn.SetMaxIdleConns(MAX_IDLE_CONNS)
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)
	conn.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	boil.SetDB(conn)      // sqlboilerにDB接続を設定
	boil.DebugMode = true // SQLログを出力する

	return nil
}
