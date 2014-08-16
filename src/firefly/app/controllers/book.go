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

	c.updateBookInfo(book)

	subActive := "intro"
	return c.Render(book, subActive)
}

func (c *Book) Borrow(id int64) revel.Result {
	book := models.FindBookById(c.q, id)

	c.updateBookInfo(book)

	borrows := models.FindBorrowsByBookId(c.q, id)

	subActive := "borrow"
	c.Render(book, subActive, borrows)
	return c.RenderTemplate("book/Show.html")
}

func (c *Book) updateBookInfo(book *models.Book) revel.Result {
	if book.Id == 0 {
		return c.NotFound("书籍不存在")
	}

	book.AddHited(c.q)
	user := c.connected()
	if user != nil {
		collect := models.FindCollect(c.q, user.Id, book.Id)
		if collect != nil {
			book.SetCollected(true)
		} else {
			book.SetCollected(false)
		}

		status := models.BorrowStatus(c.q, user.Id, book.Id)
		book.SetBorrow(status)
	}

	return nil
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
	var ok bool
	if user != nil {
		ok = true
		models.RemoveCollect(c.q, user.Id, id)

		book := models.FindBookById(c.q, id)
		book.SubCollect(c.q)
	}

	return c.RenderJson(ok)
}

type Ret struct {
	Ok    bool
	Count int
}

func (c *Book) Booking(id int64) revel.Result {
	user := c.connected()
	var ret Ret
	if user != nil {
		status := models.BorrowStatus(c.q, user.Id, id)
		if status != models.BOOK && status != models.OWN {
			ret.Ok, ret.Count = models.AddBooking(c.q, user.Id, id)
		}

	}

	return c.RenderJson(ret)
}

func (c *Book) UnBooking(id int64) revel.Result {
	user := c.connected()
	var ret Ret
	if user != nil {
		ret.Ok, ret.Count = models.RemoveBooking(c.q, user.Id, id)
	}

	return c.RenderJson(ret)
}

func (c *Book) Intro(id int64) revel.Result {
	book := models.FindBookById(c.q, id)
	return c.RenderJson(book.Introduction)
}
