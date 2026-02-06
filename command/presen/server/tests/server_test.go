package server_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tsunakit99/commandservice/command/infra/sqlboiler/handler"
)

func TestHelperPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/serverパッケージのテスト")
}

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../../infra/sqlboiler/config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := handler.DBConncet()
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのでテストを中止します。")
})
