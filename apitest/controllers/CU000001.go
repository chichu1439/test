package controllers

import (
	"github.com/astaxie/beego"
)

// Check New Customer
type CU000001Controller struct {
	beego.Controller
}

// @Title
// @Description Check New Customer
// @Param	body		body 	models.SCU0000001I	"Input parameter"
// @Success 200  {object} models.SCU0000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000001Controller) Post() ()  {
}
