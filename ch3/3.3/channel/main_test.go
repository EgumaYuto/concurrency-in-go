package main

import "testing"

func TestSimplest(t *testing.T) {
	simplest()
}

func TestOpenSample(t *testing.T) {
	openChannel()
}

func TestCloseSample(t *testing.T) {
	closeChannel()
}

func TestFlashStream(t *testing.T) {
	flashChannel()
}

func TestFlashChannelByClosing(t *testing.T) {
	flashChannelByClosing()
}

func TestBufferedChannels(t *testing.T) {
	bufferedChannels()
}
