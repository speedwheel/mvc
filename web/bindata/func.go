package bindata

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func HtmlBindata(b []byte) iris.Handler {
	return func(ctx iris.Context) {
		context.AddGzipHeaders(ctx.ResponseWriter()) // import "github.com/kataras/iris/context"
		ctx.ContentType(context.ContentHTMLHeaderValue)
		ctx.Write(b)
	}
}
