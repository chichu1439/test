package controllers

import (
	"github.com/astaxie/beego"
)

// Create New Customer
type CU000002Controller struct {
	beego.Controller
}

// @Title
// @Description Create New Customer
// @Param	body		body 	models.SCU0000002I	"Input parameter"
// @Success 200  {object} models.SCU0000002O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000002Controller) Post() ()  {
}
