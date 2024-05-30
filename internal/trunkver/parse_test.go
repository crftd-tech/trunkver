package trunkver

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsesMajorVersion(t *testing.T) {
	refTime, _ := time.Parse(time.RFC3339, "2024-05-28T22:50:00+00:00")

	trunkVer, err := ParseTrunkVer("20240528225000.0.0-g5066c12-9277357275-1")
	assert.NoError(t, err)

	assert.Equal(t, refTime.UTC(), trunkVer.Timestamp)
	assert.Equal(t, SourceRef{CommitHash: "5066c12", ScmPrefix: "g", RawRef: "g5066c12"}, trunkVer.SourceRef)
	assert.Equal(t, "9277357275-1", trunkVer.BuildRef)
}

func TestParsesPrereleaseVersion(t *testing.T) {
	refTime, _ := time.Parse(time.RFC3339, "2024-05-28T22:50:00+00:00")

	trunkVer, err := ParseTrunkVer("1.0.0-20240528225000-g5066c12-9277357275-1")
	assert.NoError(t, err)

	assert.Equal(t, refTime.UTC(), trunkVer.Timestamp)
	assert.Equal(t, SourceRef{CommitHash: "5066c12", ScmPrefix: "g", RawRef: "g5066c12"}, trunkVer.SourceRef)
	assert.Equal(t, "9277357275-1", trunkVer.BuildRef)
}

func TestParsesRawSourceRefCorrrectly(t *testing.T) {
	refTime, _ := time.Parse(time.RFC3339, "2024-05-28T22:50:00+00:00")

	trunkVer, err := ParseTrunkVer("20240528225000.0.0-someOtherScm-9277357275-1")
	assert.NoError(t, err)

	assert.Equal(t, refTime.UTC(), trunkVer.Timestamp)
	assert.Equal(t, SourceRef{CommitHash: "", ScmPrefix: "", RawRef: "someOtherScm"}, trunkVer.SourceRef)
	assert.Equal(t, "9277357275-1", trunkVer.BuildRef)
}
