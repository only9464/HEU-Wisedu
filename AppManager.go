package main

import (
	"HEU-Wisedu/apps/FAWKC"
	"HEU-Wisedu/apps/Settings"
	"HEU-Wisedu/apps/TJKC"
	"HEU-Wisedu/apps/XGKC"
	"HEU-Wisedu/apps/login"
	"HEU-Wisedu/apps/only9464"
	"context"
)

// 创建新的AppManager并注册所有app
func NewAppManager() *AppManager {
	return &AppManager{
		apps: []AppInterface{
			NewApp(),
			TJKC.NewApp(),
			FAWKC.NewApp(),
			XGKC.NewApp(),
			login.NewApp(),
			only9464.NewApp(),
			Settings.NewApp(),
			// 在这里添加新的app即可(嘻嘻)*****
		},
	}
}

///所有APP统一管理

// App接口定义所有app必须实现的方法
type AppInterface interface {
	Startup(ctx context.Context)
}

// 存储所有app实例
type AppManager struct {
	apps []AppInterface
}

// 统一处理Startup
func (am *AppManager) StartupHandler(ctx context.Context) {
	for _, app := range am.apps {
		app.Startup(ctx)
	}
}

// 获取需要绑定的接口列表
func (am *AppManager) GetBindings() []interface{} {
	bindings := make([]interface{}, len(am.apps))
	for i, app := range am.apps {
		bindings[i] = app
	}
	return bindings
}
