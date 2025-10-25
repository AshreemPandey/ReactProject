package models

type Career struct {
	Title      string `boil:"title"`
	Position   string `boil:"position"`
	BaseSalary string `boil:"base_salary"`
	Details    string `boil:"details"`
}

type LoginCredentials struct {
	Username string `boil:"username"`
	Password string `boil:"password"`
}

type Dashboard struct {
	Username string `boil:"username"`
	FileName string `boil:"filename"`
	Filepath string `boil:"filepath"`
}
