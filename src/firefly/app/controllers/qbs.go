package controllers

import (
	"firefly/app/models"
	"fmt"
	_ "github.com/coocood/mysql"
	"github.com/coocood/qbs"
)

type Qbs struct {
	q *qbs.Qbs
}

func (c *Qbs) Dial() {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
	}
	c.q = q
}

func (c *Qbs) Close() {
	c.q.Close()
}

func registerDb(driver, dbname, user, password, host string) {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)
	qbs.Register(driver, params, dbname, qbs.NewMysql())
	err := createTabel()
	if err != nil {
		fmt.Println(err)
	}
}

func createTabel() error {
	migration, err := qbs.GetMigration()
	if err != nil {
		return err
	}
	defer migration.Close()

	err = migration.CreateTableIfNotExists(new(models.User))
	err = migration.CreateTableIfNotExists(new(models.Book))
	err = migration.CreateTableIfNotExists(new(models.Collect))
	err = migration.CreateTableIfNotExists(new(models.Borrow))
	//err = migration.CreateTableIfNotExists(new(models.Permissions))

	return err
}
