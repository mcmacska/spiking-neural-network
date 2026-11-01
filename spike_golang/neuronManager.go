package main

import (
	"fmt"
	"time"
)

func manageNeurons(stepTime int16) {
	fmt.Printf("%s\n", "Setting up SNN...")

	var allNeurons []*Neuron = []*Neuron{}

	var n1 Neuron
	n1.constructor("n1")
	var n2 Neuron
	n2.constructor("n2")
	var n3 Neuron
	n3.constructor("n3")
	var n4 Neuron
	n4.constructor("n4")

	n1.connectNeuron(&n4)
	n2.connectNeuron(&n4)
	n3.connectNeuron(&n4)
	n4.connectNeuron(&n2)
	n4.connectNeuron(&n1)

	// initial current
	n1.inputCurrent = n1.threshold
	n2.inputCurrent = n2.threshold
	n3.inputCurrent = n3.threshold

	allNeurons = append(allNeurons, &n1)
	allNeurons = append(allNeurons, &n2)
	allNeurons = append(allNeurons, &n3)
	allNeurons = append(allNeurons, &n4)

	fmt.Printf("%s\n", "Starting SNN...")
	for {
		stepNeurons(allNeurons)
		time.Sleep(time.Duration(stepTime) * time.Millisecond)
	}

}

func stepNeurons(allNeurons []*Neuron) {

	// phase 1
	for _, n := range allNeurons {
		n.step()
		n.printMembranePotential()
	}

	// phase 2
	for _, n := range allNeurons {
		if n.spiked {
			n.fire(n.signalCurrent)
			n.printFire()
		}
	}

	fmt.Println("")
}
