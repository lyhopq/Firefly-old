package controllers

import (
	"firefly/app/models"
	"firefly/app/routes"
	"firefly/app/util"
	"fmt"
	"github.com/huichen/sego"
	"github.com/robfig/config"
	"github.com/robfig/revel"
	"strings"
)

type Application struct {
	*revel.Controller
	Qbs
}

func (c *Application) Begin() revel.Result {
	c.Dial()
	return nil
}

func (c *Application) End() revel.Result {
	c.Close()
	return nil
}
func (c *Application) inject() revel.Result {
	c.RenderArgs["active"] = c.Name
	user := c.connected()
	if user != nil {
		user.UpdateBooKEx(c.q, user.Id)
		c.RenderArgs["user"] = user
		fmt.Println(user.BookCount, user.OwnCount, user.CollectCount)
	}

	// 检查是否需要授权
	fmt.Println("Action: ", c.Action)
	fmt.Println("Args: ", c.RenderArgs)
	value, ok := Permissions[strings.TrimSuffix(c.Action, "Post")]
	if ok {
		if user == nil {
			c.Flash.Error("请先登录")
			c.Session["preUrl"] = c.Request.Request.URL.String()
			return c.Redirect(routes.User.Signin())
		} else {
			if value != user.Type {
				return c.Forbidden("抱歉，您没有得到授权！")
			}
		}
	}

	return nil
}

func (c *Application) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c *Application) getUser(username string) *models.User {
	//user := new(models.User)
	user := models.NewUser()
	c.q.WhereEqual("name", username).Find(user)

	if user.Id == 0 {
		return nil
	}

	return user
}

func Init() {
	basePath = revel.BasePath
	uploadPath = basePath + "/public/upload/"

	c, _ := config.ReadDefault(basePath + "/conf/my.conf")
	driver, _ := c.String("database", "db.driver")
	dbname, _ := c.String("database", "db.dbname")
	user, _ := c.String("database", "db.user")
	password, _ := c.String("database", "db.password")
	host, _ := c.String("database", "db.host")
	registerDb(driver, dbname, user, password, host)

	doubanKey, _ = c.String("keys", "keys.douban")
	proxy, _ = c.String("http", "http.proxy")

	segmenter.LoadDictionary(basePath + "/conf/dict.txt")
}

type App struct {
	Application
}

func (c *App) Index() revel.Result {
	books := models.Recommend(c.q, "hited")

	return c.Render(books)
}

func (c *App) Search(q string, page int) revel.Result {
	var (
		books []*models.Book
		rows  int64
	)

	books, rows = models.SearchBooks(c.q, page, []string{strings.TrimSpace(q)})
	pagination := models.GetPagination(page, rows, routes.App.Search(q, page))
	if rows > 0 {
		return c.Render(books, rows, pagination)
	}

	text := []byte(q)
	segments := segmenter.Segment(text)
	keys := sego.SegmentsToSlice(segments, true)
	keys = util.Filter(keys, util.IsNotIn([]string{" ", "的", "和", "我", "与"}))

	books, rows = models.SearchBooks(c.q, page, keys)
	pagination = models.GetPagination(page, rows, routes.App.Search(q, page))

	return c.Render(books, rows, pagination)
}

func (c *App) Notice() revel.Result {
	type result struct {
		Err bool
		N   *models.Notice
	}

	res := new(result)
	notice := models.LastNotice(c.q)
	if notice != nil {
		res.Err = false
		res.N = notice
	}

	return c.RenderJson(res)
}

func (c *App) Stat() revel.Result {
	type Stat struct {
		User   int64
		Book   int64
		Topic  int64
		Replay int64
	}

	stat := new(Stat)
	stat.Book = models.BookCount(c.q)
	stat.User = models.UserCount(c.q)
	stat.Topic = 0
	stat.Replay = 0

	return c.RenderJson(stat)
}
