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
	ID            int
}

// GetCategories returns data us</p><p>ed to populate categories.
func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	return result
}
