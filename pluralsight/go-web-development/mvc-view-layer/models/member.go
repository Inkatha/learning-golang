package models

type Member struct {
	email     string
	id        int
	password  string
	firstName string
}

// Email - Get member's email
func (member *Member) Email() string {
	return member.email
}

// ID - Get member's Id
func (member *Member) ID() int {
	return member.id
}

// Password - Get member's Password
func (member *Member) Password() string {
	return member.password
}

// FirstName - Get member's FirstName
func (member *Member) FirstName() string {
	return member.firstName
}

// SetEmail - Sets email for a member
func (member *Member) SetEmail(value string) {
	member.email = value
}

// SetID - Sets id for a member
func (member *Member) SetID(value int) {
	member.id = value
}

// SetPassword - Sets password for a member
func (member *Member) SetPassword(value string) {
	member.password = value
}

// SetFirstName - Sets the first name of a member
func (member *Member) SetFirstName(value string) {
	member.firstName = value
}
