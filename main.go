package main

import (
	"fmt"
	"net/http"

	"training/webgo/internal/api"
	"training/webgo/internal/pages"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "text/html; charset=utf-8"))
		r.Method(http.MethodGet, "/", pages.Home)
		r.Method(http.MethodGet, "/contact", pages.Contact)
		r.Method(http.MethodGet, "/signup", pages.SignUp)
		r.Post("/signup", pages.SignUp.ServeHTTP)
		r.Method(http.MethodGet, "/faq", pages.Faq)
		r.Method(http.MethodGet, "/show-up/{username}", pages.ShowUp)
		r.Method(http.MethodGet, "/joke", pages.Joke)
		r.NotFound(pages.NotFoudPage.ServeHTTP)
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Post("/users", api.UsersSignUp.ServeHTTP)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	err := B()
// 	// TODO: Determine if the `err` variable is an `ErrNotFound`
// 	if errors.Unwrap(err) == ErrNotFound {
// 		fmt.Println("Error not found ")
// 	}
// 	fmt.Println(err.Error())
// 	fmt.Println(errors.Is(err, ErrNotFound))
// 	fmt.Println(errors.As(err, &ErrNotFound))
// }

// var ErrNotFound = errors.New("not found")

// func A() error {
// 	return ErrNotFound
// }

// func B() error {
// 	err := A()
// 	if err != nil {
// 		return fmt.Errorf("b: %w", err)
// 	}
// 	return nil
// }
