package main

import (
	"fmt"
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

var i2c = machine.I2C0

func main() {
	i2c.Configure(machine.I2CConfig{SCL: machine.SCL0_PIN, SDA: machine.SDA0_PIN})

	accel := lis3dh.New(i2c)
	accel.Address = lis3dh.Address0 // address on the Wio Terminal
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	println(accel.Connected())

	for {
		x, y, z, _ := accel.ReadAcceleration()
		xx := float64(x) / 1000000
		yy := float64(y) / 1000000
		zz := float64(z) / 1000000

		fmt.Printf("%f %f\r\n", math.Atan2(xx, zz), math.Atan2(yy, zz))

		time.Sleep(time.Millisecond * 100)
	}
}
