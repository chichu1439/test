package controllers

import (
	"github.com/astaxie/beego"
)

// Collection Summary of Total Overdue Amount
type IC000040Controller struct {
	beego.Controller
}

// @Title
// @Description Collection Summary of Total Overdue Amount
// @Param	body		body 	models.IC000040I	"Input parameter"
// @Success 200  {object} models.IC000040O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000040Controller) Post() ()  {
}
