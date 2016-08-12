package randport

import "testing"

func TestGet(t *testing.T) {
	port := Get()
	if port <= 0 || port > 65535 {
		t.Fatalf("export port range 1-65535, got: %d", port)
	}
	t.Logf("%d", port)
}
