package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"strings"
	"time"
)

type Book struct {
	Id              int64
	Title           string `qbs:"size:128,notnull"`
	Cover           string `qbs:"size:128"`
	Author          string `qbs:"size:64,notnull"`
	Translator      string `qbs:"size:64"`
	Pages           int    `qbs:"size:32"`
	Introduction    string
	Publisher       string `qbs:"size:64"`
	Language        string `qbs:"size:16"`
	PublicationDate time.Time
	Isbn            string `qbs:"size:16,unique,notnull,index"`

	ShelfTime time.Time `qbs:"created"`
	Holding   int
	Existing  int
	Hited     int64
	Readed    int64
	Commented int64
	Collected int64

	IsCollected bool `qbs:"-"`
	IsBooked    bool `qbs:"-"`
	IsOwned     bool `qbs:"-"`
}

func Recommend(q *qbs.Qbs, column string) []*Book {
	books, _ := GetBooks(q, 1, "", "", column)
	return books
}

func FindBookById(q *qbs.Qbs, id int64) *Book {
	book := new(Book)
	err := q.WhereEqual("book.id", id).Find(book)
	if err != nil {
		fmt.Println(err)
	}

	return book
}

func GetPagination(page int, rows int64, url string) *Pagination {
	if page < 1 {
		page = 1
	}
	url = url[:strings.Index(url, "=")+1]
	return NewPagination(page, int(rows), url)
}

func GetBooks(q *qbs.Qbs, page int, column string, value interface{}, order string) ([]*Book, int64) {
	if page < 1 {
		page = 1
	}
	page -= 1

	var books []*Book
	var rows int64
	if column == "" {
		rows = q.Count("book")
		err := q.OmitFields("Introduction").OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&books)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		rows = q.WhereEqual(column, value).Count("book")
		err := q.WhereEqual(column, value).
			OmitFields("Introduction").OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&books)
		if err != nil {
			fmt.Println(err)
		}
	}

	return books, rows
}

func (b *Book) CoverImgSrc() string {
	if strings.HasPrefix(b.Cover, "thumb") {
		return fmt.Sprintf("/public/upload/%s", b.Cover)
	}
	return fmt.Sprintf("/public/img/%s", b.Cover)
}

func (b *Book) AddHited(q *qbs.Qbs) {
	b.Hited += 1
	type Book struct {
		Hited int64
	}
	book := new(Book)
	book.Hited = b.Hited
	q.WhereEqual("id", b.Id).Update(book)
}

func (b *Book) SetCollected(c bool) {
	b.IsCollected = c
}

func (b *Book) AddCollect(q *qbs.Qbs) {
	b.Collected += 1
	type Book struct {
		Collected int64
	}
	book := new(Book)
	book.Collected = b.Collected
	q.WhereEqual("id", b.Id).Update(book)
}

func (b *Book) SubCollect(q *qbs.Qbs) {
	b.Collected -= 1
	if b.Collected < 0 {
		b.Collected = 0
	}
	type Book struct {
		Collected int64
	}
	book := new(Book)
	book.Collected = b.Collected
	q.WhereEqual("id", b.Id).Update(book)
}

func (b *Book) SetBorrow(status int) {
	switch status {
	case BOOK:
		b.IsBooked = true
	default:
		b.IsOwned = true
	}
}
