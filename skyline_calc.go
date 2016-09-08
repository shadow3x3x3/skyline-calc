package main

import "fmt"

const (
	notMatchSizeError = iota
	dominate          = iota
	beDominated       = iota
	nonDominate       = iota
)

type SkylineData struct {
	id         string
	attributes []float32
}

func (sd *SkylineData) dimSize() int {
	return len(sd.attributes)
}

func skylineCalc(data1 SkylineData, data2 SkylineData) int {
	if data1.dimSize() != data2.dimSize() {
		return notMatchSizeError
	}
	data2Attrs := data2.attributes

	flag := 0
	checkFlag := data1.dimSize()
	for index, value := range data1.attributes {
		switch {
		case value > data2Attrs[index]:
			flag++
		case value < data2Attrs[index]:
			flag--
		case value == data2Attrs[index]:
			checkFlag--
		}
	}

	if flag == checkFlag {
		return beDominated
	} else if flag == 0-checkFlag {
		return dominate
	}
	return nonDominate
}

func main() {
	b := SkylineData{id: "4", attributes: []float32{1, 3, 4, 5}}
	a := SkylineData{id: "2", attributes: []float32{2, 4, 6, 9}}
	fmt.Println(skylineCalc(a, b))
}
