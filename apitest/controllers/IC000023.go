package controllers

import (
	"github.com/astaxie/beego"
)

// Loan detail
type IC000023Controller struct {
	beego.Controller
}

// @Title
// @Description Loan detail
// @Param	body		body 	models.SIC0000023I	"Input parameter"
// @Success 200  {object} models.SIC0000023O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000023Controller) Post() ()  {
}
