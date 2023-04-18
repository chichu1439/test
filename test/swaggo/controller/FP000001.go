package controller

import "github.com/gin-gonic/gin"

// FP000001 godoc
// @Summary     participant inquiry
// @Description participant inquiry
// @Tags        paymentSwitch
// @Param       FP000001Request body     models.FP000001I true "participant inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000001O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-list/inquiry [post]
func (c *Controller) FP000001(ctx *gin.Context) {

}
