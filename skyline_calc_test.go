package skylineCalc

import (
	"testing"
)

func Test_SkylineCalc_Dominate(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{2, 4, 5}}

	result, _ := skylineCalc(sd1, sd2)
	if result != dominate {
		t.Errorf("{1, 2, 3} should dominate {2, 4, 5}")
	}

}

func Test_SkylineCalc_NonDominate(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{2, 1, 4}}

	result, _ := skylineCalc(sd1, sd2)
	if result != nonDominate {
		t.Errorf("{1, 2, 3} should not dominate or be dominated by {2, 1, 4}")
	}

}

func Test_SkylineCalc_Be_Dominated(t *testing.T) {
	sd1 := SkylineData{"1", []float32{2, 4, 5}}
	sd2 := SkylineData{"2", []float32{1, 2, 3}}

	result, _ := skylineCalc(sd1, sd2)
	if result != beDominated {
		t.Errorf("{2, 4, 5} should be dominated by {1, 2, 3}")
	}

}

func Test_SkylineCalc_Error(t *testing.T) {
	sd1 := SkylineData{"1", []float32{1, 2, 3}}
	sd2 := SkylineData{"2", []float32{1, 5}}

	_, err := skylineCalc(sd1, sd2)
	if err == nil {
		t.Errorf("should occur Size Error ")
	}
}
