package controllers

import (
	"firefly/app/models"
	"firefly/app/routes"
	"fmt"

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

	if page < 1 {
		page = 1
	}
	books, rows := models.GetBooks(c.q, page, "", "", "readed")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	c.Render(title, subActive, books, pagination)
	return c.RenderTemplate("book/Index.html")
}

func (c *Book) MostComment(page int) revel.Result {
	title := "最多评论"
	subActive := "mostComment"

	if page < 1 {
		page = 1
	}
	books, rows := models.GetBooks(c.q, page, "", "", "Commented")
	pagination := models.GetPagination(page, rows, routes.Book.Index(page))

	c.Render(title, subActive, books, pagination)

	return c.RenderTemplate("book/Index.html")
}

func (c *Book) Show(id int64, opr string) revel.Result {
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

	user := c.connected()
	if user != nil {

		collect := models.FindCollect(c.q, user.Id, id)
		fmt.Println(collect, "7777777777z")
		fmt.Println(user, "7777777777z")
		if collect != nil {
			book.IsCollected = true
		}
	}
	//collect := models.FindCollect(c.q, user.Id, id)
	//fmt.Println(collect, "7777777777z")

	/*
		switch opr {
		case "show":
			book.Hited += 1
			type Book struct {
				Hited int64
			}
			b := new(Book)
			b.Hited = book.Hited
			c.q.WhereEqual("id", id).Update(b)

		case "collect":
			book.Collected += 1
			book.IsCollected = true
			type Book struct {
				Collected int64
			}
			b := new(Book)
			b.Collected = book.Collected
			c.q.WhereEqual("id", id).Update(b)
		case "uncollect":
			book.Collected -= 1
			if book.Collected < 0 {
				book.Collected = 0
			}
			book.IsCollected = false
			type Book struct {
				Collected int64
			}
			b := new(Book)
			b.Collected = book.Collected
			c.q.WhereEqual("id", id).Update(b)
		case "book":
			book.Existing -= 1
			if book.Existing < 0 {
				book.Existing = 0
			}
		}
	*/

	return c.Render(book)
}

func (c *Book) Collect(id int64) revel.Result {
	fmt.Println(id, "1111111111111111")
	user := c.connected()
	signin := false
	if user != nil {
		collect := new(models.Collect)
		collect.User = user.Id
		collect.Book = id
		c.q.Save(collect)
		signin = true

		book := models.FindBookById(c.q, id)
		book.Collected += 1

		type Book struct {
			Collected int64
		}
		b := new(Book)
		b.Collected = book.Collected
		c.q.WhereEqual("id", id).Update(b)
	}

	return c.RenderJson(signin)
}

func (c *Book) UnCollect(id int64) revel.Result {
	user := c.connected()
	var signin bool
	if user != nil {
		collect := models.FindCollect(c.q, user.Id, id)
		c.q.Delete(collect)
		fmt.Println(id, "1111111111111111")
		signin = true

		book := models.FindBookById(c.q, id)
		book.Collected -= 1
		if book.Collected < 0 {
			book.Collected = 0
		}

		type Book struct {
			Collected int64
		}
		b := new(Book)
		b.Collected = book.Collected
		c.q.WhereEqual("id", id).Update(b)
	}

	return c.RenderJson(signin)
}

/*
func (c *Book) Collect(id int64) revel.Result {
	book := models.FindBookById(c.q, id)
	if book.Id == 0 {
		return c.NotFound("书籍不存在")
	}

	book.Collected += 1
	book.IsCollected = true

	type Book struct {
		Collected int64
	}

	b := new(Book)
	b.Collected = book.Collected
	c.q.WhereEqual("id", id).Update(b)

	c.Render(book)
	return c.RenderTemplate("book/Show.html")
}*/
func (c *Book) Booking(id int64) revel.Result {
	return nil
}
