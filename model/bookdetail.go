package model

type BookDetail struct {
	BookName string
	Author string
	Publicer string
	Pages string
	Price string
	Score string
	Intro string
}

func (book BookDetail)String() string {
	return "书籍:" + book.BookName + " 作者:" + book.Author + " 出版社:" + book.Publicer + " 书籍页数:" + book.Pages + " 价格:" + book.Price +
		"评分:" + book.Score + "\n内容简介:" + book.Intro
}