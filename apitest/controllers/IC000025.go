package controllers

import (
	"github.com/astaxie/beego"
)

// Application of manual disbursement
type IC000025Controller struct {
	beego.Controller
}

// @Title
// @Description Application of manual disbursement
// @Param	body		body 	models.IC000025I	"Input parameter"
// @Success 200  {object} models.IC000025O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000025Controller) Post() ()  {
}
