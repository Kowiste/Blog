package main

import "time"

func main() {
	println("Begin")
	//go nostop()
	withChannel()
	//oneWayChannel()
	//timerChannel()
}

func nostop() {
	println("bmw")
}
func withChannel() {
	ch := make(chan string)
	go stop(ch)
	<-ch
	println("end")

}
func stop(ch chan string) {
	println("going to stop")
	ch <- ""
}

func oneWayChannel() {
	ch := make(chan bool)
	go onlyread(ch)
	ch <- true
	go onlyread(ch)
	ch <- false
	chb := make(chan byte)
	go onlywrite(chb)

	println("Finish with: ", <-chb)
}
func onlyread(ch <-chan bool) {
	println("ford")
	if <-ch {
		println("is true")
		return
	}
	println("is false")
}
func onlywrite(ch chan<- byte) {
	println("Going to write in 20 second")
	time.Sleep(20000)
	println("or not ")
	ch <- 33
}

func timerChannel() {
	kill := make(chan bool)
	go waiting(kill)
	for {
		println("write")
		if <-kill {
			println("bye")
			break
		}
		time.Sleep(500)
	}
}
func waiting(ch chan<- bool) {
	tm := time.NewTicker(5 * time.Second)
	t1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t1.C:
			println("1 second")
		case <-tm.C:
			ch <- true
			return
		}
	}

}
