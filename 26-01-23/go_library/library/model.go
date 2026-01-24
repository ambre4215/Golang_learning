package library

type Book struct {
	Title  string
	Author string
	Stock  int
	Total  int
}

type User struct {
	Name          string
	BorrowedBooks []string
}
