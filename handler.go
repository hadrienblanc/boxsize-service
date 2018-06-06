package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrPayloadValidation will be used in case the box mensuration is negative
	ErrPayloadValidation = errors.New("The box mensuration need to be been greater than 0")

	// ErrInvalidNumberValidation is the error when we use the holy number 42.
	ErrInvalidNumberValidation = errors.New("The box mensuration can't have a 42 number")
)

func BoxSizeHandler(w http.ResponseWriter, r *http.Request) {
	var payload MyPayload

	if r.Body == nil {
		http.Error(w, "We received an empty request.", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "[json] An error occured while decoding the json format request.",
			http.StatusUnprocessableEntity)
		return
	}

	slice, err := BoxSize(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(slice)
	if err != nil {
		http.Error(w, "[Encoding] An Error occured", http.StatusInternalServerError)
		return
	}
}

type MyPayload struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Length int `json:"length"`
}

func (p MyPayload) validate() error {
	if p.Height < 1 || p.Width < 1 || p.Length < 1 {
		return ErrPayloadValidation
	}

	if p.Height == 42 || p.Width == 42 || p.Length == 42 {
		return ErrInvalidNumberValidation
	}

	return nil
}

type MyServerAnswer struct {
	Strings []string `json:"results"`
}

func BoxSize(payload MyPayload) (MyServerAnswer, error) {
	slice := []string{}

	errs := payload.validate()

	if errs != nil {
		return MyServerAnswer{Strings: slice}, errs
	}

	volume := payload.Height * payload.Width * payload.Length
	slice = append(slice, fmt.Sprintf("The box volume is %d cm³", volume))

	liters := volume / 1000
	slice = append(slice, fmt.Sprintf("It's also %d liters.", liters))

	return MyServerAnswer{Strings: slice}, nil
}
