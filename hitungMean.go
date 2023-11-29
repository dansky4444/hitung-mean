package main

import "fmt"

var count = 0

func inputBilangan_1302220048(bil *int) {
	fmt.Scanln(bil)
	if *bil < 0 {
		inputBilangan_1302220048(bil)
	}
}

func stop_1302220048(bil int) bool {
	if bil == 0 {
		return true
	}
	return false
}

func hitung_1302220048(mean *float64, min, max *int) {
	var bil = 2
	inputBilangan_1302220048(&bil)
	if !stop_1302220048(bil) {
		if *min == 0 && *max == 0 {
			*max = bil
			*min = bil
		}
		if bil > *max && bil > *min {
			*min = *max
			*max = bil
		} else if bil < *min && bil < *max {
			*min = bil
		}
		*mean = (*mean + float64(bil))
		count += 1
		hitung_1302220048(mean, min, max)
	}
}

func main() {
	var mean float64
	var min, max int
	hitung_1302220048(&mean, &min, &max)
	fmt.Println(mean/float64(count), min, max, count)

}
