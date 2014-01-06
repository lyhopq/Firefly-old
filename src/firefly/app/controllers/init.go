package controllers

import "github.com/robfig/revel"

func init() {
	revel.OnAppStart(Init)

	revel.InterceptMethod((*Application).Begin, revel.BEFORE)
	revel.InterceptMethod((*Application).inject, revel.BEFORE)
	revel.InterceptMethod((*Application).End, revel.AFTER)
}
