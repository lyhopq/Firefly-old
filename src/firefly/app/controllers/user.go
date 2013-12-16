package controllers

import (
	"firefly/app/models"
	"firefly/app/routes"

	"github.com/coocood/qbs"
	"github.com/robfig/revel"
)

//"code.google.com/p/go-uuid/uuid"
//"fmt"
//"github.com/coocood/qbs"
//"github.com/disintegration/imaging"

//"image"
//"strings"

type User struct {
	Application
}

func (c *User) Signup() revel.Result {
	return c.Render()
}

func (c *User) SignupPost(user models.User) revel.Result {
	user.Validate(c.q, c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Signup())
	}

	user.Type = MemberGroup
	user.Avatar = defaultAvatar

	if !user.Save(c.q) {
		c.Flash.Error("注册用户失败")
		return c.Redirect(routes.User.Signup())
	}

	//perm := new(models.Permissions)
	//perm.UserId = user.Id
	//perm.Perm = MemberGroup
	//perm.Save(c.q)

	return c.Redirect(routes.User.Signin())
}

//
func (c *User) Signin() revel.Result {
	return c.Render()
}

func (c *User) SigninPost(name, password string) revel.Result {
	c.Validation.Required(name).Message("请输入用户名")
	c.Validation.Required(password).Message("请输入密码")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Signin())
	}

	user := new(models.User)
	condition := qbs.NewCondition("name = ?", name).
		And("hashed_password = ?", models.EncryptPassword(password))
	c.q.Condition(condition).Find(user)

	if user.Id == 0 {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Out["user"] = name
		c.Flash.Error("用户名或密码错误")
		return c.Redirect(routes.User.Signin())
	}

	c.Session["user"] = name

	preUrl, ok := c.Session["preUrl"]
	if ok {
		return c.Redirect(preUrl)
	}

	return c.Redirect(routes.App.Index())
}

func (c *User) Signout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.App.Index())
}

func (c *User) Borrow(page int) revel.Result {
	title := "已预借"
	subActive := "book"

	borrows, rows := models.GetBorrows(c.q, page, "status", models.BOOK, "id")
	pagination := models.GetPagination(page, rows, routes.User.Borrow(page))

	return c.Render(title, subActive, borrows, pagination)
}

func (c *User) Owned(page int) revel.Result {
	title := "已借阅"
	subActive := "own"

	borrows, rows := models.GetBorrows(c.q, page, "status", models.OWN, "id")
	pagination := models.GetPagination(page, rows, routes.User.Borrow(page))

	c.Render(title, subActive, borrows, pagination)
	return c.RenderTemplate("user/borrow.html")
}

func (c *User) BorrowHis(page int) revel.Result {
	title := "借阅历史"
	subActive := "his"

	borrows, rows := models.GetBorrows(c.q, page, "", "", "id")
	pagination := models.GetPagination(page, rows, routes.User.Borrow(page))

	c.Render(title, subActive, borrows, pagination)
	return c.RenderTemplate("user/borrow.html")
}

//
//func (c *User) Edit() revel.Result {
//	id := c.RenderArgs["user"].(*models.User).Id
//	user := findUserById(c.q, id)
//	if user.Id == 0 {
//		return c.NotFound("用户不存在")
//	}
//
//	return c.Render(user, avatars)
//}
//
//func (c *User) EditPost(avatar string) revel.Result {
//	id := c.RenderArgs["user"].(*models.User).Id
//	checkFileExt(c.Controller, imageExts, "picture", "Only image")
//	user := findUserById(c.q, id)
//	if user.Id == 0 {
//		return c.NotFound("用户不存在")
//	}
//
//	if c.Validation.HasErrors() {
//		c.Validation.Keep()
//		c.FlashParams()
//		return c.Redirect(routes.User.Edit())
//	}
//
//	if ok, _ := getFileExt(c.Request, "picture"); ok {
//		picture := saveFile(c.Request, "picture")
//		src, _ := imaging.Open(uploadPath + picture)
//		var dst *image.NRGBA
//
//		dst = imaging.Thumbnail(src, 48, 48, imaging.CatmullRom)
//		avatar = "thumb" + picture
//		imaging.Save(dst, uploadPath+avatar)
//		deleteFile(picture)
//	}
//
//	if avatar != "" {
//		if strings.HasPrefix(user.Avatar, "thumb") {
//			deleteFile(user.Avatar)
//		}
//		user.Avatar = avatar
//	}
//
//	if user.Save(c.q) {
//		c.Flash.Success("保存成功")
//	} else {
//		c.Flash.Error("保存信息失败")
//	}
//
//	return c.Redirect(routes.User.Edit())
//}
//
//func (c *User) Validate(code string) revel.Result {
//	user := findUserByCode(c.q, code)
//	if user.Id == 0 {
//		return c.NotFound("用户不存在或校验码错误")
//	}
//
//	user.IsActive = true
//	user.Save(c.q)
//
//	c.Flash.Success("您的账号成功激活，请登录！")
//
//	return c.Redirect(routes.User.Signin())
//}
//
//func findUserById(q *qbs.Qbs, id int64) *models.User {
//	user := new(models.User)
//	q.WhereEqual("id", id).Find(user)
//
//	return user
//}
//
//func findUserByCode(q *qbs.Qbs, code string) *models.User {
//	user := new(models.User)
//	q.WhereEqual("validate_code", code).Find(user)
//
//	return user
//}
