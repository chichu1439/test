package controllers

import (
	"github.com/astaxie/beego"
)

// Non-performing Loan Monthly Report
type IC000052Controller struct {
	beego.Controller
}

// @Title
// @Description Non-performing Loan Monthly Report
// @Param	body		body 	models.IC000052I	"Input parameter"
// @Success 200  {object} models.IC000052O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000052Controller) Post() ()  {
}
