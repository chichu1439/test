package controllers

import (
	"github.com/astaxie/beego"
)

// Query Write Off Register List
type AL000014Controller struct {
	beego.Controller
}

// @Title
// @Description Query Write Off Register List
// @Param	body		body 	models.AL000014I	"Input parameter"
// @Success 200  {object} models.AL000014O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000014Controller) Post() ()  {
}
