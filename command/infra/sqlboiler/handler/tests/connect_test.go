package handler_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/handler"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/handlerパッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("接続が成功した場合、nilが返る", Label("DB接続"), func() {
		absPath, _ := filepath.Abs("../../config/database.toml")
		// 環境変数に設定ファイルのパスをセット
		os.Setenv("DATABASE_TOML_PATH", absPath)
		result := handler.DBConncet()
		Expect(result).To(BeNil())
	})
})
