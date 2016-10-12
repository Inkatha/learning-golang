package viewmodels

// Home represents a title and active value for the homepage
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

// Login represents a title and active value for the login page
type Login struct {
	Title  string
	Active string
}

// GetLogin sets the title and current active link for the login page
func GetLogin() Login {
	result := Login{
		Title:  "Lemonade Stand Society - Login",
		Active: "",
	}

	return result
}
