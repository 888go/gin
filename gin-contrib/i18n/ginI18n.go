package i18n

import (
	"errors"
	"fmt"
	"path"
	
	"github.com/888go/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ GinI18n = (*ginI18nImpl)(nil)

type ginI18nImpl struct {
	bundle          *i18n.Bundle
	localizerByLng  map[string]*i18n.Localizer
	defaultLanguage language.Tag
	getLngHandler   GetLngHandler
}

// getMessage 通过 lng（语言代码）和 messageID 获取本地化消息
func (i *ginI18nImpl) getMessage(ctx *gin.Context, param interface{}) (string, error) {
	lng := i.getLngHandler(ctx, i.defaultLanguage.String())
	localizer := i.getLocalizerByLng(lng)

	localizeConfig, err := i.getLocalizeConfig(param)
	if err != nil {
		return "", err
	}

	message, err := localizer.Localize(localizeConfig)
	if err != nil {
		return "", err
	}

	return message, nil
}

// mustGetMessage ... 必须获取消息...
func (i *ginI18nImpl) mustGetMessage(ctx *gin.Context, param interface{}) string {
	message, _ := i.getMessage(ctx, param)
	return message
}

func (i *ginI18nImpl) setBundle(cfg *BundleCfg) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(cfg.FormatBundleFile, cfg.UnmarshalFunc)

	i.bundle = bundle
	i.defaultLanguage = cfg.DefaultLanguage

	i.loadMessageFiles(cfg)
	i.setLocalizerByLng(cfg.AcceptLanguage)
}

func (i *ginI18nImpl) setGetLngHandler(handler GetLngHandler) {
	i.getLngHandler = handler
}

// loadMessageFiles 加载所有本地化文件到资源包中
func (i *ginI18nImpl) loadMessageFiles(config *BundleCfg) {
	for _, lng := range config.AcceptLanguage {
		src := path.Join(config.RootPath, lng.String()) + "." + config.FormatBundleFile
		if err := i.loadMessageFile(config, src); err != nil {
			panic(err)
		}
	}
}

func (i *ginI18nImpl) loadMessageFile(config *BundleCfg, src string) error {
	buf, err := config.Loader.LoadMessage(src)
	if err != nil {
		return err
	}

	if _, err = i.bundle.ParseMessageFileBytes(buf, src); err != nil {
		return err
	}
	return nil
}

// setLocalizerByLng 根据语言设置本地化器
func (i *ginI18nImpl) setLocalizerByLng(acceptLanguage []language.Tag) {
	i.localizerByLng = map[string]*i18n.Localizer{}
	for _, lng := range acceptLanguage {
		lngStr := lng.String()
		i.localizerByLng[lngStr] = i.newLocalizer(lngStr)
	}

// 如果默认语言不存在，则设置默认语言
	defaultLng := i.defaultLanguage.String()
	if _, hasDefaultLng := i.localizerByLng[defaultLng]; !hasDefaultLng {
		i.localizerByLng[defaultLng] = i.newLocalizer(defaultLng)
	}
}

// newLocalizer 通过语言创建一个本地化器
func (i *ginI18nImpl) newLocalizer(lng string) *i18n.Localizer {
	lngDefault := i.defaultLanguage.String()
	lngs := []string{
		lng,
	}

	if lng != lngDefault {
		lngs = append(lngs, lngDefault)
	}

	localizer := i18n.NewLocalizer(
		i.bundle,
		lngs...,
	)
	return localizer
}

// getLocalizerByLng 通过语言获取本地化器
func (i *ginI18nImpl) getLocalizerByLng(lng string) *i18n.Localizer {
	localizer, hasValue := i.localizerByLng[lng]
	if hasValue {
		return localizer
	}

	return i.localizerByLng[i.defaultLanguage.String()]
}

func (i *ginI18nImpl) getLocalizeConfig(param interface{}) (*i18n.LocalizeConfig, error) {
	switch paramValue := param.(type) {
	case string:
		localizeConfig := &i18n.LocalizeConfig{
			MessageID: paramValue,
		}
		return localizeConfig, nil
	case *i18n.LocalizeConfig:
		return paramValue, nil
	case i18n.LocalizeConfig:
		return &paramValue, nil
	}

	msg := fmt.Sprintf("un supported localize param: %v", param)
	return nil, errors.New(msg)
}
