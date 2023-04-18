package controller

import "github.com/gin-gonic/gin"

// FP000060 godoc
// @Summary     values creation
// @Description values creation
// @Tags        paymentSwitch
// @Param       FP000060Request body     models.FP000060I true "values creation"
// @Success     200                        {object} controller.Response{response=models.FP000060O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/fields/values/creation [post]
func (c *Controller) FP000060(ctx *gin.Context) {

}
