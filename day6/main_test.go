package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	assert.Equal(t, 5, ProccessSignal("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, ProccessSignal("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, ProccessSignal("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, ProccessSignal("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestGame2(t *testing.T) {
	assert.Equal(t, 19, ProccessMessage("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, ProccessMessage("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, ProccessMessage("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, ProccessMessage("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, ProccessMessage("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
