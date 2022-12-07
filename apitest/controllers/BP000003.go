package controllers

import (
	"github.com/astaxie/beego"
)

// Payment confirm
type BP000003Controller struct {
	beego.Controller
}

// @Title
// @Description  Payment confirm
// @Param	body		body 	models.BP000003Request	"Input parameter"
// @Success 200  {object} models.BP000003Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000003Controller) Post() ()  {
}
