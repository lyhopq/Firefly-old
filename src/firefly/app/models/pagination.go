package models

import "fmt"

const (
	ItemsPerPage = 10 //每页几条记录
)

type Pagination struct {
	page  int //当前页码（从1开始）
	rows  int //记录总数
	url   string
	pages int //总页数
}

func NewPagination(page int, rows int, url string) *Pagination {
	p := Pagination{}
	if page < 1 {
		page = 1
	}
	p.page = page
	p.rows = rows
	p.url = url
	return &p
}

func (p *Pagination) CurPage() int {
	return p.page
}

func (p *Pagination) PrePage() string {
	page := p.page
	page -= 1
	if page < 1 {
		page = 1
	}
	return fmt.Sprintf("%s%d", p.url, page)
}
func (p *Pagination) NextPage() string {
	p.calPages()
	page := p.page
	page += 1
	if page > p.pages {
		page = p.pages
	}
	return fmt.Sprintf("%s%d", p.url, page)
}

func (p *Pagination) calPages() {
	p.pages = p.rows / ItemsPerPage
	if p.pages*ItemsPerPage < p.rows {
		p.pages += 1
	}
}

func (p *Pagination) PageStr() string {
	p.calPages()
	return fmt.Sprintf("%d/%d", p.page, p.pages)
}

func (p *Pagination) HasPre() (pre bool) {
	p.calPages()

	pre = true
	if p.pages == 1 || p.page == 1 {
		pre = false
	}
	return
}
func (p *Pagination) HasNext() (next bool) {
	p.calPages()

	next = true
	if p.pages == 1 || p.page == p.pages {
		next = false
	}
	return
}
