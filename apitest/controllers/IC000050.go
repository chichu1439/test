package controllers

import (
	"github.com/astaxie/beego"
)

// Closed Loan Monthly Report
type IC000050Controller struct {
	beego.Controller
}

// @Title
// @Description Closed Loan Monthly Report
// @Param	body		body 	models.IC000050I	"Input parameter"
// @Success 200  {object} models.IC000050O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000050Controller) Post() ()  {
}
