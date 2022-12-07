package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment plan modificaiton
type IC000054Controller struct {
	beego.Controller
}

// @Title
// @Description Repayment plan modificaiton
// @Param	body		body 	models.IC000054I	"Input parameter"
// @Success 200  {object} models.IC000054O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000054Controller) Post() ()  {
}
