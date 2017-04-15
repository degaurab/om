package commands_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
)

var _ = Describe("DeleteProduct", func() {
	var (
		productsService *fakes.ProductUploader
		logger          *fakes.Logger
	)

	BeforeEach(func() {
		productsService = &fakes.ProductUploader{}
		logger = &fakes.Logger{}
	})

	FIt("deletes a product", func() {
		client := commands.NewDeleteProduct(productsService, logger)

		err := client.Execute([]string{
			"--product-name", "cf",
			"--product-version", "1.2.3",
		})
		Expect(err).NotTo(HaveOccurred())

		// Expect(productsService.StagedProductsCallCount()).To(Equal(1))
		// Expect(productsService.DeleteProductArgsForCall(0)).To(Equal(api.ProductsConfigurationInput{
		// 	GUID: "some-product-guid",
		// }))

		err = client.Execute([]string{
			"--product-name", "cf",
			"--product-version", "1.2.3",
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(productsService.TrashCallCount()).To(Equal(1))

		format, v := logger.PrintfArgsForCall(0)
		Expect(fmt.Sprintf(format, v...)).To(Equal("trashing products"))

		format, v = logger.PrintfArgsForCall(1)
		Expect(fmt.Sprintf(format, v...)).To(Equal("done"))
	})

	Context("failure cases", func() {
		Context("when the trash call returns an error", func() {
			It("returns an error", func() {
				productsService.TrashReturns(errors.New("some error"))
				command := commands.NewDeleteProduct(productsService, logger)
				err := command.Execute([]string{})
				Expect(err).To(MatchError("could not delete products: some error"))
			})
		})
	})
})
