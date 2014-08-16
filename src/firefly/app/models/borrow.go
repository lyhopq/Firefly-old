package models

import (
	"fmt"

	"github.com/coocood/qbs"

	"time"
)

const (
	NOTBORROW = iota
	BOOK      // 预借
	OWN       // 已借
	DUE       // 超期
	PRERET    // 预还
	RETURN    // 已还
)

type Borrow struct {
	Id     int64
	UserId int64 `qbs:"fk:User"`
	User   *User
	BookId int64 `qbs:"fk:Book"`
	Book   *Book
	Status int

	Created time.Time
	Updated time.Time
}

func (b *Borrow) ShowStatus() string {
	switch b.Status {
	case BOOK:
		return "预借"
	case OWN:
		return "已借"
	case DUE:
		return "超期"
	case PRERET:
		return "预还"
	case RETURN:
		return "已还"
	default:
		return ""
	}
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

func FindBorrow(q *qbs.Qbs, uid, bid int64) []*Borrow {
	var borrows []*Borrow
	condition := qbs.NewEqualCondition("user_id", uid).AndEqual("book_id", bid)
	err := q.Condition(condition).OrderByDesc("id").FindAll(&borrows)
	if err != nil {
		fmt.Println(err)
	}

	return borrows
}

func FindBorrowsByBookId(q *qbs.Qbs, bid int64) []*Borrow {
	var borrows []*Borrow
	err := q.WhereEqual("book_id", bid).FindAll(&borrows)
	if err != nil {
		fmt.Println(err)
	}

	return borrows
}

func BorrowStatus(q *qbs.Qbs, uid, bid int64) int {
	borrows := FindBorrow(q, uid, bid)
	status := NOTBORROW
	for _, bor := range borrows {
		if bor.Status == BOOK {
			status = BOOK
			break
		} else if bor.Status > BOOK && bor.Status < RETURN {
			status = OWN
			break
		}
	}

	return status
}

func AddBooking(q *qbs.Qbs, uid, bid int64) (bool, int) {
	book := FindBookById(q, bid)
	if book.Existing != 0 {
		book.Existing -= 1
	} else {
		return false, book.Existing
	}

	bor := new(Borrow)
	bor.UserId = uid
	bor.BookId = bid
	bor.Status = BOOK
	if _, err := q.Save(bor); err != nil {
		return false, book.Existing
	}

	q.Save(book)

	return true, book.Existing
}

func RemoveBooking(q *qbs.Qbs, uid, bid int64) (bool, int) {
	book := FindBookById(q, bid)
	book.Existing += 1

	bor := FindBorrow(q, uid, bid)
	if len(bor) > 0 {
		if _, err := q.Delete(bor[0]); err != nil {
			return false, book.Existing
		}

		q.Save(book)
	}

	return true, book.Existing
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

func userBorrows(q *qbs.Qbs, page int, uid int64, con *qbs.Condition) ([]*Borrow, int64) {
	if page < 1 {
		page = 1
	}
	page -= 1

	var borrows []*Borrow
	var rows int64
	rows = q.Condition(con).Count("borrow")
	err := q.Condition(con).Limit(ItemsPerPage).OrderByDesc("created").
		Offset(page * ItemsPerPage).FindAll(&borrows)
	if err != nil {
		fmt.Println(err)
	}

	return borrows, rows

}

func UserBook(q *qbs.Qbs, page int, uid int64) ([]*Borrow, int64) {
	condition := qbs.NewEqualCondition("user_id", uid).AndEqual("status", BOOK)
	return userBorrows(q, page, uid, condition)
}

func UserOwn(q *qbs.Qbs, page int, uid int64) ([]*Borrow, int64) {
	condition := qbs.NewEqualCondition("user_id", uid)
	condition1 := qbs.NewEqualCondition("status", OWN).OrEqual("status", PRERET)
	condition.AndCondition(condition1)
	return userBorrows(q, page, uid, condition)
}

func UserHis(q *qbs.Qbs, page int, uid int64) ([]*Borrow, int64) {
	condition := qbs.NewEqualCondition("user_id", uid).AndEqual("status", RETURN)
	return userBorrows(q, page, uid, condition)
}

func BorrowCount(q *qbs.Qbs, uid int64, status int) int64 {
	con := qbs.NewEqualCondition("user_id", uid).AndEqual("status", status)
	return q.OmitJoin().Condition(con).Count("borrow")
}

func (b *Borrow) SetBorrowStatus(q *qbs.Qbs, st int) error {
	b.Status = st
	_, err := q.Save(b)

	return err
}

func (b *Borrow) IsPreRet() bool {
	return b.Status == PRERET
}
