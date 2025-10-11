package models

type Career struct {
	Title      string `boil:"title"`
	Position   string `boil:"position"`
	BaseSalary string `boil:"base_salary"`
	Details    string `boil:"Details"`
}
