package library

import (
	"encoding/json"
	"fmt"
	"go_library/utils"
	"io"
	"os"
)

type Library struct {
	Books  map[string]*Book
	Logger utils.LogInterface
}

func (l *Library) LogAndPrint(formate string) {
	msg := fmt.Sprintf(formate)
	fmt.Printf(msg)
	if l.Logger != nil {
		l.Logger.Log(msg)
	}
}

const LibaryFileName = "library.json"

func SaveData(data map[string]*Book) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("数据初始化失败[%v]\n", err)
	}
	err = os.WriteFile(LibaryFileName, jsonData, 0644)
	if err != nil {
		fmt.Printf("数据写入失败[%v]\n", err)
	}
	fmt.Printf("数据写入成功！")
}

func LoadData() *Library {
	file, err := os.Open(LibaryFileName)
	lib := &Library{
		Books:  make(map[string]*Book),
		Logger: &utils.FileLogger{FileName: "system.log"},
	}
	if err != nil {
		fmt.Printf("数据读取失败[%v]\n", err)
		return lib
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("数据加载失败[%v]\n", err)
	}
	err = json.Unmarshal(content, &lib.Books)
	if err != nil {
		fmt.Printf("数据初始化失败[%v]\n", err)
	}
	return lib
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
			totalCount++
			l.LogAndPrint(fmt.Sprintf("添加书籍[%s]%d本到库存中", book.Title, book.Total))
		} else {
			librarybook.Stock += book.Total
			librarybook.Total += book.Total
			totalCount += book.Total
			l.LogAndPrint(fmt.Sprintf("添加书籍[%s]%d本到库存中", book.Title, book.Total))
		}
	}
	l.LogAndPrint(fmt.Sprintf("+++共添加%v新书,总添加%v到库存中+++", titleCount, totalCount))
	SaveData(l.Books)
}

func (l *Library) BorrowBook(bookName string) {
	librarybook, exists := l.Books[bookName]
	if !exists {
		fmt.Printf("书籍不存在\n")
	}
	if librarybook.Stock <= 0 {
		fmt.Printf("书籍库存不足\n")
	}
	librarybook.Stock -= 1
	fmt.Printf("租借%s成功\n", bookName)
	SaveData(l.Books)
}

func (l *Library) ReturnBook(bookName string) {
	librarybook, exists := l.Books[bookName]
	if !exists {
		fmt.Printf("书籍不存在\n")
	}
	if librarybook.Stock+1 > librarybook.Total {
		fmt.Printf("数量超出总库存,请询问还书者")
	}
	librarybook.Stock += 1
	fmt.Printf("归还%s成功\n", bookName)
	SaveData(l.Books)
}
