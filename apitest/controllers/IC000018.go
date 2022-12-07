package controllers

import (
	"github.com/astaxie/beego"
)

// Trial calculation of repayment plan
type IC000018Controller struct {
	beego.Controller
}

// @Title
// @Description Trial calculation of repayment plan
// @Param	body		body 	models.IC000018I	"Input parameter"
// @Success 200  {object} models.IC000018O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000018Controller) Post() ()  {
}
