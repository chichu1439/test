package controllers

import (
	"github.com/astaxie/beego"
)
// Query Loan Fee strategy
type PD000047Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Loan Fee strategy
// @Param	body		body 	models.SPD0000047I	"Input parameter"
// @Success 200  {object} models.SPD0000047O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000047Controller) Post() {
}
