package controller

import "github.com/gin-gonic/gin"

// FP000068 godoc
// @Summary     scenario enum query
// @Description scenario enum query
// @Tags        paymentSwitch
// @Param       FP000068Request body     models.FP000068I true "scenario enum query"
// @Success     200                        {object} controller.Response{response=models.FP000068O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/enum/query [post]
func (c *Controller) FP000068(ctx *gin.Context) {

}
