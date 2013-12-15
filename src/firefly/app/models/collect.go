package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"time"
)

type Collect struct {
	Id   int64
	User int64
	Book int64

	Time time.Time
}

func FindCollect(q *qbs.Qbs, uid, bid int64) *Collect {
	col := new(Collect)
	err := q.WhereEqual("collect.User", uid).WhereEqual("collect.Book", bid).Find(col)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return col
}

func AddCollect(q *qbs.Qbs, uid, bid int64) {
	collect := new(Collect)
	collect.User = uid
	collect.Book = bid
	q.Save(collect)
}

func RemoveCollect(q *qbs.Qbs, collect *Collect) {
	q.Delete(collect)
}
