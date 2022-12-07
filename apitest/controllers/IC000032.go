package controllers

import (
	"github.com/astaxie/beego"
)

// Loan history inquiry
type IC000032Controller struct {
	beego.Controller
}

// @Title
// @Description Loan history inquiry
// @Param	body		body 	models.IC000032I	"Input parameter"
// @Success 200  {object} models.IC000032O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000032Controller) Post() ()  {
}
