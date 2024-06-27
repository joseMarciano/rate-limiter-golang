package model

import (
	"time"
)

type TypeClient int

const (
	Ip     TypeClient = iota
	ApiKey            = iota
)

type Client struct {
	Id       string
	Type     TypeClient
	LastHit  time.Time
	IsLocked bool
}
