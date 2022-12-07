package controllers

import (
	"github.com/astaxie/beego"
)

// Update Customer Stauts
type CU000010Controller struct {
	beego.Controller
}

// @Title
// @Description Update Customer Stauts
// @Param	body		body 	models.SCU0000010I	"Input parameter"
// @Success 200  {object} models.SCU0000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000010Controller) Post() ()  {
}
