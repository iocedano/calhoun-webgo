package pages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"training/webgo/internal"
)

var Joke = internal.CreatePage("joke", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	resp, _ := http.Get("https://api.chucknorris.io/jokes/random")

	body, _ := io.ReadAll(resp.Body)
	var joke struct {
		IconURL string `json:"icon_url"`
		Value   string `json:"value"`
	}

	if err := json.Unmarshal(body, &joke); err != nil {
		fmt.Println(err)
	}

	return joke
})
