package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	m := macaron.New()
	m.Use(macaron.Recovery())
	m.Use(session.Sessioner(session.Options{
		// Name of provider. Default is "memory".
		Provider: "memory",
		// Provider configuration, it's corresponding to provider.
		ProviderConfig: "",
		// Cookie name to save session ID. Default is "MacaronSession".
		CookieName: "MacaronSession",
		// Cookie path to store. Default is "/".
		CookiePath: "/",
		// GC interval time in seconds. Default is 3600.
		Gclifetime: 3600,
		// Max life time in seconds. Default is whatever GC interval time is.
		Maxlifetime: 3600,
		// Use HTTPS only. Default is false.
		Secure: false,
		// Cookie life time. Default is 0.
		CookieLifeTime: 0,
		// Cookie domain name. Default is empty.
		Domain: "",
		// Session ID length. Default is 16.
		IDLength: 16,
		// Configuration section name. Default is "session".
		Section: "session",
	}))
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "/go/src/app/templates",
	}))

	m.Use(func(ctx *macaron.Context) {
		log.WithFields(log.Fields{
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Info("Request logging")
		ctx.Next()
	})

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Redirect("/api", 302)
	})
	m.Get("/api", func(ctx *macaron.Context) {
		ctx.Data["Headline"] = "API Docs"
		ctx.HTML(200, "apidoc")
	})

	m.Group("/persons", func() {
		m.Get("/", overviewHandler)
		m.Get("/:id", detailHandler)
		m.Post("/", binding.Bind(Person{}), createHandler)
		m.Put("/:id", binding.Bind(Person{}), upgradeHandler)
		m.Patch("/:id", binding.Bind(Person{}), updateHandler)
		m.Delete("/:id", deleteHandler)
		m.Group("/:id/chapters", func() {
			m.Get("/:id", overviewHandler)
			m.Post("/new", overviewHandler)
			m.Put("/update/:id", overviewHandler)
			m.Delete("/delete/:id", overviewHandler)
		})
	})

	m.NotFound(func(ctx *macaron.Context) {
		log.WithFields(log.Fields{
			"error": ctx.Req.RequestURI,
		}).Error("Not found")
		ctx.JSON(404, map[string]string{"error": "Route not defined"})
	})

	m.Run()
}
