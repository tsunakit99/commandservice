package handler

import (
	"errors"
	"log"
	"net"

	"github.com/tsunakit99/commandservice/command/errs"

	"github.com/go-sql-driver/mysql"
)

// データベースアクセスエラーのハンドリング
func DBErrHandler(err error) error {
	var opErr *net.OpError
	var driverErr *mysql.MySQLError
	if errors.As(err, &opErr) { // 接続がタイムアウトかネットワーク関連のエラーの場合
		log.Println(err.Error())
		return errs.NewInternalError(opErr.Error())
	} else if errors.As(err, &driverErr) { // MySQLのドライバエラーの場合
		log.Printf("Code:%d Message:%s\n", driverErr.Number, driverErr.Message)
		if driverErr.Number == 1062 { // 一意制約違反
			return errs.NewCRUDError("一意制約違反です。")
		} else {
			return errs.NewInternalError(driverErr.Message)
		}
	} else { // その他のエラー
		log.Println(err.Error())
		return errs.NewInternalError(err.Error())
	}
}
