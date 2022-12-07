package controllers

import (
	"github.com/astaxie/beego"
)
// Query Product Fee Info for Batch Calculation
type PD000084Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Product Fee Info for Batch Calculation
// @Param	body		body 	models.PD000084I	"Input parameter"
// @Success 200  {object} models.PD000084O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000084Controller) Post() {
}
