package main

import (
	"fmt"
	"github.com/warthog618/gpio"
	"os"
	"time"
)

var (
// Pin Outline:       R1, B1, G1, R2, B2, G2
	pins [6]gpio.Pin
	// Outline	      A,  B,  C,  D,  E,  F,  G1,G2
	disp [8]gpio.Pin
  Cyclechan = make(chan bool)
)

func setup() {
	for i := 0; i < 6; i++ {
		pins[i].Output()
	}
	for j := 0; j < 8; j++ {
		disp[j].Output()
	}
	pins[5].High()
	pins[0].High()
}

func teardown() {
	for i := 0; i < 6; i++ {
		pins[i].Low()
	}
	for j := 0; j < 8; j++ {
		disp[j].Low()
	}
	fmt.Println("Teardown complete")
}

func cycle() {
  <-Cyclechan
  fmt.Println("Channel read")

	pins[5].Low()
	for i := 0; i < 3; i++{
		pins[4].High()
		time.Sleep(time.Second / 2)
		pins[4].Low()
		time.Sleep(time.Second / 2)
	}
	pins[3].High()
	pins[0].Low()
	pins[2].High()
	j := 9
	for j > 4 {
		numbers(j)
		time.Sleep(time.Second)
		j--
	}
	pins[2].Low()
	for j > 0 {
		numbers(j)
		pins[1].High()
		time.Sleep(time.Second / 2)
		pins[1].Low()
		time.Sleep(time.Second / 2)
		j--
	}
	numbers(j)
	pins[0].High()
	pins[3].Low()
	pins[5].High()

  time.Sleep(time.Second * 3)
  numbers(-1)
  time.Sleep(time.Second * 17)
}

func numbers(number int) {
	switch number {
	case 0:
		disp[0].High()
		disp[1].High()
		disp[2].High()
		disp[3].High()
		disp[4].High()
		disp[5].High()
		disp[6].Low()
		disp[7].Low()
	case 1:
		disp[0].Low()
		disp[1].High()
		disp[2].High()
		disp[3].Low()
		disp[4].Low()
		disp[5].Low()
		disp[6].Low()
		disp[7].Low()
	case 2:
		disp[0].High()
		disp[1].High()
		disp[2].Low()
		disp[3].High()
		disp[4].High()
		disp[5].Low()
		disp[6].High()
		disp[7].High()
	case 3:
		disp[0].High()
		disp[1].High()
		disp[2].High()
		disp[3].High()
		disp[4].Low()
		disp[5].Low()
		disp[6].High()
		disp[7].High()
	case 4:
		disp[0].Low()
		disp[1].High()
		disp[2].High()
		disp[3].Low()
		disp[4].Low()
		disp[5].High()
		disp[6].High()
		disp[7].High()
	case 5:
		disp[0].High()
		disp[1].Low()
		disp[2].High()
		disp[3].High()
		disp[4].Low()
		disp[5].High()
		disp[6].High()
		disp[7].High()
	case 6:
		disp[0].High()
		disp[1].Low()
		disp[2].High()
		disp[3].High()
		disp[4].High()
		disp[5].High()
		disp[6].High()
		disp[7].High()
	case 7:
		disp[0].High()
		disp[1].High()
		disp[2].High()
		disp[3].Low()
		disp[4].Low()
		disp[5].Low()
		disp[6].Low()
		disp[7].Low()
	case 8:
		disp[0].High()
		disp[1].High()
		disp[2].High()
		disp[3].High()
		disp[4].High()
		disp[5].High()
		disp[6].High()
		disp[7].High()
	case 9:
		disp[0].High()
		disp[1].High()
		disp[2].High()
		disp[3].High()
		disp[4].Low()
		disp[5].High()
		disp[6].High()
		disp[7].High()
	case -1:
		for i := 0; i < 8; i++ {
			disp[i].Low()
		}
	}
}

func handler(*gpio.Pin){
  Cyclechan <- true
  fmt.Println("Channel written")
}

func main() {
	if err := gpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer gpio.Close()

	var button = gpio.NewPin(21)

  // Pin Outline:       R1, B1, G1, R2, B2, G2
	pins = [6]gpio.Pin{*gpio.NewPin(17), *gpio.NewPin(22), *gpio.NewPin(27),
		*gpio.NewPin(18), *gpio.NewPin(24), *gpio.NewPin(23)}
	// Outline	      A,  B,  C,  D,  E,  F,  G1,G2
	disp = [8]gpio.Pin{*gpio.NewPin(13), *gpio.NewPin(19), *gpio.NewPin(25), *gpio.NewPin(16),
		*gpio.NewPin(20), *gpio.NewPin(26), *gpio.NewPin(6), *gpio.NewPin(12)}

	defer teardown()

  go cycle()

  button.Watch(gpio.EdgeRising, handler)

	setup()

  for {}
	/*
  for {
		but := button.Read()
		if but == gpio.High {
			cycle(pins, disp)
			time.Sleep(time.Second * 3)
			numbers(-1, disp)
			time.Sleep(time.Second * 17)
		}
	}
  */
}
