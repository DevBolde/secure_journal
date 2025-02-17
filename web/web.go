package web

import (
	"fmt"
	"secure_journal/login"

	"github.com/rohanthewiz/rweb"
)

func InitWeb() (s *rweb.Server) {
	s = rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		},
	)

	head := "<head><title>My Journal</title>"
	head += "<style>body {background-color: lightblue;}</style></head>"
	pageStart := "<html>"
	pageEnd := "</html>"

	rootHandler := func(ctx rweb.Context) error {
		body := "<body><h1>My Journal</h1>" + RegisterMenu() + "</body>"
		page := pageStart + head + body + pageEnd
		fmt.Println(page)
		return ctx.WriteHTML(page)
	}

	s.Get("/", rootHandler)

	s.Get("/register", func(ctx rweb.Context) (err error) {
		body := "<body><h1>My Journal</h1>" + RegisterMenu() +
			`<p style="color: navy">Register</p>` +
			`<form action="/register" method="POST">
                <label for="username">Username:</label><br>
                <input type="text" name="username" id="username"><br>
                <label for="password">Password:</label><br>
                <input type="password" name="password" id="password"><br>
                <label for="password">Confirm Password:</label><br>
                <input type="password" name="confirm_password" id="confirm_password"><br>
                <input type="submit" value="Register">
            </form>` +
			"</body>"
		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		confirm_password := ctx.Request().FormValue("confirm_password")

		if password != confirm_password {
			errorBody := "<body><h1>My Journal</h1>" + RegisterMenu() +
				`<p style="color: red">Registration failed: Passwords don't match!</p>` +
				`<a href="/register">Try again</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}

		err = login.Register(username, password)
		if err != nil {
			// Return an error page instead of just the error
			errorBody := "<body><h1>My Journal</h1>" + RegisterMenu() +
				`<p style="color: red">Registration failed: ` + err.Error() + `</p>` +
				`<a href="/register">Try again</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}

		successMsg := `<div style="margin: 20px;"><p style="color: green">Registration successful!</p></div>`
		body := "<body><h1>My Journal</h1>" + successMsg + LogMenu() + "</body>"
		page := pageStart + head + body + pageEnd

		return ctx.WriteHTML(page)

	})

	s.Get("/login", func(ctx rweb.Context) (err error) {
		body := "<body><h1>My Journal</h1>" + RegisterMenu() +
			`<p style="color: navy">Login</p>` +
			`<form action="/login" method="POST">
                <label for="username">Username:</label><br>
                <input type="text" name="username" id="username"><br>
                <label for="password">Password:</label><br>
                <input type="password" name="password" id="password"><br>
                <input type="submit" value="Login">
            </form>` +
			"</body>"

		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Post("/login", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")

		err = login.Login(username, password)
		if err != nil {
			errorBody := "<body><h1>My Journal</h1>" + RegisterMenu() +
				`<p style="color: red">Login failed: ` + err.Error() + `</p>` +
				`<a href="/login">Try again</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}

		body := "<body><h1>My Journal</h1>" + JournalMenu() +
			`<p style="color: green">Welcome to your Journals!</p>` +
			"</body>"
		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Get("/my-journals", func(ctx rweb.Context) (err error) {
		body := "<body><h1>My Journal</h1>" +
			"<h2>Your Very First Journal Entry!</h2>" +
			"<p>I can do all things through christ who strengthens me!</p>" +
			"</body>"

		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Get("/log-out", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})
	s.Get("/delete-user", func(ctx rweb.Context) (err error) {
		body := "<body><h1>My Journal</h1>" + RegisterMenu() +
			`<p style="color: navy">Delete User</p>` +
			`<form action="/delete-user" method="POST">
                <label for="username">Username:</label><br>
                <input type="text" name="username" id="username"><br>
                <label for="password">Password:</label><br>
                <input type="password" name="password" id="password"><br>
                <input type="submit" value="Delete">
            </form>` +
			"</body>"
		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")

		err = login.Delete(username, password)
		if err != nil {
			errorBody := "<body><h1>My Journal</h1>" + RegisterMenu() +
				`<p style="color: red">Deletion failed: ` + err.Error() + `</p>` +
				`<a href="/login">Try again</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}
		successMsg := "<body><h1>My Journal</h1>" + RegisterMenu() +
			`<p style="color: green">Deletion successful!</p>` +
			"</body>"
		page := pageStart + head + successMsg + pageEnd
		return ctx.WriteHTML(page)
	})
	//initweb return
	return
}
