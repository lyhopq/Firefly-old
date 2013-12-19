package models

import (
	"fmt"
	"github.com/coocood/qbs"

	"time"
)

const (
	BOOK   = iota // 预借
	OWN           // 已借
	DUE           // 超期
	PRERET        // 预还
	RETURN        // 已还
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

func FindBorrowById(q *qbs.Qbs, id int64) *Borrow {
	bor := new(Borrow)
	err := q.WhereEqual("borrow.id", id).Find(bor)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bor
}

func FindBorrow(q *qbs.Qbs, uid, bid int64) *Borrow {
	bor := new(Borrow)
	condition := qbs.NewEqualCondition("user_id", uid).AndEqual("book_id", bid)
	err := q.Condition(condition).Find(bor)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return bor
}

func FindBorrowsByBookId(q *qbs.Qbs, bid int64) []*Borrow {
	var borrows []*Borrow
	err := q.WhereEqual("book_id", bid).FindAll(&borrows)
	if err != nil {
		fmt.Println(err)
	}

	return borrows
}

func AddBooking(q *qbs.Qbs, uid, bid int64) bool {
	book := FindBookById(q, bid)
	if book.Existing != 0 {
		book.Existing -= 1
	} else {
		return false
	}

	bor := new(Borrow)
	bor.UserId = uid
	bor.BookId = bid
	bor.Status = BOOK
	if _, err := q.Save(bor); err != nil {
		return false
	}

	q.Save(book)

	return true
}

func RemoveBooking(q *qbs.Qbs, uid, bid int64) bool {
	book := FindBookById(q, bid)
	book.Existing += 1

	bor := FindBorrow(q, uid, bid)
	if bor != nil {
		if _, err := q.Delete(bor); err != nil {
			return false
		}

		q.Save(book)
	}

	return true
}

func GetBorrows(q *qbs.Qbs, page int, column string, value interface{}, order string) ([]*Borrow, int64) {
	if page < 1 {
		page = 1
	}
	page -= 1

	var borrows []*Borrow
	var rows int64
	if column == "" {
		rows = q.Count("borrow")
		err := q.OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&borrows)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		rows = q.WhereEqual(column, value).Count("borrow")
		err := q.WhereEqual(column, value).OrderByDesc(order).
			Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&borrows)
		if err != nil {
			fmt.Println(err)
		}
	}

	return borrows, rows
}
