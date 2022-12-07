package controllers

import (
	"github.com/astaxie/beego"
)
// Modify Loan interest strategy
type PD000045Controller struct {
	beego.Controller
}

// @Title
// @Description Modify Loan interest strategy
// @Param	body		body 	models.SPD0000045I	"Input parameter"
// @Success 200  {object} models.SPD0000045O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000045Controller) Post() {
}
