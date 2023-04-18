package controller

import "github.com/gin-gonic/gin"

// FP000020 godoc
// @Summary     alert-balance setting
// @Description alert-balance setting
// @Tags        paymentSwitch
// @Param       FP000020Request body     models.FP000020I true "alert-balance setting"
// @Success     200                        {object} controller.Response{response=models.FP000020O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert-balance/setting [post]
func (c *Controller) FP000020(ctx *gin.Context) {

}
