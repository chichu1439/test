package controllers

import (
	"github.com/astaxie/beego"
)

// Loan approve
type IC000007Controller struct {
	beego.Controller
}

// @Title
// @Description Loan approve
// @Param	body		body 	models.SIC0000007I	"Input parameter"
// @Success 200  {object} models.SIC0000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000007Controller) Post() ()  {
}
