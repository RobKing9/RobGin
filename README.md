# RobGin
用Golang实现的基于gin的web框架
<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

## 目录

- [项目介绍](#项目介绍)
- [文件目录说明](#文件目录说明)
- [技术栈](#技术栈)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [作者](#作者)
- [鸣谢](#鸣谢)


## 项目介绍
- 使用 Golang 实现，参考 Gin 实现的网络框架
- 设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持
- 使用 Trie 树实现动态路由( dynamic route )解析
- 实现路由分组控制( Route Group Control )
- 设计并实现 Web 框架的中间件(Middlewares)机制
- 实现静态资源服务( Static Resource)，支持HTML模板渲染。
## 文件目录说明
```
RobGin
|-- Day1 http-base  http基础
|   |-- main.go
|   |-- rob
|   |   `-- rob.go
|-- Day2-context   上下文
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- rob.go
|   |   `-- router.go
|-- Day3-trie   前缀树路由
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- rob.go
|   |   |-- router.go
|   |   |-- router_test.go
|   |   `-- trie.go
|-- Day4-groupcontrol   路由分组控制
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- rob.go
|   |   |-- router.go
|   |   `-- trie.go
|-- Day5-middleware   中间件
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- logger.go
|   |   |-- rob.go
|   |   |-- router.go
|   |   `-- trie.go
|-- Day6-template   模板
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- logger.go
|   |   |-- rob.go
|   |   |-- router.go
|   |   `-- trie.go
|   |-- static
|   |   `-- css
|   |       `-- test.css
|   |-- templates
|   |   |-- arr.tmpl
|   |   |-- css.tmpl
|   |   `-- custom_func.tmpl
|-- Day7-errorRecovery   错误恢复
|   |-- go.mod
|   |-- main.go
|   |-- rob
|   |   |-- context.go
|   |   |-- logger.go
|   |   |-- recovery.go 
|   |   |-- rob.go
|   |   |-- rob_test.go
|   |   |-- router.go
|   |   |-- router_test.go
|   |   `-- trie.go
|-- README.md


```

### 技术栈
- [gin](https://gin-gonic.com/zh-cn/)

### 贡献者

请阅读**README.md** 查阅为该项目做出贡献的开发者。

### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


### 作者

QQ邮箱：2768817839@qq.com


*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt)

### 鸣谢


- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)

<!-- links -->
[your-project-path]:RobKing8/RobGin
[contributors-shield]: https://img.shields.io/github/contributors/RobKing9/RobGin.svg?style=flat-square
[contributors-url]: https://github.com/RobKing9/RobGin/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/RobKing9/RobGin.svg?style=flat-square
[forks-url]: https://github.com/RobKing9/RobGin/network/members
[stars-shield]: https://img.shields.io/github/stars/RobKing9/RobGin.svg?style=flat-square
[stars-url]: https://github.com/RobKing9/RobGin/stargazers
[issues-shield]: https://img.shields.io/github/issues/RobKing9/RobGin.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg
[license-shield]: https://img.shields.io/github/license/RobKing9/RobGin.svg?style=flat-square
[license-url]: https://github.com/RobKing9/RobGin/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/RobKing