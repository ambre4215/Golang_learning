package main

import (
	"fmt"
	"go_library/library"
	"strconv"
	"strings"
)

func main() {
	lib := library.LoadData()
	for {
		fmt.Printf("请输入需要的操作\n")
		fmt.Printf("[1] 批量上架\n")
		fmt.Printf("[2] 借书\n")
		fmt.Printf("[3] 还书\n")
		fmt.Printf("[4] 查看所有书籍\n")
		var inputOp int
		fmt.Scan(&inputOp)
		switch inputOp {
		case 1:
			var inputbooks string
			var batchbooks []library.Book
			fmt.Scan(&inputbooks)
			lines := strings.Split(inputbooks, "#")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				part := strings.Split(line, ",")
				if len(part) == 3 {
					title := part[0]
					author := part[1]
					count, _ := strconv.Atoi(part[2])

					b := library.Book{
						Title:  title,
						Author: author,
						Stock:  count,
						Total:  count,
					}
					batchbooks = append(batchbooks, b)
				}
			}
			lib.Addbooks(batchbooks...)
		case 2:
			var inputTitle string
			fmt.Scan(&inputTitle)
			lib.BorrowBook(inputTitle)
		case 3:
			var inputTitle string
			fmt.Scan(&inputTitle)
			lib.ReturnBook(inputTitle)
		case 4:
			for _, l := range lib.Books {
				fmt.Printf("标题[%s]作者[%s]库存[%d]总数[%d]\n", l.Title, l.Author, l.Stock, l.Total)
			}
		default:
			fmt.Printf("无效输入！")
		}
	}
}
