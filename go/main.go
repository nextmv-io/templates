// package main is the starting point for your go template.
package main

import (
	"context"
	"log"
	"time"

	"github.com/nextmv-io/sdk/run"
)

func main() {
	err := run.CLI(solver).Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

// solver is the entrypoint of the program where a model is defined and solved.
func solver(_ context.Context, input input, options options) (output, error) {
	startTime := time.Now()
	// Build the greeting message.
	if options.Recipient == "" {
		options.Recipient = "world"
	}
	greetingMessage := input.Greeting + ", " + options.Recipient + "!"

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// Format the solution into the desired output format and add custom
	// statistics.
	output := format(greetingMessage, duration)

	return output, nil
}

// format the solution from the solver into the desired output format.
func format(greetingMessage string, duration time.Duration) output {
	return output{
		Solutions: []solution{
			{
				GreetingMessage: greetingMessage,
			},
		},
		Statistics: statistics{
			Run: runInfo{
				Duration: float64(duration.Seconds()),
			},
			Result: result{
				Value:    float64(len(greetingMessage)),
				Duration: float64(duration.Seconds()),
				Custom: map[string]string{
					"greeting_message": greetingMessage,
				},
			},
			Schema: "v1",
		},
	}
}
