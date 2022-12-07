package controllers

import (
	"github.com/astaxie/beego"
)

// Query loan account period info
type AL000054Controller struct {
	beego.Controller
}

// @Title
// @Description Query loan account period info
// @Param	body		body 	models.AL000054I	"Input parameter"
// @Success 200  {object} models.AL000054O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000054Controller) Post() ()  {
}
