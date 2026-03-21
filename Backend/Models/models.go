package models

type Career struct {
	Title      string `boil:"title"`
	Position   string `boil:"position"`
	BaseSalary string `boil:"base_salary"`
	Details    string `boil:"details"`
}

type LoginCredentials struct {
	UserId   string `boil:"user_id"`
	UserName string `boil:"user_name"`
	Password string `boil:"password"`
}

type Dashboard struct {
	UserName string `boil:"user_name"`
	FileName string `boil:"file_name"`
	FilePath string `boil:"file_path"`
}
