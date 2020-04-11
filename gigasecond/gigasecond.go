package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond accept an time as input and returns the input time plus 10 ^ 9 seconds
func AddGigasecond(t time.Time) time.Time {
	// Using .Add method to and the time.Second incremented using math.pow of 1,000,000,000 secs.
	return t.Add(time.Second * time.Duration(math.Pow(10, 9)))
}
