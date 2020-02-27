package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *LoginController) StaticBlock() {

}

// @router /all/:key [get]
func (this *LoginController) AllBlock() {
	this.Ctx.WriteString("Login/all")
}
