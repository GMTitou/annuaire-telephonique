package annuaire

type Contact struct {
	ID        int    `json:"id"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	phone     string `json:"phone"`
}
