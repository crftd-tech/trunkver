package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestApplicationVersion(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2024-05-22T16:25:48+02:00")
	if err != nil {
		panic(err)
	}
	assert.Equal(t,
		"20240522142548.0.0-gb4dc0d3-12345",
		formatTrunkver(now, "gb4dc0d3", "12345"),
		"Should create a semver-compatible version",
	)
}
