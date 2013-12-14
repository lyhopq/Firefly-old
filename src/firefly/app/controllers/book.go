package controllers

import (
	"firefly/app/models"
	"firefly/app/routes"
	"github.com/robfig/revel"
)

type Book struct {
	Application
}

func (c *Book) Index(page int) revel.Result {
	title := "最近上架"
	subActive := "latestShelf"

	if page < 1 {
		page = 1
	}
	books, rows := models.GetBooks(c.q, page, "", "", "id")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	return c.Render(title, subActive, books, pagination)
}

func (c *Book) MostRead(page int) revel.Result {
	title := "最多借阅"
	subActive := "mostRead"

	books, rows := models.GetBooks(c.q, page, "", "", "readed")
	rows += 0

	c.Render(title, subActive, books)
	return c.RenderTemplate("book/Index.html")
}

func (c *Book) MostComment(page int) revel.Result {
	title := "最多评论"
	subActive := "mostComment"

	books, rows := models.GetBooks(c.q, page, "", "", "Commented")
	rows += 0

	c.Render(title, subActive, books)
	return c.RenderTemplate("book/Index.html")
}
func (c *Book) Show(id int64) revel.Result {
	book := models.FindBookById(c.q, id)
	if book.Id == 0 {
		return c.NotFound("书籍不存在")
	}

	book.Hited += 1

	type Book struct {
		Hited int64
	}

	b := new(Book)
	b.Hited = book.Hited
	c.q.WhereEqual("id", id).Update(b)

	return c.Render()
}
