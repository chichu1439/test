package controllers

import (
	"github.com/astaxie/beego"
)

// New Loan Daily/Monthly Report
type IC000049Controller struct {
	beego.Controller
}

// @Title
// @Description New Loan Daily/Monthly Report
// @Param	body		body 	models.IC000049I	"Input parameter"
// @Success 200  {object} models.IC000049O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000049Controller) Post() ()  {
}
