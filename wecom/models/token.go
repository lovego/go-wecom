package models

import (
	"time"

	"github.com/lovego/go-wecom/commons/gorms"
)

type Token struct {
	gorms.Model
	Type      string
	CorpID    string
	ExpiredAt *time.Time
}
