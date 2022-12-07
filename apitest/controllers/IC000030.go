package controllers

import (
	"github.com/astaxie/beego"
)

// Loan pdf
type IC000030Controller struct {
	beego.Controller
}

// @Title
// @Description Loan pdf
// @Param	body		body 	models.IC000030I	"Input parameter"
// @Success 200  {object} models.IC000030O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000030Controller) Post() ()  {
}
