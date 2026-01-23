package library

import (
	"fmt"
	"go_library/utils"
)

type Library struct {
	Books  map[string]*Book
	Logger utils.LogInterface
}

func (l *Library) Addbooks(books ...Book) {
	var (
		titleCount int
		totalCount int
	)
	if l.Books == nil {
		l.Books = make(map[string]*Book)
	}
	for _, book := range books {
		librarybook, exists := l.Books[book.Title]
		if !exists {
			newBook := book
			l.Books[book.Title] = &newBook
			titleCount++
			l.Logger.Log(fmt.Sprintf("成功添加书籍[%s]%d本到库存中", book.Title, book.Total))
		} else {
			librarybook.Stock += book.Total
			librarybook.Total += book.Total
			totalCount += book.Total
			l.Logger.Log(fmt.Sprintf("成功添加书籍[%s]%d本到库存中", book.Title, book.Total))
		}
		l.Logger.Log(fmt.Sprintf("共添加%v新书,总添加%v到库存中"), titleCount, totalCount)
	}
}
