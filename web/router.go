package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var templ *template.Template

func ConfigGinRouter(router gin.IRoutes) {
	templ = template.Must(template.ParseGlob("./web/template/*.gohtml"))
	template.Must(templ.ParseGlob("./web/template/admin/*"))
	configGinStaticRouter(router)
	configMainRoute(router)
	return
}

func configGinStaticRouter(router gin.IRoutes) {
	router.Static("/assets", "./web/assets")
	router.Static("/dist", "./node_modules")
	router.StaticFile("/favicon.ico", "./web/assets/favicon.ico")
}

func configMainRoute(router gin.IRoutes) {
	router.GET("/about", AboutPage)
	router.GET("/admin", AdminPage)
	router.GET("/admin/create-course", CreateCoursePage)
	router.GET("/login", LoginPage)
	router.GET("/logout", LogoutPage)
	router.GET("/setPassword/:key", CreatePasswordPage)
	router.GET("/watch/:id", WatchPage)
	router.GET("/", MainPage)
}
