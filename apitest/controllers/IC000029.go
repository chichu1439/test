package controllers

import (
	"github.com/astaxie/beego"
)

// Loan Details Edit approve
type IC000029Controller struct {
	beego.Controller
}

// @Title
// @Description Loan Details Edit approve
// @Param	body		body 	models.SIC0000029I	"Input parameter"
// @Success 200  {object} models.SIC0000029O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000029Controller) Post() ()  {
}
