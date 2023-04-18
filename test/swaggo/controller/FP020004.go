package controller

import "github.com/gin-gonic/gin"

// FP020004 godoc
// @Summary     addressing inquiry
// @Description addressing inquiry
// @Tags        paymentSwitch
// @Param       FP020004Request body     models.FP020004I true "addressing inquiry"
// @Success     200                        {object} controller.Response{response=models.FP020004O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/addressing/inquiry [post]
func (c *Controller) FP020004(ctx *gin.Context) {

}
