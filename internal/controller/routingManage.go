package controller

import (
	"GoWAFer/internal/service"
	"GoWAFer/internal/types"
	"GoWAFer/pkg/pagination"
	"GoWAFer/pkg/utils"
	"GoWAFer/pkg/utils/api_helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RoutingManageController 路由管理控制器
type RoutingManageController struct {
	routingManageService *service.RoutingManageService
}

// NewRoutingManageController 实例化路由管理控制器
func NewRoutingManageController(routingManageService *service.RoutingManageService) *RoutingManageController {
	return &RoutingManageController{routingManageService: routingManageService}
}

// AddRouting godoc
// @Summary 新增路由管理记录
// @Description 新增路由管理记录
// @Tags Routing
// @Accept json
// @Product json
// @Param types.AddRoutingRequest body types.AddRoutingRequest true "请求体"
// @Success 200 {object} api_helper.Response
// @Router /waf/api/v1/routing [post]
func (c *RoutingManageController) AddRouting(g *gin.Context) {
	var req types.AddRoutingRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.Error(types.ErrInvalidBody)
		return
	}

	// 检查路由格式是否正确
	if !utils.ValidateRouting(req.Path) {
		g.Error(types.ErrRoutingNotMatch)
		return
	}

	err := c.routingManageService.AddRouting(req.Path, req.IsBlack)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{Status: 0, Msg: "操作成功！"})
}

// GetRouting godoc
// @Summary 分页查询路由管理记录
// @Description 分页查询路由管理记录
// @Tags Routing
// @Produce json
// @Param query query string false "查询关键字"
// @Param isBlack query string false "是否为黑名单"
// @Param page query int false "页码"
// @Param perPage query int false "页面大小"
// @Success 200 {object} api_helper.Response
// @Router /waf/api/v1/routing [get]
func (c *RoutingManageController) GetRouting(g *gin.Context) {
	// 通过请求实例化分页结构体
	page := pagination.NewFromRequest(g)
	query := g.Query("query")
	isBlack := g.DefaultQuery("isBlack", "false")
	var isBlackBool bool
	if isBlack == "true" {
		isBlackBool = true
	} else {
		isBlackBool = false
	}
	page = c.routingManageService.GetRoutingWithPagination(page, isBlackBool, query)
	g.JSON(http.StatusOK, api_helper.Response{Status: 0, Msg: "success", Data: page})
}

// DeleteRouting godoc
// @Summary 删除路由管理记录
// @Description 删除路由管理记录
// @Tags Routing
// @Accept json
// @Product json
// @Param types.DeleteRoutingRequest body types.DeleteRoutingRequest true "请求体"
// @Success 200 {object} api_helper.Response
// @Router /waf/api/v1/routing [delete]
func (c *RoutingManageController) DeleteRouting(g *gin.Context) {
	var req types.DeleteRoutingRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.Error(types.ErrInvalidBody)
		return
	}

	err := c.routingManageService.DeleteRouting(req.Path, req.IsBlack)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, api_helper.Response{Status: 0, Msg: "操作成功！"})
}
