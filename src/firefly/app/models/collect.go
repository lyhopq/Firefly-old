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
