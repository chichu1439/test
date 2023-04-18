package controller

import "github.com/gin-gonic/gin"

// FP013028 godoc
// @Summary     adjust mockinfo
// @Description adjust mockinfo
// @Tags        paymentSwitch
// @Param       FP013028Request body     models.FP013028I true "adjust mockinfo"
// @Success     200                        {object} controller.Response{response=models.FP013028O}
// @Failure     500                {object} controller.Error
// @Router     /v3/nfps/adjust/mockinfo [post]
func (c *Controller) FP013028(ctx *gin.Context) {

}
