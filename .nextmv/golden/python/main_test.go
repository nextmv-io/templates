package mip

import (
	"os"
	"testing"

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
				"-recipient",
				"nextmv",
			},
			TransientFields: []golden.TransientField{
				{
					Key:         ".statistics.result.duration",
					Replacement: golden.StableFloat,
				},
				{
					Key:         ".statistics.run.duration",
					Replacement: golden.StableFloat,
				},
				{
					Key:         ".statistics.result.custom.nextmv_version",
					Replacement: "v0.0.0",
				},
			},
			ExecutionConfig: &golden.ExecutionConfig{
				Command:    "python3",
				Args:       []string{"../../../python/main.py"},
				InputFlag:  "-input",
				OutputFlag: "-output",
			},
		},
	)
}
