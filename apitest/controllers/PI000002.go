package controllers

import (
	"github.com/astaxie/beego"
)
// query base interest rate
type PI000002Controller struct {
	beego.Controller
}

// @Title
// @Description  query base interest rate
// @Param	body		body 	models.PI000002I	"Input parameter"
// @Success 200  {object} models.PI000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000002Controller) Post() {
}
