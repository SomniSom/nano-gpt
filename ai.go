package nanogpt

import (
	"context"
	"net/http"
	"net/url"

	"github.com/SomniSom/goreq"
)

// Authentication
// Use either header:
// Authorization: Bearer YOUR_API_KEY
// x-api-key: YOUR_API_KEY

var gpt *NanoGpt

type NanoGpt struct {
	Auth    string
	BaseUrl *url.URL
	Client  *http.Client
}

func (g NanoGpt) Balance() (Balance, error) {
	return goreq.New[Balance](context.Background(), g.BaseUrl.String()).
		Method(http.MethodPost).
		Headers("x-api-key", g.Auth, "Content-Type", "application/json").
		Path("/api/check-balance").Fetch()
}

func (g NanoGpt) Message(r *Request) (Response, error) {
	g.BaseUrl.Path = "/api/v1/messages"
	return goreq.New[Response](context.Background(), g.BaseUrl.String()).
		Headers("x-api-key", g.Auth, "Content-Type", "application/json").
		BodyJson(r).
		Fetch()
}

func Init(client *http.Client, auth string) {
	gpt = new(NanoGpt)
	gpt.Client = client
	gpt.Auth = auth
	gpt.BaseUrl = &url.URL{
		// Default url: https://nano-gpt.com/api/v1/messages
		Scheme: "https",
		Host:   "nano-gpt.com",
	}
}
