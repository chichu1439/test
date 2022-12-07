package controllers

import (
	"github.com/astaxie/beego"
)

// Loan Details Edit application
type IC000028Controller struct {
	beego.Controller
}

// @Title
// @Description Loan Details Edit application
// @Param	body		body 	models.SIC0000028I	"Input parameter"
// @Success 200  {object} models.SIC0000028O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000028Controller) Post() ()  {
}
