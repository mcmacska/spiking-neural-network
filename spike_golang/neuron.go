package main

import (
	"fmt"
)

type Neuron struct {
	name              string
	membranePotential float32
	threshold         float32
	rest              float32
	leak              float32
	inputCurrent      float32
	signalCurrent     float32
	spiked            bool
	connectedNeurons  []*Neuron // store the pointers
}

func (thisN *Neuron) constructor(name string) {
	thisN.name = name
	thisN.membranePotential = 0.0
	thisN.threshold = 1.0
	thisN.rest = 0.0
	thisN.leak = 0.98
	thisN.signalCurrent = 0.4
	thisN.spiked = false
	thisN.inputCurrent = 0.0
}

func (thisN *Neuron) step() {
	//  add input
	thisN.membranePotential += thisN.inputCurrent

	if thisN.membranePotential >= thisN.threshold {
		thisN.spiked = true
		//  reset
		thisN.membranePotential = thisN.rest
	} else {
		thisN.spiked = false
	}
	//  leak
	thisN.membranePotential *= thisN.leak

	//  reset input
	thisN.inputCurrent = 0.0
}

func (thisN *Neuron) fire(current float32) {
	for i := 0; i < len(thisN.connectedNeurons); i++ {
		var n *Neuron = thisN.connectedNeurons[i]
		n.inputCurrent += current
	}
}

func (thisN *Neuron) connectNeuron(n *Neuron) {
	thisN.connectedNeurons = append(thisN.connectedNeurons, n)
}

func (thisN *Neuron) printMembranePotential() {
	fmt.Printf("%s: %f\n", thisN.name, thisN.membranePotential)
}

func (thisN *Neuron) printInputCurrent() {
	fmt.Printf("%s: %f\n", thisN.name, thisN.inputCurrent)
}

func (thisN *Neuron) printConnectedNeurons() {
	fmt.Println(thisN.name, ": ", thisN.connectedNeurons)
}

func (thisN *Neuron) printFire() {
	fmt.Println(thisN.name, " fired")
}
