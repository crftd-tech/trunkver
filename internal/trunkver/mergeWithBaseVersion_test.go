package trunkver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseVersionWithV(t *testing.T) {
	assert.Equal(t,
		"1.2.3-20240522142548-gb4dc0d3-12345",
		MergeWithBaseVersion("v1.2.3", "20240522142548-gb4dc0d3-12345"),
		"Should remove the leading 'v' from the base version and replace the prerelease part with the trunkver",
	)
}

func TestBaseVersionWithoutV(t *testing.T) {
	assert.Equal(t,
		"1.2.3-20240522142548-gb4dc0d3-12345",
		MergeWithBaseVersion("1.2.3", "20240522142548-gb4dc0d3-12345"),
		"Should replace the prerelease part with the trunkver",
	)
}

func TestCoercesBaseVersionIntoSemver(t *testing.T) {
	assert.Equal(t,
		"1.0.0-20240522142548-gb4dc0d3-12345",
		MergeWithBaseVersion("1", "20240522142548-gb4dc0d3-12345"),
		"Should make the best out of 1",
	)
}
