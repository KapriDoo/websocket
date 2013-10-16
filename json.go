// Copyright 2013 Gary Burd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
)

// WriteJSON writes the JSON encoding of v to the connection.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func WriteJSON(c *Conn, v interface{}) error {
	w, err := c.NextWriter(TextMessage)
	if err != nil {
		return err
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// See the documentation for the encoding/json Marshal function for details
// about the conversion of JSON to a Go value.
func ReadJSON(c *Conn, v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	return json.NewDecoder(r).Decode(v)
}
