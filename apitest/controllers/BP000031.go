package controllers

import (
	"github.com/astaxie/beego"
)

// Get order info
type BP000031Controller struct {
	beego.Controller
}

// @Title
// @Description  Get order info
// @Param	body		body 	models.BP000031Request	"Input parameter"
// @Success 200  {object} models.BP000031Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000031Controller) Post() ()  {
}
