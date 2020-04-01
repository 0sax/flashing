package flashing

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// FlashMessage represents a flash message
type FlashMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Set encodes a FlashMessage struct as a json string to a cookie
func (fm *FlashMessage) Set(w http.ResponseWriter, cookieName string) error {
	// encode struct to json
	js, err := json.Marshal(fm)
	if err != nil {
		return err
	}

	c := &http.Cookie{Name: cookieName,
		Value: string(js)}
	http.SetCookie(w, c)

	return nil
}

// GetFlash checks the request for a cookie a cookie and returns its contents
// as a FlashMessage
func GetFlash(w http.ResponseWriter, r *http.Request, cookieName string) (*FlashMessage, error) {

	c, err := r.Cookie(cookieName)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, err
		default:
			return nil, errors.New("cookie not getting for unknown reason: " + err.Error())
		}
	}

	var fm *FlashMessage

	err = json.Unmarshal([]byte(c.Value), &fm)
	if err != nil {
		fmt.Println("Error decoding flash cookie value")
		return nil, err
	}
	// write empty cookie
	dc := &http.Cookie{Name: cookieName, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)

	// return cookie as flashMessageStruct
	return fm, nil
}
