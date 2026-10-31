package stringsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClip(t *testing.T) {
	for _, c := range []struct {
		name, s string
		max     int
		want    string
	}{{"empty", "", 2, ""}, {"zero", "abc", 0, ""}, {"negative", "abc", -1, ""}, {"equal", "abc", 3, "abc"}, {"large", "abc", 8, "abc"}, {"short", "abc", 2, "ab"}, {"unicode", "Привет 👋", 8, "Привет 👋"}, {"cyrillic", "Привет", 3, "При"}, {"emoji", "👋Go", 1, "👋"}} {
		t.Run(c.name, func(t *testing.T) { assert.Equal(t, c.want, Clip(c.s, c.max)) })
	}
}
