package core

import (
    "encoding/json"
)

type RestResponse struct {
    ErrorCode string        `json:"errorCode,omitempty"`
    Error     string        `json:"error,omitempty"`
    Errors    []string      `json:"errors,omitempty"`
    Message   string        `json:"message,omitempty"`
    Status    bool          `json:"status,omitempty"`
    Data      []interface{} `json:"data,omitempty"`
}

func (r *RestResponse) Serialize() []byte {
    str, err := json.Marshal(r)

    if err != nil {
        App.Log.Error(err)
    }

    return str
}