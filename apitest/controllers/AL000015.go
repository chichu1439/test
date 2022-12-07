package controllers

import (
	"github.com/astaxie/beego"
)
// Query Write Off Flow Record
type AL000015Controller struct {
	beego.Controller
}

// @Title
// @Description Query Write Off Flow Record
// @Param	body		body 	models.AL000015I	"Input parameter"
// @Success 200  {object} models.AL000015O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000015Controller) Post() ()  {
}
