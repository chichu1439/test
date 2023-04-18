package controller

import "github.com/gin-gonic/gin"

// FP000100 godoc
// @Summary     aog wait
// @Description aog wait
// @Tags        paymentSwitch
// @Param       FP000100Request body     models.FP000100I true "aog wait"
// @Success     200                        {object} controller.Response{response=models.FP000100O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/simulato1/aog/wait [post]
func (c *Controller) FP000100(ctx *gin.Context) {

}
