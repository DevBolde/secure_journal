package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func errorHandler(ctx rweb.Context, errorMessage string, menufunc ...element.Component) error {
	b := element.NewBuilder()
	e := b.Ele
	t := b.Text

	e("html").R(
		e("head").R(
			e("title").R(
				t("My Journal"),
			),
			e("style").R(
				t("body {background-color: lightblue;} h1 a {text-decoration: none; color: inherit;}"),
			),
		),
		e("body").R(
			e("h1").R(
				t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
			),
			e("div").R(
				menufunc,
			),
			e("div").R(
				e("p", "style", "color: red").R(
					t(errorMessage),
				),
				e("a", "href", "/register").R(
					t("Try again"),
				),
			),
		),
	)
	return ctx.WriteHTML(b.String())
}
