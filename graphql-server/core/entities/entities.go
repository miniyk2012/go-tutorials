package entities

type Movie struct {
	Title     string
	Year      int
	Actors    []Cast
	Directors []Cast
}

type Cast struct {
	Name   string // Both first name and last name or other formats if not English
	Gender bool   // true for male and false for female
}

type User struct {
	UID string
	AuthorizedActions []string
}
