package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/handler"
)

// トランザクション制御
type transaction struct{}

// トランザクションを開始する
func (inc *transaction) begin(ctx context.Context) (*sql.Tx, error) {
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, handler.DBErrHandler(err)
	}
	return tran, nil
}

// トランザクションを終了する
func (inc *transaction) complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをロールバックしました。")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをコミットしました。")
		}
	}
	return nil
}
