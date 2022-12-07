package controllers

import (
	"github.com/astaxie/beego"
)
// delete base interest rate
type PI000004Controller struct {
	beego.Controller
}

// @Title
// @Description  delete base interest rate
// @Param	body		body 	models.PI000004I	"Input parameter"
// @Success 200  {object} models.PI000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000004Controller) Post() {
}
