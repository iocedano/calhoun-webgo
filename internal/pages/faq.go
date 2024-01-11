package pages

import (
	"net/http"
	"training/webgo/internal"
)

var Faq = internal.CreatePage("faq", func(ResponseWriter http.ResponseWriter, Request *http.Request) interface{} {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
	}

	return questions
})
