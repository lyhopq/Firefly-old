package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

type Collect struct {
	Id     int64
	UserId int64
	User   *User
	BookId int64
	Book   *Book

	Created time.Time
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

func UserCollectCount(q *qbs.Qbs, uid int64) int64 {
	con := qbs.NewEqualCondition("user_id", uid)
	return q.OmitJoin().Condition(con).Count("collect")
}

func UserCollected(q *qbs.Qbs, page int, uid int64) ([]*Collect, int64) {
	if page < 1 {
		page = 1
	}
	page -= 1

	var collects []*Collect
	con := qbs.NewEqualCondition("user_id", uid)
	rows := q.Condition(con).Count("collect")
	err := q.OrderByDesc("id").Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&collects)
	if err != nil {
		fmt.Println(err)
	}
	return collects, rows
}

func BookCollectCount(q *qbs.Qbs, bid int64) int64 {
	con := qbs.NewEqualCondition("book_id", bid)
	return q.OmitJoin().Condition(con).Count("collect")
}
