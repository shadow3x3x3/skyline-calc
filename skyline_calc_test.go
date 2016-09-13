package skylineCalc

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
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

func generatorDatas(dim int) (sds []SkylineData) {
	rand.Seed(time.Now().UnixNano())

	for times := 0; times < 1000000; times++ {
		var randAttrs []float32
		for index := 0; index < dim; index++ {
			randAttrs = append(randAttrs, float32(rand.Intn(10000))/100)
		}
		sds = append(sds, SkylineData{strconv.Itoa(times), randAttrs})
		randAttrs = nil
	}
	return
}

func Benchmark_2_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(2)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_3_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(3)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_4_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(4)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_5_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(5)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_6_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(6)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_7_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(7)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_8_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(8)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}

func Benchmark_30_Dim(b *testing.B) {
	b.StopTimer()
	sds := generatorDatas(30)
	b.StartTimer()
	for i := 1; i < 1000000; i++ {
		skylineCalc(sds[0], sds[i])
	}
}
