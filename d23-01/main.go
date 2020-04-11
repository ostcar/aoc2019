package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"

	"github.com/ostcar/aoc-2019/intcode"
)

func main() {
	code, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Can not read input: %v", err)
	}
	ns := newNetSwitch(string(code))
	for i := 0; i < 50; i++ {
		ns.start()
	}

	// go func() {
	// 	time.Sleep(time.Second)
	// 	var keys []int
	// 	for k := range ns.ports {
	// 		keys = append(keys, k)
	// 	}
	// 	sort.Ints(keys)
	// 	for _, key := range keys {
	// 		fmt.Println(key, ns.ports[key])
	// 	}
	// }()

	v := ns.distribute()
	fmt.Println(v)
}

type xy struct {
	x, y int
}

type netSwitch struct {
	portsMu     sync.Mutex
	ports       map[int][]xy
	code        string
	distributor map[int]chan int
	nextPort    int
}

func newNetSwitch(code string) *netSwitch {
	return &netSwitch{
		ports:       make(map[int][]xy),
		code:        code,
		distributor: make(map[int]chan int),
	}
}

func (n *netSwitch) start() {
	port := n.nextPort
	n.nextPort++
	n.distributor[port] = make(chan int)

	c := intcode.New(
		n.code,
		intcode.WithOutputChan(n.distributor[port]),
		intcode.WithInputFunc(n.getFunc(port)),
		intcode.WithName(strconv.Itoa(port)),
	)
	go c.Run()
}

func (n *netSwitch) getFunc(port int) func() int {
	var v xy
	var second bool // Send y
	var first bool  // Send port nr
	return func() int {
		if !first {
			first = true
			return port
		}

		if second {
			second = false
			return v.y
		}

		n.portsMu.Lock()
		defer n.portsMu.Unlock()

		if len(n.ports[port]) > 0 {
			second = true
			v = n.ports[port][0]
			n.ports[port] = n.ports[port][1:]
			return v.x
		}
		return -1
	}
}

func (n *netSwitch) distribute() int {
	type dist struct {
		port int
		xy   xy
	}

	values := make(chan dist)
	for i := range n.distributor {
		go func(i int) {
			for {
				port := <-n.distributor[i]
				x := <-n.distributor[i]
				y := <-n.distributor[i]
				values <- dist{port, xy{x, y}}
			}
		}(i)
	}

	for {
		v := <-values

		if v.port == 255 {
			return v.xy.y
		}

		n.portsMu.Lock()
		n.ports[v.port] = append(n.ports[v.port], v.xy)
		n.portsMu.Unlock()
	}
}
