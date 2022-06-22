package rob

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	//将路由封装
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handle HandlerFunc) {
	engine.router.addRouter(method, pattern, handle)
}

func (engine *Engine) Get(pattern string, handle HandlerFunc) {
	engine.addRouter("GET", pattern, handle)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}
