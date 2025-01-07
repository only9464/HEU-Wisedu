# Acrylic

Acrylic为Wails框架的Vue模板，用于开发桌面应用程序，主要主题色为半透明毛玻璃效果的白色，因而取名Acrylic。(实际上是因为懒得想名字 (bushi))


# 一、开发环境

## 1.安装依赖

**以下依赖按照顺序逐个安装即可：**

  <a href="https://nodejs.org/">
    <img src="https://img.shields.io/badge/Node.js-v20.12.2-green?logo=node.js" alt="Node">
  </a>
<br>
<a href="https://www.npmjs.com/">
  <img src="https://img.shields.io/badge/npm-v9.0.0-red?logo=npm" alt="npm">
</a>
<br>
  <a href="https://vuejs.org/">
    <img src="https://img.shields.io/badge/dynamic/json?url=https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/GlideWay/main/frontend/package.json&query=$.dependencies.vue&label=Vue&color=4FC08D&logo=vue.js" alt="Vue">
  </a>
<br>
  <a href="https://go.dev/">
    <img src="https://img.shields.io/github/go-mod/go-version/only9464/GlideWay?logo=go&label=Golang&color=00ADD8" alt="Go">
  </a>
<br>
  <a href="https://wails.io">
    <img src="https://img.shields.io/github/v/release/wailsapp/wails?label=Wails&color=red&logo=wails" alt="Wails">
  </a>
<br>

## 2.调试运行

**在项目的根目录下执行：**

```bash
wails dev
```

## 3.打包

**在项目的根目录下执行：**

```bash
wails build
```
# 二、二次开发

~~**以下操作中起名均是为了规范化项目代码，帮助新手快速上手。如果懂Wails、Vue框架原理,可自行修改调整**~~

## 1.快速开发`前端界面`

### 1.1在[/frontend/src/components/](./frontend/src/components/)文件夹下创建一个新的文件夹，作为你的`新应用`的组件文件夹，例如：`test`，作为`<新应用名>`

### 1.2在[/frontend/src/components/<新应用名>/](./frontend/src/components/<新应用名>/)文件夹下创建你的组件文件，自由发挥即可，最好必须有一个`<新应用名>.vue`，例如：`test.vue`

```vue
<template>
    <!-- 这里是组件的模板 -->
</template>

<script setup>
    <!-- 这里是组件的逻辑 -->
</script>

<style scoped>
    <!-- 这里是组件的样式 -->
</style>
```
### 1.3在[/frontend/src/components/views/](./frontend/src/components/views/)文件夹下创建你的组件对应的页面文件，最好是`<新应用名>View.vue`，例如：`testView.vue`
以下是示例的完整代码：
```vue
<template>
  <div>
    <<新应用名> />
  </div>
</template>

<script setup>
import <新应用名> from '../components/<新应用名>/<新应用名>.vue'; // 导入 <新应用名> 组件
</script>

<style scoped>
</style>
```

### 1.4在[/frontend/src/router/index.js](./frontend/src/router/index.js)中添加你的组件对应的页面文件，例如：`testView.vue`

**每开发一个新应用，都要在此添加这两行：**

```javascript
import FirstView from '../views/FirstView.vue'
import SecondView from '../views/SecondView.vue'
import <新应用名>View from '../views/<新应用名>View.vue'

<!-- 现有的代码...... -->

const routes = [
  { path: '/', component: FirstView }, // 第一个页面
  { path: '/second', component: SecondView }, // 第二个页面
  { path: '/<新应用名>', component: <新应用名>View }, // 新应用页面
]
```
### 1.5在[/frontend/src/components/nav/Sidebar.vue](./frontend/src/components/nav/Sidebar.vue)文件中的`<li>`标签中添加你的组件对应的导航栏文件，例如：`<新应用名>Nav.vue`

```vue
<template>
  <div class="sidebar-container">
    <div class="sidebar acrylic-effect">
      <div class="menu-container">
        <ul>
          <li>
            <router-link to="/" class="nav-link">
              <span>第一个界面</span>
            </router-link>
          </li>
          <li>
            <router-link to="/second" class="nav-link">
              <span>第二个界面</span>
            </router-link>
          </li>

          <!--  这里是新应用界面的导航-->
          <!-- <li>
            <router-link to="/<新应用名>" class="nav-link">
              <span><新应用名></span>
            </router-link>
          </li> -->

        </ul>
      </div>
    </div>
  </div>
</template>
```

## 2.快速开发`后端逻辑`

### 2.1在[/apps/](./apps/)文件夹下创建你的组件对应的应用文件夹`<新应用名>`，例如：`test`

### 2.2在[/apps/<新应用名>/](./apps/<新应用名>/)文件夹下创建你的组件对应的`<新应用名>.go`文件，例如：`test.go`以及固定名称`app.go`
**app.go：** ~~（参考[/apps/Template/app.go](./apps/Template/app.go)）~~


```go
package <新应用名>

import (
    "context"
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

func (s *App) <新应用名>(a, b int) int {
    return <新应用名>(a, b)
}

```


**<新应用名>.go：** ~~（参考[/apps/Template/Template.go](./apps/Template/Template.go)）~~

```go
package <新应用名>

import "fmt"

//以下都是供本包内的函数进行调用的
func <新应用名>(a, b int) int {
    fmt.Println("<新应用名>")
    return a / b
}


```

### 2.3在[/AppManager.go](./AppManager.go)文件中添加你的应用的app实例~~（实际上就是加一行）~~：

```go
// 创建新的AppManager并注册所有app
func NewAppManager() *AppManager {
    return &AppManager{
        apps: []AppInterface{
            NewApp(),
            First.NewApp(),
            Second.NewApp(),
            <新应用名>.NewApp(),
            // 在这里添加新的app即可(嘻嘻)*****
        },
    }
}
```

## 3.让你的前端js代码“调用”你的后端go代码

如果在`2.快速开发后端逻辑`中，你是在按照我的步骤进行的，那么在前端js代码中，想要“**调用**”后端的go函数就**非常简单**

只需要 使用 `window.go.<新应用名>.App.<新应用名中的某个函数名>` 即可

例如：

```javascript
window.go.First.App.First(1, 2)
```
