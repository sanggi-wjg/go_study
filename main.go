package main

import (
	"fmt"
	"math"
	"time"
)

// func recordMetrics() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total",
// 		Help: "The total number of processed events",
// 	})
// )

//func main() {
//	var sample map[string]string
//	sample = make(map[string]string, 2)
//	sample["key1"] = "key2"
//	fmt.Println(sample)
//
//	for key, value := range sample {
//		fmt.Println(key, value)
//	}
//}

//func somethingMethod(mustTrue bool) (string, error) {

//	if mustTrue {
//		return "yes! It's true", nil
//	}
//	return "", errors.New("error : It's not true")
//}
//
//func showResult(res string, err error) {
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(res)
//	}
//}
//
//func main() {
//	res, err := somethingMethod(true)
//	showResult(res, err)
//
//	res, err = somethingMethod(false)
//	showResult(res, err)
//}

//func show(i int, c chan bool) {
//	fmt.Println(i)
//	c <- true
//}
//
//func main() {
//	var result []bool
//	channel := make(chan bool)
//
//	for i := 0; i <= 10; i++ {
//		go show(i, channel)
//	}
//	for i := 0; i <= 10; i++ {
//		result = append(result, <-channel)
//	}
//	fmt.Println(result)
//}

//type mySample struct {
//	name string
//	data map[int]string
//}
//
//func newMySample(name string) *mySample {
//	m := mySample{name: name}
//	m.data = map[int]string{}
//	return &m
//}
//
//func (m *mySample) setData(data map[int]string) {
//	m.data = data
//}
//
//func (m *mySample) getData() map[int]string {
//	return m.data
//}

//func main() {
//	sample := newMySample("SampleOne")
//	d := map[int]string{0: "DataOne", 1: "DataTwo"}
//	sample.setData(d)
//
//	fmt.Println(sample.name)
//	fmt.Println(sample.getData())
//}

const RANGE = 10000

func main() {
	start := time.Now()
	var result []float64
	channel := make(chan float64)

	for i := 0; i < RANGE; i++ {
		go goRoutine(i, channel)
	}
	for i := 0; i < RANGE; i++ {
		result = append(result, <-channel)
	}
	fmt.Printf("%s", time.Since(start))
}

func goRoutine(i int, c chan float64) {
	defer fmt.Println(i)
	res := math.Pow(float64(i), float64(i))
	c <- res
}

//type Something struct{}
type something struct{}

func NewSomething() *something {
	return &something{}
}
