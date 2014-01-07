package controllers

import (
	"firefly/app/models"
	"firefly/app/routes"

	"strconv"
	"strings"

	"github.com/lyhopq/douban"
	"github.com/robfig/revel"
)

type Admin struct {
	Application
}

func (c *Admin) Index() revel.Result {
	return c.Render()
}

func (c *Admin) ListUser() revel.Result {
	var users []*models.User
	c.q.FindAll(&users)

	return c.Render(users)
}

func (c *Admin) DeleteUser(id int64) revel.Result {
	user := new(models.User)
	user.Id = id
	c.q.Delete(user)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) ListBook(page int) revel.Result {
	title := "图书列表"
	books, rows := models.GetBooks(c.q, page, "", "", "id")
	pagination := models.GetPagination(page, rows, routes.Admin.ListBook(page))

	return c.Render(title, books, pagination)
}

func (c *Admin) NewBook() revel.Result {
	title := "添加图书"
	return c.Render(title)
}

func (c *Admin) NewBookPost(book models.Book) revel.Result {
	book.Validate(c.q, c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewBook())
	}

	if book.Existing == 0 {
		book.Existing = book.Holding
	}

	if !book.Save(c.q) {
		c.Flash.Error("添加图书失败")
	}

	page := 1
	return c.Redirect(routes.Admin.ListBook(page))
}

func (c *Admin) EditBook(id int64) revel.Result {
	title := "编辑图书"

	book := models.FindBookById(c.q, id)
	if book.Id == 0 {
		return c.NotFound("书籍不存在")
	}

	c.Render(title, book)

	return c.RenderTemplate("admin/NewBook.html")
}

func (c *Admin) EditBookPost(id int64, book models.Book) revel.Result {
	book.Id = id
	book.Validate(c.q, c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.EditBook(id))
	}

	if !book.Save(c.q) {
		c.Flash.Error("编辑图书失败")
	}

	page := 1
	return c.Redirect(routes.Admin.ListBook(page))
}

func (c *Admin) DeleteBook(id int64) revel.Result {
	book := new(models.Book)
	book.Id = id
	c.q.Delete(book)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) ListBorrow(page int) revel.Result {
	title := "借阅请求"
	borrows, rows := models.GetBorrows(c.q, page, "status", models.BOOK, "id")
	pagination := models.GetPagination(page, rows, routes.Admin.ListBorrow(page))

	return c.Render(title, borrows, pagination)
}

func (c *Admin) ListReturn(page int) revel.Result {
	title := "归还请求"
	borrows, rows := models.GetBorrows(c.q, page, "status", models.PRERET, "id")
	pagination := models.GetPagination(page, rows, routes.Admin.ListBorrow(page))

	return c.Render(title, borrows, pagination)
}

func (c *Admin) ConfirmBorrow(id int64) revel.Result {
	borrow := models.FindBorrowById(c.q, id)
	borrow.SetBorrowStatus(c.q, models.OWN)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) RejectBorrow(id int64) revel.Result {
	borrow := models.FindBorrowById(c.q, id)
	c.q.Delete(borrow)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) ConfirmReturn(id int64) revel.Result {
	borrow := models.FindBorrowById(c.q, id)
	borrow.SetBorrowStatus(c.q, models.RETURN)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) RejectReturn(id int64) revel.Result {
	borrow := models.FindBorrowById(c.q, id)
	borrow.SetBorrowStatus(c.q, models.OWN)

	return c.RenderJson([]byte("true"))
}

func fetchBook(isbn string, book *models.Book) {
	api := douban.NewApi()
	dBook := api.GetBookByIsbn(isbn)

	book.Title = dBook.Title
	book.Author = strings.Join(dBook.Author, ", ")
	book.Translator = strings.Join(dBook.Translator, ", ")
	pages, err := strconv.ParseInt(dBook.Pages, 10, 0)
	if err != nil {
		book.Pages = 0
	} else {
		book.Pages = int(pages)
	}
	book.Introduction = dBook.Summary
	book.Publisher = dBook.Publisher
	book.PublicationDate = dBook.PubDate

	fileName := "thumb_" + isbn + ".jpg"
	_, err = getImg(dBook.Cover.Large, uploadPath+fileName)
	if err == nil {
		book.Cover = fileName
	}
}

func (c *Admin) FetchBook(isbn string) revel.Result {
	book := new(models.Book)
	fetchBook(isbn, book)
	return c.RenderJson(book)
}

func (c *Admin) UpdateBook(id int64) revel.Result {
	book := models.FindBookById(c.q, id)
	fetchBook(book.Isbn, book)
	return c.RenderJson(book)
}

/*
func (c *Admin) ListCategory() revel.Result {
	categories := getCategories(c.q)

	return c.Render(categories)
}

func (c *Admin) DeleteCategory(id int64) revel.Result {
	fmt.Println(id, "1111111111111111")
	category := new(models.Category)
	category.Id = id
	c.q.Delete(category)

	return c.RenderJson([]byte("true"))
}

func (c *Admin) NewCategory() revel.Result {
	title := "新建分类"
	return c.Render(title)
}

func (c *Admin) NewCategoryPost(category models.Category) revel.Result {
	category.Validate(c.q, c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	if !category.Save(c.q) {
		c.Flash.Error("添加分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}

func (c *Admin) EditCategory(id int64) revel.Result {
	title := "编辑分类"

	category := findCategoryById(c.q, id)
	if category.Id == 0 {
		return c.NotFound("分类不存在")
	}

	c.Render(title, category)

	return c.RenderTemplate("admin/NewCategory.html")
}

func (c *Admin) EditCategoryPost(id int64, category models.Category) revel.Result {
	category.Id = id
	category.Validate(c.q, c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Admin.NewCategory())
	}

	if !category.Save(c.q) {
		c.Flash.Error("编辑分类失败")
	}

	return c.Redirect(routes.Admin.ListCategory())
}

func getCategories(q *qbs.Qbs) []*models.Category {
	var categories []*models.Category
	if err := q.FindAll(&categories); err != nil {
		fmt.Println(err)
	}

	return categories
}

func findCategoryById(q *qbs.Qbs, id int64) *models.Category {
	category := new(models.Category)
	q.WhereEqual("id", id).Find(category)

	return category
}
*/
