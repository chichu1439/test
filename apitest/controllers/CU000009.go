package controllers

import (
	"github.com/astaxie/beego"
)

// Update customer’s contract info
type CU000009Controller struct {
	beego.Controller
}

// @Title
// @Description Update customer’s contract info
// @Param	body		body 	models.SCU0000009I	"Input parameter"
// @Success 200  {object} models.SCU0000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *CU000009Controller) Post() ()  {
}
