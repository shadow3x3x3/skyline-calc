package skylineCalc

import "errors"

const (
	skylineError = 0
	dominate     = iota
	beDominated
	nonDominate
)

// SkylineData records multiple variables' data
type SkylineData struct {
	label      string
	attributes []float32
}

func (sd *SkylineData) dimSize() int {
	return len(sd.attributes)
}

func skylineCalc(data1 SkylineData, data2 SkylineData) (int, error) {
	if data1.dimSize() != data2.dimSize() {
		return skylineError, errors.New("Size Not Match: Two attributes' size not match.")
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
		return beDominated, nil
	} else if flag == 0-checkFlag {
		return dominate, nil
	}
	return nonDominate, nil
}
