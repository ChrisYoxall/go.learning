package main

type device interface {
	on()
	off()
	increaseVolume()
	decreaseVolume()
}
