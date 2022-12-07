package controllers

import (
	"github.com/astaxie/beego"
)
// Query Micro Loan product full info
type PD000052Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Micro Loan product full info
// @Param	body		body 	models.SPD0000052I	"Input parameter"
// @Success 200  {object} models.SPD0000052O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000052Controller) Post() {
}
