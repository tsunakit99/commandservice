package categories_tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 商品カテゴリエンティティ、値オブジェクトテスト用エントリポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                        // テストが失敗したときに呼ばれるハンドラを登録
	RunSpecs(t, "domain/models/categoriesパッケージのテスト") // 登録されたすべてのテストを実行
}
