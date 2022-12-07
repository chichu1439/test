package controllers

import (
	"github.com/astaxie/beego"
)
// delete interest strategy
type PI000009Controller struct {
	beego.Controller
}

// @Title
// @Description  delete interest strategy
// @Param	body		body 	models.PI000009I	"Input parameter"
// @Success 200  {object} models.PI000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000009Controller) Post() {
}
