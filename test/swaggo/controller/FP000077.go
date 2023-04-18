package controller

import "github.com/gin-gonic/gin"

// FP000077 godoc
// @Summary     alert inquiry
// @Description alert inquiry
// @Tags        paymentSwitch
// @Param       FP000077Request body     models.FP000077I true "alert inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000077O}
// @Failure     500                {object} controller.Error
// @Router     /v2/b2b/nfps/alert/inquiry [post]
func (c *Controller) FP000077(ctx *gin.Context) {

}
