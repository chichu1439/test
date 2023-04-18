package controller

import "github.com/gin-gonic/gin"

// FP000059 godoc
// @Summary     step delete
// @Description step delete
// @Tags        paymentSwitch
// @Param       FP000059Request body     models.FP000059I true "step delete"
// @Success     200                        {object} controller.Response{response=models.FP000059O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/step/delete [post]
func (c *Controller) FP000059(ctx *gin.Context) {

}
