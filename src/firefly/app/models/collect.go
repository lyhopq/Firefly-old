package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

type Collect struct {
	Id     int64
	UserId int64
	BookId int64

	Time time.Time
}

func FindCollect(q *qbs.Qbs, uid, bid int64) *Collect {
	col := new(Collect)
	condition := qbs.NewEqualCondition("user_id", uid).AndEqual("book_id", bid)
	err := q.Condition(condition).Find(col)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return col
}

func AddCollect(q *qbs.Qbs, uid, bid int64) {
	collect := new(Collect)
	collect.UserId = uid
	collect.BookId = bid
	q.Save(collect)
}

func RemoveCollect(q *qbs.Qbs, uid, bid int64) {
	collect := FindCollect(q, uid, bid)
	if collect != nil {
		q.Delete(collect)
	}
}
