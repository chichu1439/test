package controllers

import (
	"github.com/astaxie/beego"
)

// Update Customer Info
type CU000004Controller struct {
	beego.Controller
}

// @Title
// @Description Update Customer Info
// @Param	body		body 	models.SCU0000004I	"Input parameter"
// @Success 200  {object} models.SCU0000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000004Controller) Post() ()  {
}
