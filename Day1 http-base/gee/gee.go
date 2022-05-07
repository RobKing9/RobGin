package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

//引擎 启动器

type Engine struct {
	//路由表 value是一个函数
	router map[string]HandlerFunc
}

//初始化引擎

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

//启动器 添加路由的功能
func (engine *Engine) addRouter(method string, pattern string, handle HandlerFunc) {
	// get-user
	k := method + "-" + pattern
	//添加到路由表
	engine.router[k] = handle
}

func (engine *Engine) Get(pattern string, handle HandlerFunc) {
	engine.addRouter("GET", pattern, handle)
}

func (engine *Engine) POST(pattern string, handle HandlerFunc) {
	engine.addRouter("POST", pattern, handle)
}

//启动器 添加启动服务的功能

func (engine *Engine) Run(addr string) (err error) {
	//engine的ServerHTTP实现了接口handle 此方法的第二个参数为接口handle
	return http.ListenAndServe(addr, engine)
}

//实现了handle接口
//ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应
//Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//路由表的key 方法+/路径
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
