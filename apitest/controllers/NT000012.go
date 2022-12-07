package controllers

import (
	"github.com/astaxie/beego"
)

// Update Message Target object
type NT000012Controller struct {
	beego.Controller
}

// @Title
// @Description Update Message Target object
// @Param	body		body 	models.NT000012I	"Input parameter"
// @Success 200  {object} models.NT000012O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000012Controller) Post() ()  {
}
