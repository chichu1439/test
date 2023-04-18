package controller

import "github.com/gin-gonic/gin"

// FP000062 godoc
// @Summary     values delete
// @Description values delete
// @Tags        paymentSwitch
// @Param       FP000062Request body     models.FP000062I true "values delete"
// @Success     200                        {object} controller.Response{response=models.FP000062O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/fields/values/delete [post]
func (c *Controller) FP000062(ctx *gin.Context) {

}
