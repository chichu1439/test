package controllers

import (
	"github.com/astaxie/beego"
)

// Get loan product info
type BP000034Controller struct {
	beego.Controller
}

// @Title
// @Description  Get loan product info
// @Param	body		body 	models.BP000034Request	"Input parameter"
// @Success 200  {object} models.BP000034Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000034Controller) Post() ()  {
}
