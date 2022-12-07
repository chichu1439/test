package controllers

import (
	"github.com/astaxie/beego"
)

// Get payment method list
type BP000035Controller struct {
	beego.Controller
}

// @Title
// @Description  Get payment method list
// @Param	body		body 	models.BP000035Request	"Input parameter"
// @Success 200  {object} models.BP000035Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000035Controller) Post() ()  {
}
