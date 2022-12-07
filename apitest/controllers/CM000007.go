package controllers

import (
	"github.com/astaxie/beego"
)

// file download
type CM000007Controller struct {
	beego.Controller
}

// @Title
// @Description file download
// @Param	body		body 	models.CM000007I	"Input parameter"
// @Success 200  {object} models.CM000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CM000007Controller) Post() ()  {
}
