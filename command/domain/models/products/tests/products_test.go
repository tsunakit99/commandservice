package products_tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 商品エンティティ、値オブジェクトテスト用エントリポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                      // テストが失敗したときに呼ばれるハンドラを登録
	RunSpecs(t, "domain/models/productsパッケージのテスト") // 登録されたすべてのテストを実行
}
