package controllers

import (
	"github.com/astaxie/beego"
)

// loan repayment status inquiry
type AL000063Controller struct {
	beego.Controller
}

// @Title
// @Description loan repayment status inquiry
// @Param	body		body 	models.AL000063I	"Input parameter"
// @Success 200  {object} models.AL000063O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *AL000063Controller) Post() ()  {
}
