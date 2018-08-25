package controllers

import (
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/speedwheel/mvc/web/bindata"
)

var (
	_bindataWebPublicIndexhtml []byte
	pages                      map[string]iris.Handler
	cacheMiddleware            = cache.StaticCache(720 * time.Hour)
	assetHandler               = iris.StaticEmbeddedHandler("./web/public", bindata.GzipAsset, bindata.GzipAssetNames, true)
)

func init() {
	_bindataWebPublicIndexhtml, _ = bindata.GzipAsset("web/public/index.html")
	pages = map[string]iris.Handler{
		//"/favicon.ico": htmlBindata(_bindataWebPublicImagesFaviconico),
		"/":          bindata.HtmlBindata(_bindataWebPublicIndexhtml),
		"index.html": bindata.HtmlBindata(_bindataWebPublicIndexhtml),
	}
}

type IndexController struct {
	Ctx iris.Context
}

func (c *IndexController) Get() {
	if h, ok := pages[c.Ctx.Path()]; ok {
		h(c.Ctx)
		return
	}
	for _, asset := range bindata.GzipAssetNames() {
		if c.Ctx.Path() == strings.Replace(asset, "web/public", "", 1) {
			cacheMiddleware(c.Ctx)
			cache.ETag(c.Ctx)
			assetHandler(c.Ctx)
			return
		}
	}
	pages["/"](c.Ctx)
}
