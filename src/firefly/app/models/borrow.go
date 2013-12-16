package models

import (
	"fmt"
	"github.com/coocood/qbs"

	"time"
)

const (
	BOOK = iota
	OWN
	DUE
)

type Borrow struct {
	Id     int64
	UserId int64 `qbs:"fk:User"`
	User   *User
	BookId int64 `qbs:"fk:Book"`

	Book   *Book
	Status int

	Updated time.Time
}

func FindBorrow(q *qbs.Qbs, uid, bid int64) *Borrow {
	bor := new(Borrow)
	err := q.WhereEqual("user_id", uid).WhereEqual("book_id", bid).Find(bor)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bor
}

func AddBooking(q *qbs.Qbs, uid, bid int64) (ok bool) {
	bor := new(Borrow)
	bor.UserId = uid
	bor.BookId = bid
	bor.Status = BOOK
	ok = true
	if _, err := q.Save(bor); err != nil {
		ok = false
	}

	return
}

func RemoveBooking(q *qbs.Qbs, uid, bid int64) (ok bool) {
	bor := FindBorrow(q, uid, bid)
	ok = true
	if bor == nil {
		ok = false
		return
	}

	if _, err := q.Delete(bor); err != nil {
		ok = false
	}

	return
}
