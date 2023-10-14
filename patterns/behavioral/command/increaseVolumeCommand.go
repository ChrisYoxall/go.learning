package main

type increaseVolumeCommand struct {
	device device
}

func (c *increaseVolumeCommand) execute() {
	c.device.increaseVolume()
}
