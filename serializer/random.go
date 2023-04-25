package serializer

import (
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomCPUBrand() string{
	return randomStringFromSet("Apple", "NVA")
}

func randomCPUName(brand string) string{
	if brand == "Apple"{
		return randomStringFromSet("A4", "A5", "A5x", "A6")
	}
	return randomStringFromSet("Rayzen 7 pro", "Rayzen 8 pro", "Rayzen 9 pro")
}
func randomStringFromSet(a ...string) string{
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomInt(min, max int) int{
	return min + rand.Intn(max-min+1)
}

func randFlout64(min, max float64) float64{
	return min + rand.Float64() * (max - min)
}