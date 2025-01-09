package login

import (
	"context"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// 供前端直接调用的函数都写在这里

func (a *App) LoginToJwgl(username, password, ocrAPIURL string) (map[string]*http.Cookie, error) {
	return LoginToJwgl(username, password, ocrAPIURL)
}
