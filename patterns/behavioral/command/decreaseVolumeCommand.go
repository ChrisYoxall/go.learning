package main

type decreaseVolumeCommand struct {
	device device
}

func (c *decreaseVolumeCommand) execute() {
	c.device.decreaseVolume()
}
