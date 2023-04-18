package controller

import "github.com/gin-gonic/gin"

// FP000011 godoc
// @Summary     rtgs-participant inquiry
// @Description rtgs-participant inquiry
// @Tags        paymentSwitch
// @Param       FP000011Request body     models.FP000011I true "rtgs-participant inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000011O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/rtgs-participant/inquiry [post]
func (c *Controller) FP000011(ctx *gin.Context) {

}
