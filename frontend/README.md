# HEU-Wisedu

HEU-Wisedu是一款基于[哈尔滨工程大学教务选课系统](https://jwxk.hrbeu.edu.cn/)`API`、[Acrylic](https://github.com/only9464/Acrylic)`模版`开发的跨平台桌面应用程序，主要功能为查看课程信息、选课等。

> [!WARNING]
> 本软件完全免费且开源，可二次开发。仅供学习交流使用，请勿用于商业用途，一经发现，将拉黑、并依规向校方举报！！！
>

Github地址：[https://github.com/only9464/HEU-Wisedu](https://github.com/only9464/HEU-Wisedu) <br>
Atomgit地址：[https://atomgit.com/only9464/HEU-Wisedu](https://atomgit.com/only9464/HEU-Wisedu)<br>
Gitee地址：[https://gitee.com/only9464/HEU-Wisedu](https://gitee.com/only9464/HEU-Wisedu)<br>
欢迎各位同志提交[PR](https://github.com/only9464/HEU-Wisedu/pulls)，共同完善。
**BUG反馈**或者**功能建议**欢迎提交[Issues](https://github.com/only9464/HEU-Wisedu/issues)或者发邮件：[sky9464@qq.com](mailto:sky9464@qq.com)

# 一、优势
- [x] **开源免费**，无需付款，更适合HEU宝宝体质。(担心信息泄露的自己下载源代码编译程序,参考下文[**二次开发**](#五二次开发)部分)
- [x] **直接**调用[选课系统](https://jwxk.hrbeu.edu.cn/)API进行课程相关操作，省去繁杂加载，**高峰期选(qiang)课快人一步**（
- [x] 支持跨平台，支持Windows、MacOS、Linux 
- [x] 根据已修学分(查成绩)，选课更方便、快捷、具有目的性
- [x] 采用Golang的通道技术多个课程依次选(qiang)课、多线程多个课程同时选(qiang)课，自己设置间隔时间（操作简单，一看就会）
- [x] 界面简约美观(乐)，支持明暗双主题（暗色太拉胯了）

[更新日志](#八更新日志)

# 二、功能

- [X] 登录 `教务选课系统`、`教务管理系统(统一身份认证)`
- [X] 查看 `培养方案内课程`
- [X] 查看 `跨专业选修课`
- [X] 查看 `公选课`
- [x] 查看 `本批次已选课程`
- [x] 查看 `所有学期已选课程` 及其`成绩` (说白了，只能查到出成绩的课)
- [x] 查看 `已修学分` 
- [x] 退选
- [x] 自动选(qiang)课

[相关功能截图](#七功能截图)

# 三、下载

- [X] Windows[下载](https://ghproxy.heu.us.kg/https://raw.githubusercontent.com/only9464/HEU-Wisedu/master/build/bin/HEU-Wisedu.exe)
- [ ] MacOS  [下载](#3打包)
- [ ] Linux  [下载](#3打包)


MacOS、Linux 懒得编译了，**暂时不能直接下载可执行程序**，有需要的可以自己下载源码编译，可以参考下面的**二次开发[打包](#3打包)** 部分

# 四、支持作者

`赛博乞讨(bushi)：`

路过的各位同志们，大家好！我是共产主义接班人(从小,老师就这么告诉我的)，Vivo ￥0.01，助力我完成共产主义事业！

![1736260193604](image/README/1736260193604.png)

# 五、二次开发

## 1.安装依赖

**以下依赖按照顺序逐个安装即可：**

- [![Go](https://img.shields.io/github/go-mod/go-version/only9464/HEU-Wisedu?logo=go&label=Golang&color=00ADD8)](https://go.dev/)
- [![Wails](https://img.shields.io/github/v/release/wailsapp/wails?label=Wails&color=red&logo=wails)](https://wails.io)
- [![Node](https://img.shields.io/badge/Node.js-v20.12.2-green?logo=node.js)](https://nodejs.org/)
- [![npm](https://img.shields.io/badge/npm-v9.0.0-red?logo=npm)](https://www.npmjs.com/)
- [![Vue](https://img.shields.io/badge/dynamic/json?url=https://raw.githubusercontent.com/only9464/HEU-Wisedu/master/frontend/package.json&query=$.dependencies.vue&label=Vue&color=4FC08D&logo=vue.js)](https://vuejs.org/)

### Windows

暂无

### Mac

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

# 六、Star History

[![Star History Chart](https://api.star-history.com/svg?repos=only9464/HEU-Wisedu&type=Date)](https://star-history.com/#only9464/HEU-Wisedu&Date)

# 七、功能截图
![1736841700461](image/README/1736841700461.png)
![1736838920570](image/README/1736838920570.png)
![1736838696602](image/README/1736838696602.png)
# 八、更新日志


-    V0.0.3 【2025年1月14日15:16】<br>
    - 修复系统默认暗色导致表格字体显示不清楚的Bug<br>
    - 新增选(qiang)课模式：老实人模式、狂暴模式<br>
    - 更改默认选(qiang)课间隔时间：0.1s --> 0.275s<br>
    - 优化登录界面<br>

