package library

type Book struct {
	Title  string
	Auther string
	Stock  int
	Total  int
}

type User struct {
	Name          string
	BorrowedBooks []string
}
