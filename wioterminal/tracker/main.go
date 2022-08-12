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
		ax := float64(x) / 1000000
		ay := float64(y) / 1000000
		az := float64(z) / 1000000

		// それぞれの軸に対するθを求める
		a1 := math.Atan2(ax, az)
		a2 := math.Atan2(ay, az)
		fmt.Printf("%f %f\r\n", a1, a2)

		time.Sleep(time.Millisecond * 100)
	}
}
