package controller

import "github.com/gin-gonic/gin"

// FP000043 godoc
// @Summary     report log query
// @Description report log query
// @Tags        paymentSwitch
// @Param       FP000043Request body     models.LogSearchReq true "report log query"
// @Success     200                        {object} controller.Response{response=models.Pagination}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/report/log/query [post]
func (c *Controller) FP000043(ctx *gin.Context) {

}
