package commands

type DeleteProduct struct {
	productsService productUploader
	logger          logger
	Options         struct {
		ProductName    string `short:"n"  long:"product-name" description:"name of the product being configured"`
		ProductVersion string `short:"v"  long:"product-version" description:"vesrion of the product being configured"`
	}
}

func NewDeleteProduct(productUploader productUploader, logger logger) DeleteProduct {
	return DeleteProduct{
		productsService: productUploader,
		logger:          logger,
	}
}

func (dup DeleteProduct) Usage() Usage {
	return Usage{
		Description:      "This command deletes products in the targeted Ops Manager",
		ShortDescription: "deletes products on the Ops Manager targeted",
	}
}

func (dup DeleteProduct) Execute(args []string) error {
	dup.logger.Printf("trashing products")

	dup.logger.Printf("done")

	return nil
}
