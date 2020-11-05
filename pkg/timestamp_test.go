package pkg

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

func TestTimeToTimestampProto(t *testing.T) {
	currentTime := time.Now()
	expected, _ := ptypes.TimestampProto(currentTime)

	assert.Equal(t, expected, TimeToTimestampProto(currentTime))
}

func TestTimestampProtoNow(t *testing.T) {
	assert.NotEmpty(t, TimestampProtoNow())
}
