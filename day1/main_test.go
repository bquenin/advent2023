package main

import "testing"
import "github.com/stretchr/testify/assert"

func Test_getCalibrationValue(t *testing.T) {
	cases := []struct {
		line string
		want int
	}{
		{"sixdddkcqjdnzzrgfourxjtwosevenhg9", 69},
	}
	for _, c := range cases {
		got := getCalibrationValue(c.line)
		assert.Equal(t, c.want, got)
	}
}
