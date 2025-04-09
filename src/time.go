package main

import "time"

func convert(delay string) (int, error) {
	// Convert the delay string to a duration
	duration, err := time.ParseDuration(delay)
	if err != nil {
		return 0, err
	}

	// Convert the duration to seconds
	return int(duration.Seconds()), nil
}
