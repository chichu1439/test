package controller

import "github.com/gin-gonic/gin"

// FP000061 godoc
// @Summary     values update
// @Description values update
// @Tags        paymentSwitch
// @Param       FP000061Request body     models.FP000061I true "values update"
// @Success     200                        {object} controller.Response{response=models.FP000061O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/fields/values/update [post]
func (c *Controller) FP000061(ctx *gin.Context) {

}
