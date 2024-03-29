//go:build go1.16
// +build go1.16

package i18n

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type server struct {
	*gin类.Engine
}

func newEmbedServer(middleware ...gin类.HandlerFunc) *server {
	server := &server{gin类.X创建()}
	server.X中间件(middleware...)

	server.X绑定GET("/", func(context *gin类.Context) {
		context.X输出文本(http.StatusOK, MustGetMessage(context, "welcome"))
	})

	server.X绑定GET("/:name", func(context *gin类.Context) {
		context.X输出文本(http.StatusOK, MustGetMessage(context, &i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.X取API参数值("name"),
			},
		}))
	})

	return server
}

func (s *server) request(lng language.Tag, name string) string {
	path := "/" + name
	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", path, nil)
	req.Header.Add("Accept-Language", lng.String())

	w := httptest.NewRecorder()
	s.ServeHTTP(w, req)

	return w.Body.String()
}

var (
	//go:embed testdata/localizeJSON/*
	fs embed.FS

	s = newEmbedServer(Localize(WithBundle(&BundleCfg{
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
		AcceptLanguage:   []language.Tag{language.English, language.German, language.Chinese},
		RootPath:         "./testdata/localizeJSON/",
		UnmarshalFunc:    json.Unmarshal,
// 在注释掉这一行后，使用defaultLoader
// 它将从文件中加载
		Loader: &EmbedLoader{fs},
	})))
)

func TestEmbedLoader(t *testing.T) {
	type args struct {
		lng  language.Tag
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{
				name: "",
				lng:  language.English,
			},
			want: "hello",
		},
		{
			name: "hello alex",
			args: args{
				name: "",
				lng:  language.Chinese,
			},
			want: "你好",
		},
		{
			name: "hello alex",
			args: args{
				name: "alex",
				lng:  language.English,
			},
			want: "hello alex",
		},
		{
			name: "hello alex german",
			args: args{
				name: "alex",
				lng:  language.Chinese,
			},
			want: "你好 alex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.request(tt.args.lng, tt.args.name)
			if got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
