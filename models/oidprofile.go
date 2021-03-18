package models

// IODProfile is a model struct that contains data to be gotten by an oid action
type IODProfile = struct {
	Aud           string
	Email         string
	EmailVerified bool
	Exp           string
	Iss           string
	Locale        string
	Name          string
}
