package controller

import "github.com/gin-gonic/gin"

// FP013018 godoc
// @Summary     v3 request to pay
// @Description v3 request to pay
// @Tags        paymentSwitch
// @Param       FP013018Request body     models.FpsPain013 true "v3 request to pay"
// @Success     200                        {object} controller.Response{response=models.FpsPain014}
// @Failure     500                {object} controller.Error
// @Router     /v3/nfps/request/to/pay [post]
func (c *Controller) FP013018(ctx *gin.Context) {

}
