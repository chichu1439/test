package controller

import "github.com/gin-gonic/gin"

// FP000051 godoc
// @Summary     scenario detail
// @Description scenario detail
// @Tags        paymentSwitch
// @Param       FP000051Request body     models.FP000051I true "scenario detail"
// @Success     200                        {object} controller.Response{response=models.FP000051O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/detail [post]
func (c *Controller) FP000051(ctx *gin.Context) {

}
