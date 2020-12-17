// Getting the orders from your woocommerce api
//
// From Jonas Kwiedor <info@jj-ideeschmiede.de>
//
// In this file you can get all orders from the woocommerce api or all orders from a time period.
// All documentation is on my github profile github.com/jojojojonas/cmdotcom

package cmdotcom

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Config for new message
type Config struct {
	ProductToken    string
	Content         string
	Number          []string
	From            string
	AllowedChannels string
}

// Structs for sending data in json
type Request struct {
	Messages Messages `json:"messages"`
}

type Messages struct {
	Authentication Authentication `json:"authentication"`
	Message        []Message      `json:"msg"`
}

type Authentication struct {
	ProductToken string `json:"productToken"`
}

type Message struct {
	Body            Body     `json:"body"`
	To              []To     `json:"to"`
	From            string   `json:"from"`
	AllowedChannels []string `json:"allowed_channels"`
}

type Body struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type To struct {
	Number string `json:"number"`
}

// Structs for getting data and decode
type ReturnData struct {
	Details   string           `json:"details"`
	ErrorCode int              `json:"errorCode"`
	Messages  []MessagesReturn `json:"messages"`
}

type MessagesReturn struct {
	To               string      `json:"to"`
	Status           string      `json:"status"`
	Reference        interface{} `json:"reference"`
	Parts            int         `json:"parts"`
	MessageDetails   interface{} `json:"messageDetails"`
	MessageErrorCode int         `json:"messageErrorCode"`
}

// Variables
var (
	data ReturnData
)

// To send a new message via the cm.com api
func NewMessage(config Config) (ReturnData, error) {

	// Save numbers
	var to []To

	// Check numbers
	for _, value := range config.Number {

		// Add number to slice
		to = append(to, To{value})

	}

	// Create new message
	message := Request{
		Messages{
			Authentication{
				config.ProductToken,
			},
			[]Message{
				{
					Body{
						"auto",
						config.Content,
					},
					to,
					config.From,
					[]string{
						config.AllowedChannels,
					},
				},
			},
		},
	}

	// Prepare the data to marshal
	prepare, err := json.Marshal(message)
	if err != nil {
		return ReturnData{}, err
	}

	// Post data
	response, err := http.Post("https://gw.cmtelecom.com/v1.0/message", "application/json", bytes.NewBuffer(prepare))
	if err != nil {
		return ReturnData{}, err
	}

	// Decode the data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return ReturnData{}, err
	}

	// Return
	return data, nil

}
