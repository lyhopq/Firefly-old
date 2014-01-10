package models

import (
	"fmt"
	"github.com/coocood/qbs"
	"github.com/robfig/revel"
	"time"
)

type Notice struct {
	Id      int64
	Title   string `qbs:"size:128,notnull"`
	Content string `qbs:"notnull"`
	Created time.Time
}

func LastNotice(q *qbs.Qbs) *Notice {
	notice := new(Notice)
	err := q.OrderByDesc("id").Limit(1).Find(notice)
	if err != nil {
		fmt.Println(err)
	}

	return notice
}

func GetNotices(q *qbs.Qbs, page int) ([]*Notice, int64) {
	if page < 1 {
		page = 1
	}
	page -= 1

	var notices []*Notice
	rows := q.Count("notice")
	err := q.OrderByDesc("id").Limit(ItemsPerPage).Offset(page * ItemsPerPage).FindAll(&notices)
	if err != nil {
		fmt.Println(err)
	}

	return notices, rows
}

func FindNoticeById(q *qbs.Qbs, id int64) *Notice {
	notice := new(Notice)
	err := q.WhereEqual("id", id).Find(notice)
	if err != nil {
		fmt.Println(err)
	}

	return notice
}

func (n *Notice) Save(q *qbs.Qbs) bool {
	_, err := q.Save(n)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (n *Notice) Validate(q *qbs.Qbs, v *revel.Validation) {
	v.Required(n.Title).Message("请输入标题")
	v.Required(n.Content).Message("请输入内容")
}
