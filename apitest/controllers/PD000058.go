package controllers

import (
	"github.com/astaxie/beego"
)
// Delete Loan Fee strategy
type PD000058Controller struct {
	beego.Controller
}

// @Title
// @Description  Delete Loan Fee strategy
// @Param	body		body 	models.SPD0000058I	"Input parameter"
// @Success 200  {object} models.SPD0000058O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000058Controller) Post() {
}
