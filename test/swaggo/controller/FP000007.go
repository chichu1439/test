package controller

import "github.com/gin-gonic/gin"

// FP000007 godoc
// @Summary     participant relation list inquiry
// @Description participant relation list inquiry
// @Tags        paymentSwitch
// @Param       FP000007Request body     models.FP000007I true "participant relation list inquiry"
// @Success     200                        {object} controller.Response{response=models.FP000007O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-relation-list/inquiry [post]
func (c *Controller) FP000007(ctx *gin.Context) {

}
