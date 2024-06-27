package model

import "golang.org/x/time/rate"

type ClientLimiter struct {
	Client *Client
	Limit  *rate.Limiter
}
