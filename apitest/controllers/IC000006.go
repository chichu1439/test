package controllers

import (
	"github.com/astaxie/beego"
)

// Loan application
type IC000006Controller struct {
	beego.Controller
}

// @Title
// @Description Loan application
// @Param	body		body 	models.SIC0000006I	"Input parameter"
// @Success 200  {object} models.SIC0000006O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000006Controller) Post() ()  {
}
