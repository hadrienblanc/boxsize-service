package main

import (
	"fmt"
	"testing"
)

type Scenario struct {
	Payload       MyPayload
	Expected      []string
	ExpectedError error
}

func Test_BoxSize(t *testing.T) {
	scenarios := []Scenario{
		{
			ExpectedError: ErrPayloadValidation,
			Payload: MyPayload{
				Height: -12,
				Length: 2,
				Width:  2,
			},
		},
		{
			ExpectedError: ErrInvalidNumberValidation,
			Payload: MyPayload{
				Height: 42,
				Length: 2,
				Width:  2,
			},
		},
		{
			Payload: MyPayload{
				Height: 5,
				Length: 4,
				Width:  3,
			},
			Expected: []string{
				"The box volume is 60 cm³",
				"It's also 0 liters.",
			},
			ExpectedError: nil,
		},
	}

	for i, scenario := range scenarios {
		message := fmt.Sprintf("scenario #%d", (i + 1))

		response, err := BoxSize(scenario.Payload)
		for line := range scenario.Expected {
			if scenario.Expected[line] != response.Strings[line] {
				t.Fatalf("[response] Expected <%s> but got <%s>",
					scenario.Expected[line],
					response.Strings[line])
			}
		}

		if err != scenario.ExpectedError {
			t.Fatalf("[error][%s] Expected <%s> but got <%s>",
				message, scenario.ExpectedError, err)
		}
	}

}
