package book

type Author struct {
	Id         int    `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
}

type Book struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Author      int    `db:"author"`
}
