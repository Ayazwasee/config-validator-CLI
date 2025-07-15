package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestValidateFlags(t *testing.T) {
	cmd := exec.Command("go", "run", "..", "validate", "--config=sample.env", "--json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run CLI: %v", err)
	}
	output := string(out)
	if !strings.Contains(output, "sample.env") {
		t.Error("Missing config path output")
	}
	if !strings.Contains(output, "true") {
		t.Error("Missing JSON output flag value")
	}
}
