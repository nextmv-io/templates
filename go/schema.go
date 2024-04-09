package main

type options struct {
	Recipient string `json:"recipient"`
}

// Input of the problem.
type input struct {
	Greeting string `json:"greeting"`
}

// solution represents the decisions made by the solver.
type solution struct {
	GreetingMessage string `json:"greeting_message,omitempty"`
}

type output struct {
	Solutions  []solution `json:"solutions"`
	Statistics statistics `json:"statistics"`
}

type statistics struct {
	Run    runInfo `json:"run,omitempty"`
	Result result  `json:"result,omitempty"`
	Schema string  `json:"schema,omitempty"`
}

type runInfo struct {
	Duration float64 `json:"duration,omitempty"`
}

type result struct {
	Value    float64 `json:"value,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Custom   any     `json:"custom,omitempty"`
}
