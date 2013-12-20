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
	q.Save(b)
}

func (b *Book) SetCollected(c bool) {
	b.IsCollected = c
}

func (b *Book) AddCollect(q *qbs.Qbs) {
	b.Collected += 1
	q.Save(b)
}

func (b *Book) SubCollect(q *qbs.Qbs) {
	b.Collected -= 1
	if b.Collected < 0 {
		b.Collected = 0
	}
	q.Save(b)
}

func (b *Book) SubExisting(q *qbs.Qbs) {
	b.Existing -= 1
	if b.Existing < 0 {
		b.Existing = 0
	}
	q.Save(b)
}

func (b *Book) AddExisting(q *qbs.Qbs) {
	b.Existing += 1
	if b.Existing > b.Holding {
		b.Existing = b.Holding
	}
	q.Save(b)
}

func (b *Book) SetBorrow(status int) {
	switch status {
	case BOOK:
		b.IsBooked = true
	case OWN, DUE:
		b.IsOwned = true
	default:
		b.IsBooked = false
		b.IsOwned = false
	}
}
