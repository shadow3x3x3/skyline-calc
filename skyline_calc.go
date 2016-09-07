package skylineCalc

import "fmt"

const (
	dominate    = iota
	beDominated = iota
	nonDominate = iota
)

type skylineData struct {
	id         string
	attributes []float32
}

func (sd *skylineData) dimSize() {
	return len(sd.attributes)
}

func skylineCalc(data1 skylineData, data2 skylineData) int {
	if data1.dimSize() != data2.dimSize() {
		return nil
	}
	flag := data1.dimSize()
	fmt.println("The dim is %d", flag)
	return dominate
}

func main() {

}
