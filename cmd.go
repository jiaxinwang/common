package common

import (
	"context"
	"os/exec"
	"time"
)

// Run ...
func Run(name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	bytes, err := cmd.CombinedOutput()
	switch ctx.Err() {
	case context.DeadlineExceeded:
		return bytes, ctx.Err()
	case nil:
	default:
		return bytes, ctx.Err()
	}
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}
