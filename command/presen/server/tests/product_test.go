package server_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tsunakit99/commandservice/command/application"
	"github.com/tsunakit99/commandservice/command/presen/adapter"
	"github.com/tsunakit99/commandservice/command/presen/server"
	"github.com/tsunakit99/samplepb/pb"
	"go.uber.org/fx"
)

var _ = Describe("productServer構造体", Ordered, Label("メソッドのテスト"), func() {
	var srv pb.ProductCommandServer
	var product *pb.Product
	var ctx context.Context
	var container *fx.App
	// 前処理
	BeforeAll(func() {
		ctx = context.Background() // Contextの生成
		container = fx.New(
			application.SrvDepend,
			fx.Provide(
				adapter.NewproductAdapterImpl,
				server.NewproductServer,
			),
			fx.Populate(&srv),
		)
		// fxを起動し、起動時にエラーがないことを確認する
		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	// 後処理
	AfterAll(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})
	// Add()メソッドのテスト
	Context("Add()メソッドのテスト", Label("Add"), func() {
		It("商品登録が成功し、ProductUpResultが返る", func() {
			param := pb.ProductUpParam{Crud: pb.CRUD_INSERT, Id: "", Name: "消しカス君", Price: 200, CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Create(ctx, &param)
			product = result.Product
			Expect(result.Error).To(BeNil())
		})
		It("商品登録が失敗し、pb.Errorを保持したProductUpResultが返る", func() {
			param := pb.ProductUpParam{Crud: pb.CRUD_INSERT, Id: product.GetId(), Name: product.GetName(), Price: product.GetPrice(), CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Create(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: "消しカス君は既に登録されています。"}
			Expect(result.Error).To(Equal(&e))
		})
	})
	// Update()メソッドのテスト
	Context("Update()メソッドのテスト", Label("Update"), func() {
		It("商品の更新が成功し、ProductUpResultが返る", func() {
			param := pb.ProductUpParam{Crud: pb.CRUD_UPDATE, Id: product.GetId(), Name: "消しカス君", Price: 220, CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Update(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("商品の更新が失敗し、ProductUpResultが返る", func() {
			id := "ac413f22-0cf1-490a-9635-7e9ca810e545"
			param := pb.ProductUpParam{Crud: pb.CRUD_UPDATE, Id: id, Name: "消しカス君", Price: 220, CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Update(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: fmt.Sprintf("商品番号:%sは存在しないため、更新できませんでした。", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
	// Delete()メソッドのテスト
	Context("Delete()メソッドのテスト", Label("Delete"), func() {
		It("商品の削除が成功し、ProductUpResultが返る", func() {
			param := pb.ProductUpParam{Crud: pb.CRUD_DELETE, Id: product.GetId(), Name: product.GetName(), Price: product.GetPrice(), CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Delete(ctx, &param)
			Expect(result.Error).To(BeNil())
		})
		It("商品の削除が失敗し、ProductUpResultが返る", func() {
			id := "ac413f22-0cf1-490a-9635-7e9ca810e545"
			param := pb.ProductUpParam{Crud: pb.CRUD_DELETE, Id: id, Name: product.GetName(), Price: product.GetPrice(), CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0"}
			result, _ := srv.Delete(ctx, &param)
			e := pb.Error{Type: "CRUD Error", Message: fmt.Sprintf("商品番号:%sは存在しないため、削除できませんでした。", id)}
			Expect(result.Error).To(Equal(&e))
		})
	})
})
