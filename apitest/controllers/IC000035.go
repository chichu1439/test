package controllers

import (
	"github.com/astaxie/beego"
)

// Transaction register info delete
type IC000035Controller struct {
	beego.Controller
}

// @Title
// @Description Transaction register info delete
// @Param	body		body 	models.IC000035I	"Input parameter"
// @Success 200  {object} models.IC000035O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *IC000035Controller) Post() ()  {
}
