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

	books, rows := models.GetBooks(c.q, page, "", "", "id")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	return c.Render(title, subActive, books, pagination)
}

func (c *Book) MostRead(page int) revel.Result {
	title := "最多借阅"
	subActive := "mostRead"

	books, rows := models.GetBooks(c.q, page, "", "", "readed")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	c.Render(title, subActive, books, pagination)
	return c.RenderTemplate("book/Index.html")
}

func (c *Book) MostComment(page int) revel.Result {
	title := "最多评论"
	subActive := "mostComment"

	books, rows := models.GetBooks(c.q, page, "", "", "Commented")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	c.Render(title, subActive, books, pagination)

	return c.RenderTemplate("book/Index.html")
}

func (c *Book) Show(id int64) revel.Result {
	book := models.FindBookById(c.q, id)
	if book.Id == 0 {
		return c.NotFound("书籍不存在")
	}

	book.AddHited(c.q)
	user := c.connected()
	collect := models.FindCollect(c.q, user.Id, id)
	if user != nil && collect != nil {
		book.SetCollected()
	}

	return c.Render(book)
}

func (c *Book) Collect(id int64) revel.Result {
	user := c.connected()
	var signin bool
	if user != nil {
		signin = true
		models.AddCollect(c.q, user.Id, id)

		book := models.FindBookById(c.q, id)
		book.AddCollect(c.q)
	}

	return c.RenderJson(signin)
}

func (c *Book) UnCollect(id int64) revel.Result {
	user := c.connected()
	var signin bool
	if user != nil {
		signin = true
		collect := models.FindCollect(c.q, user.Id, id)
		models.RemoveCollect(c.q, collect)

		book := models.FindBookById(c.q, id)
		book.SubCollect(c.q)
	}

	return c.RenderJson(signin)
}

func (c *Book) Booking(id int64) revel.Result {
	return nil
}
