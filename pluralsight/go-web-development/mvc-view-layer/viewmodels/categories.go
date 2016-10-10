package viewmodels

// Categories represents the title of the page, its active value, and a group of categories
type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

// Category represents data for a category and its orientation on the page.
type Category struct {
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// GetCategories returns data us</p><p>ed to populate categories.
func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	juiceCategory := Category{
		ImageURL:      "lemon.png",
		Title:         "Juices and Mixes",
		Description:   `Explore our wide assortment of juices and mixes expected by today's lemonade stand clientelle. Now featuring a full line of organic juices that are guaranteed to be obtained from trees that have never been treated with pesticides or artificial fertilizers.`,
		IsOrientRight: false,
	}

	supplyCategory := Category{
		ImageURL: "kiwi.png",
		Title:    "Cups, Straws, and Other Supplies",
		Description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
		IsOrientRight: true,
	}

	advertiseCategory := Category{
		ImageURL:      "pineapple.png",
		Title:         "Signs and Advertising",
		Description:   "Sure, you could just wait for people to find your stand along the side of the road, but if you want to take it to the next level, our premium line of advertising supplies.",
		IsOrientRight: false,
	}

	result.Categories = []Category{
		juiceCategory,
		supplyCategory,
		advertiseCategory,
	}

	return result
}
