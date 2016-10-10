package viewmodels

// Home represents a title and active page name
type Home struct {
	Title  string
	Active string
}

// GetHome sets the title and current active value for the homepage
func GetHome() Home {
	result := Home{
		Title:  "Lemonade Stand Society",
		Active: "home",
	}

	return result
}
