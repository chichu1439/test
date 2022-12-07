package controllers

import (
	"github.com/astaxie/beego"
)

// Customer count
type CU000011Controller struct {
	beego.Controller
}

// @Title
// @Description Customer count
// @Param	body		body 	models.SCU0000011I	"Input parameter"
// @Success 200  {object} models.SCU0000011O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000011Controller) Post() ()  {
}
