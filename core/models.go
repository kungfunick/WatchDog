package core

import (
    "github.com/ventu-io/go-shortid"
    "time"
)

type Model struct {
    ID        uint      `gorm:"primary_key" json:"-"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

// A Subject represents a service or website
type Subject struct {
    Model
    Domain  string  `json:"domain"`
    PingURI string  `json:"ping"`
    ExtId   string  `json:"extId"`
    Status  string  `json:"status"`
    Audits  []Audit `json:"audits"`
}

func (s *Subject) BeforeCreate() {
    s.ExtId, _ = shortid.Generate()
}

// An audit represents a single check
type Audit struct {
    Model
    Subject      Subject `json:"subject"`
    Result       bool    `json:"result"`
    ResponseTime int     `json:"responseTime"`
}