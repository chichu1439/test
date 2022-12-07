package controllers

import (
	"github.com/astaxie/beego"
)

// Update Template
type NT000004Controller struct {
	beego.Controller
}

// @Title
// @Description Update Template
// @Param	body		body 	models.NT000004I	"Input parameter"
// @Success 200  {object} models.NT000004O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000004Controller) Post() ()  {
}
