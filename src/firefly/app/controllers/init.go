package controllers

import (
	"github.com/robfig/revel"
	"strings"
)

func init() {
	revel.OnAppStart(Init)

	revel.InterceptMethod((*Application).Begin, revel.BEFORE)
	revel.InterceptMethod((*Application).inject, revel.BEFORE)
	revel.InterceptMethod((*Application).End, revel.AFTER)

	revel.TemplateFuncs["section"] = func(s string) []string {
		return strings.Split(s, "\n")
	}
}
