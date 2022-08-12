// 8-1 Wio Terminal Tracker
//
// Wio Terminal の傾き情報を出力します
// Processing と組み合わせて傾き情報をパソコンの画面に同期します
//
// - Wio Terminal からパソコンに角度情報を送る
// - パソコンの Processing で書く動に対応して Gopher を傾ける
//
// tinygo flash --target wioterminal --size short ./wioterminal/tracker
//
// 書き込みが終わったらパソコンで Processing を立ち上げ
// メニュー＞ファイル＞開く から ./sketch_01/sketch_01.pde を読み込ませてください
//

package main

import (
	"fmt"
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

func main() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{SCL: machine.SCL0_PIN, SDA: machine.SDA0_PIN})

	accel := lis3dh.New(i2c)
	accel.Address = lis3dh.Address0 // address on the Wio Terminal
	accel.Configure()
	accel.SetRange(lis3dh.RANGE_2_G)

	println(accel.Connected())

	for {
		x, y, z, _ := accel.ReadAcceleration()

		// TODO: x, y, z を使って角θを求める
		a1 := 0
		a2 := 0

		fmt.Printf("%f %f\r\n", a1, a2)

		time.Sleep(time.Millisecond * 100)
	}
}
