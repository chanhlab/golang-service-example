package pkg

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToTimestampProto converts from time.Time to Protobuf Timestamp
func TimeToTimestampProto(time time.Time) *timestamp.Timestamp {
	timestampProto, _ := ptypes.TimestampProto(time)
	return timestampProto
}

// TimestampProtoNow gets the current Protobuf Timestamp
func TimestampProtoNow() *timestamp.Timestamp {
	return ptypes.TimestampNow()
}
