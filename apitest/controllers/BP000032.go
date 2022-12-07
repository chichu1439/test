package controllers

import (
	"github.com/astaxie/beego"
)

// Get customer info
type BP000032Controller struct {
	beego.Controller
}

// @Title
// @Description  Get customer info
// @Param	body		body 	models.BP000032Request	"Input parameter"
// @Success 200  {object} models.BP000032Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000032Controller) Post() ()  {
}
