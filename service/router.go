package service

import (
	"github.com/gorilla/mux"
)
//返回指向mux的指针的函数。我们可以用作处理程序的路由器。
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
//迭代器遍历我们在routes.go中声明的路由，并将它们附加到router实例
	for _, route := range routes {
//附加每个路由，使用类似于生成器的模式设置每个路由。
		router.Methods(route.Method).
					Path(route.Pattern).
					Name(route.Name).
					Handler(route.HandlerFunc)
	}

	return router
}