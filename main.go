package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var registeredUsers = map[string]string{}

func main() {
	// 创建一个 Gin 实例
	r := gin.New()
	// 手动设置受信任的代理地址 以消除程序运行时的警告
	r.SetTrustedProxies([]string{})
	// 静态文件路由，用于提供 HTML、CSS、JavaScript 文件
	r.Static("/static", "./static")

	// 设置模板引擎和模板文件路径
	r.LoadHTMLGlob("templates/*")

	// 首页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 注册页面路由
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	// 注册请求处理路由
	r.POST("/register", handleRegister)

	// 登录页面路由
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	// 登录请求处理路由
	r.POST("/login", handleLogin)

	// 设置其他路由
	r.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "About page",
		})
	})

	// 启动服务器
	r.Run(":8080")
}

// 处理用户注册的函数
func handleRegister(c *gin.Context) {
	// 获取请求中的用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 在这里可以添加你的逻辑，如验证用户名是否已存在，对密码进行加密等
	// 检查用户名是否已存在
	if _, ok := registeredUsers[username]; ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名已存在",
		})
		return
	}
	// 注册新用户
	registeredUsers[username] = password

	// 返回注册成功的响应，包含跳转链接
	c.JSON(http.StatusOK, gin.H{
		"message":  "注册成功",
		"redirect": "/",
	})
}

// 处理用户登录的函数
func handleLogin(c *gin.Context) {
	// 获取请求中的用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 在这里可以添加你的逻辑，如验证用户名和密码是否匹配等
	// 检查用户名是否存在
	if userpass, ok := registeredUsers[username]; ok {
		// 验证密码是否匹配
		if userpass == password {
			// 登录成功
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
			})
			return
		}
	}

	// 用户名或密码不匹配，登录失败
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "用户名或密码错误",
	})
}
