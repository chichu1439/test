package controller

import "github.com/gin-gonic/gin"

// FP000017 godoc
// @Summary     participant agency relation deletion
// @Description participant agency relation deletion
// @Tags        paymentSwitch
// @Param       FP000017Request body     models.FP000017I true "participant agency relation deletion"
// @Success     200                        {object} controller.Response{response=models.FP000017O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-agency-relation/deletion [post]
func (c *Controller) FP000017(ctx *gin.Context) {

}
