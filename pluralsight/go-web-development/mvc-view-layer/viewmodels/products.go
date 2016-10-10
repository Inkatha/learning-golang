package viewmodels

// Products represents the title of the product page, the active link,
// and maintains a list of product data.
type Products struct {
	Title   string
	Active  string
	Product []Product
}

// Product contains data about a particular product.
type Product struct {
	ImageURL      string
	Name          string
	Description   string
	Price         string
	UnitOfMeasure string
}

// GetProducts returns data from a list of products.
func GetProducts() Products {
	result := Products{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}
	lemonJuiceProduct := Product{
		ImageURL:      "lemon.png",
		Name:          "Lemon Juice",
		Description:   "Made from fresh, organic California lemons.",
		Price:         "$1.09",
		UnitOfMeasure: "liter",
	}

	appleJuiceProduct := Product{
		ImageURL:      "apple.png",
		Name:          "Apple Juice",
		Description:   "The perfect blend of Washington apples.",
		Price:         "$0.89",
		UnitOfMeasure: "liter",
	}

	watermelonJuiceProduct := Product{
		ImageURL:      "watermelon.png",
		Name:          "Watermelon Juice",
		Description:   "From sun-drenched fields in Florida.",
		Price:         "$3.99",
		UnitOfMeasure: "liter",
	}

	kiwiJuiceProduct := Product{
		ImageURL:      "kiwi.png",
		Name:          "Kiwi Juice",
		Description:   "California sunshine and rain distilled into sweet goodness",
		Price:         "$5.99",
		UnitOfMeasure: "liter",
	}

	mangosteenJuiceProduct := Product{
		ImageURL:      "mangosteen.png",
		Name:          "Mangosteen Juice",
		Description:   "Our latest taste sensation, imported directly from Hawaii",
		Price:         "$6.87",
		UnitOfMeasure: "liter",
	}

	orangeJuiceProduct := Product{
		ImageURL:      "orange.png",
		Name:          "Orange Juice",
		Description:   "Fresh squeezed from Florida's best oranges.",
		Price:         "$1.67",
		UnitOfMeasure: "liter",
	}

	pineappleJuiceProduct := Product{
		ImageURL:      "pineapple.png",
		Name:          "Pineapple Juice",
		Description:   "An exotic and refreshing offering. Straight from Hawaii.",
		Price:         "$2.48",
		UnitOfMeasure: "liter",
	}

	strawberryJuiceProduct := Product{
		ImageURL:      "strawberry.png",
		Name:          "Strawberry Juice",
		Description:   "The perfect balance of sweet and tart.",
		Price:         "$4.36",
		UnitOfMeasure: "liter",
	}

	result.Product = []Product{
		lemonJuiceProduct,
		appleJuiceProduct,
		watermelonJuiceProduct,
		kiwiJuiceProduct,
		mangosteenJuiceProduct,
		orangeJuiceProduct,
		pineappleJuiceProduct,
		strawberryJuiceProduct,
	}

	return result
}
