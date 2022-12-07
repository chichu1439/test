package controllers

import (
	"github.com/astaxie/beego"
)

// Loan Customer Statistics
type IC000014Controller struct {
	beego.Controller
}

// @Title
// @Description Loan Customer Statistics
// @Param	body		body 	models.IC000014I	"Input parameter"
// @Success 200  {object} models.IC000014O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000014Controller) Post() ()  {
}
