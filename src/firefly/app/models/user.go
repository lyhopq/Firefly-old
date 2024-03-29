package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/coocood/qbs"
	"github.com/robfig/revel"
)

type User struct {
	Id   int64
	Name string `qbs:"size:32,unique,notnull"`
	//Email           string `qbs:"size:32,unique,notnull"`
	Password        string `qbs:"-"`
	ConfirmPassword string `qbs:"-"`
	HashedPassword  string `qbs:"size:32,notnull"`
	Type            int    `qbs:"default:2,notnull"` //1管理员，2普通用户
	Avatar          string `qbs:"size:255"`
	//ValidateCode    string      `qbs:"size:255"`
	//IsActive        bool
	Created time.Time
	//Updated         time.Time

	*BookEx `qbs:"-"`
}

var (
	nameRegex = regexp.MustCompile("^\\S*$")
)

func NewUser() *User {
	user := new(User)
	user.BookEx = new(BookEx)
	return user
}

func (user *User) Validate(q *qbs.Qbs, v *revel.Validation) {
	v.Required(user.Name).Message("请输入用户名")
	v.MinSize(user.Name, 3).Message("用户名最少三位")
	valid := v.Match(user.Name, nameRegex).Message("用户名不能有空格")
	if valid.Ok {
		if user.HasName(q) {
			err := &revel.ValidationError{
				Message: "该用户名已经注册过",
				Key:     "user.Name",
			}
			valid.Error = err
			valid.Ok = false

			v.Errors = append(v.Errors, err)
		}
	}

	v.Required(user.Password).Message("请输入密码")
	v.MinSize(user.Password, 3).Message("密码最少三位")
	v.Required(user.ConfirmPassword == user.Password).Message("密码不一致")
}

func (u *User) HasName(q *qbs.Qbs) bool {
	user := new(User)
	q.WhereEqual("name", u.Name).Find(user)

	return user.Id > 0

}

/*
func (u *User) HasEmail(q *qbs.Qbs) bool {
	user := new(User)
	q.WhereEqual("email", u.Email).Find(user)

	return user.Id > 0
}
*/

// 加密密码,转成md5
func EncryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u *User) Save(q *qbs.Qbs) bool {
	if u.Password != "" {
		u.HashedPassword = EncryptPassword(u.Password)
	}

	_, err := q.Save(u)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (u *User) IsAdmin() bool {
	return u.Type == 1
}

// 是否是默认头像
func (u *User) IsDefaultAvatar(avatar string) bool {
	return avatar == u.Avatar
}

// 头像的图片地址
func (u *User) AvatarImgSrc() string {
	if strings.HasPrefix(u.Avatar, "thumb") {
		return fmt.Sprintf("/public/upload/%s", u.Avatar)
	}
	return fmt.Sprintf("/public/img/%s", u.Avatar)
}

func UserCount(q *qbs.Qbs) int64 {
	return q.WhereEqual("type", 2).Count("user")
}

type BookEx struct {
	BookCount    int64
	OwnCount     int64
	CollectCount int64
}

func (b *BookEx) UpdateBooKEx(q *qbs.Qbs, uid int64) {
	b.BookCount = BorrowCount(q, uid, BOOK)
	b.OwnCount = BorrowCount(q, uid, OWN)
	b.CollectCount = UserCollectCount(q, uid)
	q.Reset()
}
