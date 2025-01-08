# HEU-Wisedu

HEU-Wisedu是一款基于[哈尔滨工程大学教务选课系统](https://jwxk.hrbeu.edu.cn/)`API`、[Acrylic](https://github.com/only9464/Acrylic)`模版`开发的跨平台桌面应用程序，主要功能为查看课程信息、选课等。

> [!WARNING]
> 本软件完全免费且开源，可二次开发。仅供学习交流使用，请勿用于商业用途，一经发现，将拉黑、并依规向校方举报！！！

欢迎各位同志提交[PR](https://github.com/only9464/HEU-Wisedu/pulls)，共同完善。
**BUG反馈**或者**功能建议**欢迎提交[Issues](https://github.com/only9464/HEU-Wisedu/issues)或者发邮件：[sky9464@qq.com](mailto:sky9464@qq.com)

# 一、优势
- [x] **开源免费**，无需付款，更适合HEU宝宝体质。(担心信息泄露的自己下载源代码编译程序,参考下文**二次开发**部分)
- [x] 直接调用[选课系统](https://jwxk.hrbeu.edu.cn/)API进行课程相关操作，省去繁杂加载，高峰期选(qiang)课快人一步（
- [x] 支持跨平台，支持Windows、MacOS、Linux (dddd)
- [ ] 根据已修学分、培养方案配置显示规则，选课更方便、快捷、具有目的性
- [ ] 支持多线程高并发多个课程同时选课（操作简单，一看就会）
- [x] 界面简约美观(乐)，支持明暗双主题

# 二、功能

- [X] 登录
- [X] 查看 `培养方案内课程`
- [X] 查看 `跨专业选修课`
- [X] 查看 `公选课`
- [ ] 查看 `本批次已选课程`
- [ ] 查看 `已修学分` 
- [x] 选课
- [x] 退选
- [ ] 自动选课

相关功能截图： (先欠着)

# 三、支持平台

- [X] `Windows`
- [X] `MacOS`
- [X] `Linux`

# 四、下载

- [X] [Windows 下载](https://ghproxy.mioe.me/https://raw.githubusercontent.com/only9464/HEU-Wisedu/master/build/bin/HEU-Wisedu.exe)
- [ ] [MacOS 下载]()
- [ ] [Linux 下载]()

MacOS、Linux 懒得编译了，有需要的可以自己下载源码编译，可以参考 下面的 二次开发 部分

# 五、支持作者

`赛博乞讨(bushi)：`

路过的各位同志们，大家好！我是共产主义接班人(从小,老师就这么告诉我的)，Vivo ￥0.01，助力我完成共产主义事业！

![1736260193604](image/README/1736260193604.png)

# 六、二次开发

## 1.安装依赖

**以下依赖按照顺序逐个安装即可：**

- [![Go](https://img.shields.io/github/go-mod/go-version/only9464/HEU-Wisedu?logo=go&label=Golang&color=00ADD8)](https://go.dev/)
- [![Wails](https://img.shields.io/github/v/release/wailsapp/wails?label=Wails&color=red&logo=wails)](https://wails.io)
- [![Node](https://img.shields.io/badge/Node.js-v20.12.2-green?logo=node.js)](https://nodejs.org/)
- [![npm](https://img.shields.io/badge/npm-v9.0.0-red?logo=npm)](https://www.npmjs.com/)
- [![Vue](https://img.shields.io/badge/dynamic/json?url=https://raw.githubusercontent.com/only9464/HEU-Wisedu/master/frontend/package.json&query=$.dependencies.vue&label=Vue&color=4FC08D&logo=vue.js)](https://vuejs.org/)

### Windows

暂无

### Linux

- libgtk-3-dev
- libwebkit2gtk-4.0-dev
- libglib2.0-dev

所需执行命令(仅在ubuntu-20.04.6-amd64测试通过，其余自测)：

```bash
sudo apt update
sudo apt-get install libgtk-3-dev libwebkit2gtk-4.0-dev libglib2.0-dev
export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig:$PKG_CONFIG_PATH
```

> [!NOTE]
>
> 新版 `Linux`安装 `libwebkit2gtk-4.0-dev`编译应用时需要增加 `-tags webkit2_40`

### Mac

暂无

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

`更多信息请参考：`[Acrylic 二次开发](https://github.com/only9464/Acrylic#%E4%BA%8C%E4%BA%8C%E6%AC%A1%E5%BC%80%E5%8F%91)

# 七、Star History

[![Star History Chart](https://api.star-history.com/svg?repos=only9464/HEU-Wisedu&type=Date)](https://star-history.com/#only9464/HEU-Wisedu&Date)


# 八、更新日志

