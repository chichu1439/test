package controllers

import (
	"github.com/astaxie/beego"
)
//create fee business scenario
type PR000009Controller struct {
	beego.Controller
}

// @Title
// @Description  create fee business scenario
// @Param	body		body 	models.PR000009I	"Input parameter"
// @Success 200  {object} models.PR000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PR000009Controller) Post() {
}
