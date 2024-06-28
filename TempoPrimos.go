package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	nrSizes       = 6  // número de magnitudes dos valores primos
	nrPrimos      = 10 // número de valores primos para cada magnitude
	processadores = 6  //número de processadores utilizados
)

func main() {
	fmt.Print("Número de processadores: ", processadores)
	runtime.GOMAXPROCS(processadores) // utilizando um processador

	// valores primos com respectivamente 3, 6, 9, 13, 18 casas
	// use o programa AchaNPrimos para achar primos com determinado número de casas

	primos3 := [nrPrimos]int{101, 883, 359, 941, 983, 859, 523, 631, 181, 233}
	primos6 := [nrPrimos]int{547369, 669437, 683251, 610279, 851117, 655439, 937351, 419443, 128467, 316879}
	primos9 := [nrPrimos]int{550032733, 429415309, 109543211, 882936113, 546857209, 756170741, 699422809, 469062577, 117355333, 617320027}
	primos13 := [nrPrimos]int{7069402558433, 960246047869, 5738081989711, 5358141480883, 2569391599009, 4135462531597, 7807787948171, 130788041233, 2708131414819, 1571981553097}
	primos16 := [nrPrimos]int{2207749090466833, 9361721528139247, 2657959759011013, 3551950148669023, 3460183118669741, 5503892014624961, 4067979800826917, 7848969908399551, 2806933754138389, 5211072635754109}
	primos18 := [nrPrimos]int{383376390724197361, 882611655919772761, 533290385325847007, 17969611178168479, 903013501582628521, 541906710014517121, 281512690206248899, 403936627075987639, 775148726422474717, 942319117335957539}

	todosPrimos := [nrSizes][nrPrimos]int{primos3, primos6, primos9, primos13, primos16, primos18}

	for p := 0; p < nrSizes; p++ {
		fmt.Println(" ")
		fmt.Println("****** Valores avaliados:", todosPrimos[p])
		fmt.Println(" ")

		resSeq := contaPrimosSeq(todosPrimos[p])
		fmt.Println(" ")
		fmt.Println("------ Sequencial:")
		fmt.Println("Tempo:", resSeq)

		resConc := contaPrimosConc(todosPrimos[p])
		fmt.Println(" ")
		fmt.Println("------ Concorrente:")
		fmt.Println("Tempo:", resConc)
	}
}

func contaPrimosSeq(s [nrPrimos]int) time.Duration {
	start := time.Now()
	for i := 0; i < nrPrimos; i++ {
		if isPrime(s[i]) {
			fmt.Print(" .")
		}
	}
	return time.Since(start)
}

func contaPrimosConc(s [nrPrimos]int) time.Duration {
	start := time.Now()
	end := make(chan int)

	for i := 0; i < nrPrimos; i++ {
		go func(i int) {
			if isPrime(s[i]) {
				fmt.Print(" .")
			}
			end <- 1
		}(i)
	}

	for i := 0; i < nrPrimos; i++ {
		<-end
	}

	return time.Since(start)
}

// Is p prime?
func isPrime(p int) bool {
	if p <= 1 {
		return false
	}
	if p <= 3 {
		return true
	}
	if p%2 == 0 || p%3 == 0 {
		return false
	}
	i := 5
	for i*i <= p {
		if p%i == 0 || p%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
