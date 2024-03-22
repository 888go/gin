package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	
	ginI18n "github.com/888go/gin/gin-contrib/i18n"
	"github.com/888go/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed i18n/localizeJSON/*
var fs embed.FS

func main() {
	// new gin engine
	gin类.X设置运行模式(gin类.X常量_运行模式_发布)
	router := gin类.X创建()

	// apply i18n middleware
	router.X中间件(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
		AcceptLanguage:   []language.Tag{language.English, language.German, language.Chinese},
		RootPath:         "./i18n/localizeJSON/",
		UnmarshalFunc:    json.Unmarshal,
// 在注释掉这一行后，使用defaultLoader
// 它将从文件中加载
		Loader: &ginI18n.EmbedLoader{
			FS: fs,
		},
	})))

	router.X绑定GET("/", func(ctx *gin类.Context) {
		ctx.X输出文本(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
	})

	router.X绑定GET("/:name", func(ctx *gin类.Context) {
		ctx.X输出文本(http.StatusOK, ginI18n.MustGetMessage(
			ctx,
			&i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": ctx.X取API参数值("name"),
				},
			}))
	})

	if err := router.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
