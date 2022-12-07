package controllers

import (
	"github.com/astaxie/beego"
)

// new loan
type IC000069Controller struct {
	beego.Controller
}

// @Title
// @Description  new loan
// @Param	body		body 	models.IC000069Request	"Input parameter"
// @Success 200  {object} models.IC000069Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *IC000069Controller) Post() ()  {
}
