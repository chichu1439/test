package controller

import "github.com/gin-gonic/gin"

// FP000065 godoc
// @Summary     host list
// @Description host list
// @Tags        paymentSwitch
// @Param       FP000065Request body     models.FP000065I true "host list"
// @Success     200                        {object} controller.Response{response=models.FP000065O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/host/list [post]
func (c *Controller) FP000065(ctx *gin.Context) {

}
