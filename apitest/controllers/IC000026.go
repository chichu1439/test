package controllers

import (
	"github.com/astaxie/beego"
)

// Approve manual disbursement
type IC000026Controller struct {
	beego.Controller
}

// @Title
// @Description Approve manual disbursement
// @Param	body		body 	models.IC000026I	"Input parameter"
// @Success 200  {object} models.IC000026O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000026Controller) Post() ()  {
}
