package controllers

import (
	"github.com/astaxie/beego"
)

// Create Template
type NT000001Controller struct {
	beego.Controller
}

// @Title
// @Description Create Template
// @Param	body		body 	models.NT000001I	"Input parameter"
// @Success 200  {object} models.NT000001O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000001Controller) Post() ()  {
}
