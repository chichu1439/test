package controllers

import (
	"github.com/astaxie/beego"
)

// Get loan list
type BP000004Controller struct {
	beego.Controller
}

// @Title
// @Description  Get loan list
// @Param	body		body 	models.BP000004Request	"Input parameter"
// @Success 200  {object} models.BP000004Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000004Controller) Post() ()  {
}
