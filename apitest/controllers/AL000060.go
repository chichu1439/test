package controllers

import (
	"github.com/astaxie/beego"
)

// Payment Link inquiry
type AL000060Controller struct {
	beego.Controller
}

// @Title
// @Description Payment Link inquiry
// @Param	body		body 	models.AL000060I	"Input parameter"
// @Success 200  {object} models.AL000060O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000060Controller) Post() ()  {
}
