package controllers

import (
	"github.com/astaxie/beego"
)
// query interest strategy detail
type PI000007Controller struct {
	beego.Controller
}

// @Title
// @Description  query interest strategy detail
// @Param	body		body 	models.PI000007I	"Input parameter"
// @Success 200  {object} models.PI000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PI000007Controller) Post() {
}
