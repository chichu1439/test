package controllers

import (
	"github.com/astaxie/beego"
)

// Add Message Target object
type NT000011Controller struct {
	beego.Controller
}

// @Title
// @Description Add Message Target object
// @Param	body		body 	models.NT000011I	"Input parameter"
// @Success 200  {object} models.NT000011O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000011Controller) Post() ()  {
}
