package controllers

import (
	"github.com/mylxsw/aidea-server/internal/ai/chat"

	"github.com/mylxsw/aidea-server/api/auth"
	"github.com/mylxsw/aidea-server/config"
	"github.com/mylxsw/aidea-server/internal/helper"
	"github.com/mylxsw/glacier/web"
	"github.com/mylxsw/go-utils/array"
)

// ModelController 模型控制器
type ModelController struct {
	conf *config.Config
}

// NewModelController 创建模型控制器
func NewModelController(conf *config.Config) web.Controller {
	return &ModelController{conf: conf}
}

func (ctl *ModelController) Register(router web.Router) {
	router.Group("/models", func(router web.Router) {
		router.Get("/", ctl.Models)
	})
}

// Models 获取模型列表
func (ctl *ModelController) Models(ctx web.Context, client *auth.ClientInfo) web.Response {
	models := array.Filter(chat.Models(ctl.conf), func(item chat.Model, _ int) bool {
		if item.VersionMin != "" && helper.VersionOlder(client.Version, item.VersionMin) {
			return false
		}

		if item.VersionMax != "" && helper.VersionNewer(client.Version, item.VersionMax) {
			return false
		}

		return !(client.IsCNLocalMode(ctl.conf) && item.IsSenstiveModel())
	})

	return ctx.JSON(models)
}
