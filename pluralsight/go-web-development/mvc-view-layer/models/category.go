package models

// Category represents data about a product category
type Category struct {
	imageURL      string
	title         string
	description   string
	isOrientRight bool
	id            int
}

// ImageURL - Gets Image URL
func (category *Category) ImageURL() string {
	return category.imageURL
}

// Title - Gets Titles
func (category *Category) Title() string {
	return category.title
}

// Description - Gets Descriptions
func (category *Category) Description() string {
	return category.description
}

// IsOrientRight - Gets Orientation
func (category *Category) IsOrientRight() bool {
	return category.isOrientRight
}

// ID - Gets ID
func (category *Category) ID() int {
	return category.id
}

// SetImageURL - Sets Image URL
func (category *Category) SetImageURL(value string) {
	category.imageURL = value
}

// SetTitle - Sets the Title
func (category *Category) SetTitle(value string) {
	category.title = value
}

// SetDescription - Set Description
func (category *Category) SetDescription(value string) {
	category.description = value
}

// SetIsOrientRight - Set Orientation
func (category *Category) SetIsOrientRight(value bool) {
	category.isOrientRight = value
}

// SetID - Sets ID's
func (category *Category) SetID(value int) {
	category.id = value
}

// GetCategories - Seed Data for categories
// TODO: Retrieve this data from a database
func GetCategories() []Category {
	result := []Category{
		Category{
			id:            1,
			imageURL:      "lemon.png",
			title:         "Juices and Mixes",
			description:   `Explore our wide assortment of juices and mixes expected by today's lemonade stand clientelle. Now featuring a full line of organic juices that are guaranteed to be obtained from trees that have never been treated with pesticides or artificial fertilizers.`,
			isOrientRight: false,
		},

		Category{
			id:       2,
			imageURL: "kiwi.png",
			title:    "Cups, Straws, and Other Supplies",
			description: `From paper cups to bio-degradable plastic to straws and
                            napkins, LSS is your source for the sundries that keep your stand
                            running smoothly.`,
			isOrientRight: true,
		},

		Category{
			id:            3,
			imageURL:      "pineapple.png",
			title:         "Signs and Advertising",
			description:   "Sure, you could just wait for people to find your stand along the side of the road, but if you want to take it to the next level, our premium line of advertising supplies.",
			isOrientRight: false,
		},
	}
	return result
}
