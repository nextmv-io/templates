package mip

import (
	"os"
	"testing"
	"time"

	"github.com/nextmv-io/sdk/golden"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestGolden(t *testing.T) {
	golden.FileTests(
		t,
		"inputs",
		golden.Config{
			Args: []string{
				"-recipient", "nextmv",
			},
			TransientFields: []golden.TransientField{
				{Key: ".statistics.result.duration", Replacement: golden.StableFloat},
				{Key: ".statistics.run.duration", Replacement: golden.StableFloat},
			},
			Thresholds: golden.Tresholds{
				Float:    0.01,
				Time:     time.Duration(5) * time.Second,
				Duration: time.Duration(5) * time.Second,
			},
			ExecutionConfig: &golden.ExecutionConfig{
				Command:    "go",
				Args:       []string{"run", "."},
				InputFlag:  "-runner.input.path",
				OutputFlag: "-runner.output.path",
				WorkDir:    "../../../go",
			},
		},
	)
}
