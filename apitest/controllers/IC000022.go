package controllers

import (
	"github.com/astaxie/beego"
)

// Loan list
type IC000022Controller struct {
	beego.Controller
}

// @Title
// @Description Loan list
// @Param	body		body 	models.SIC0000022I	"Input parameter"
// @Success 200  {object} models.SIC0000022O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000022Controller) Post() ()  {
}
