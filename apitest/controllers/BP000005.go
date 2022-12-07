package controllers

import (
	"github.com/astaxie/beego"
)

// Get loan detail
type BP000005Controller struct {
	beego.Controller
}

// @Title
// @Description  Get loan detail
// @Param	body		body 	models.BP000005Request	"Input parameter"
// @Success 200  {object} models.BP000005Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000005Controller) Post() ()  {
}
