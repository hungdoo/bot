package bybitapi

import (
	"encoding/json"
	"time"
)

// UnixTime is a custom type that embeds time.Time
type UnixTime struct {
	time.Time
}

// UnmarshalJSON satisfies the Unmarshaler interface for UnixTime
func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	ut.Time = time.UnixMilli(timestamp)
	return nil
}

// MarshalJSON satisfies the Marshaler interface for UnixTime
func (ut UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ut.Time.Format(time.RFC1123Z))
}
