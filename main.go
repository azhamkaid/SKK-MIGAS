package main

import (
	"fmt"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func main() {
	arrayString := [5]string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	//mengabungkan menjadi satu string
	joinString := strings.Join(arrayString[:], "")

	fmt.Println(joinString)
	value := make(map[string]int)
	//menghitung jumlah karakter dengan huruf besar
	for i := 65; i <= 90; i++ { //for menggunakan nilai rune dari A=65 hingga Z=90
		char := string(rune(i)) //rune yang di convert otomatis menjadi huruf besar
		result := strings.Count(joinString, string(char))
		//jika angka tidak ada maka result bernilai 0
		if result > 0 {
			fmt.Printf("Huruf '%s' Muncul Sebanyak %d Kali.\n", char, result)
			value[char] = result //input map dengan karakter dan jumlah kemunculan nya dalam kata
		}
	}
	//menghitung jumlah karakter dengan huruf kecil
	for i := 65; i <= 90; i++ { //for menggunakan nilai rune dari A=65 hingga Z=90
		Char := (rune(i))
		char := strings.ToLower(string(Char)) //convert huruf besar menjadi huruf kecil
		result := strings.Count(joinString, string(char))
		//jika angka tidak ada maka result bernilai 0
		if result > 0 {
			fmt.Printf("Huruf '%s' Muncul Sebanyak %d Kali.\n", char, result)
			value[char] = result //input map dengan karakter dan jumlah kemunculan nya dalam kata
		}
	}
	fmt.Println(value)
	sortedKeys := sortByValue(value)
	new := strings.Join(sortedKeys, "")
	fmt.Println("urutan berdasarkan jumlah kemunculan kata, jika jumlah sama urutkan huruf besar terlebih dahulu, lalu urutkan berdasarkan abjad:")
	fmt.Println(sortedKeys, new)
}

// funsi untuk mengurutkan map yang memiliki huruf dan jumlah kemunculan nya
func sortByValue(m map[string]int) []string {
	var items []kv
	for key, Count := range m {
		items = append(items, kv{key, Count})
	}

	sort.Slice(items, func(i, j int) bool {

		if items[i].Value == items[j].Value {
			// jika nilai sama dahulukan huruf besar
			if items[i].Key < items[j].Key {
				return true
			} else if items[i].Key > items[j].Key {
				return false
			}
			// jika nilai sama urutkan berdasarkan alfabet
			return items[i].Key > items[j].Key
		}
		return items[i].Value > items[j].Value
	})
	var sortedKeys []string
	for _, item := range items {
		sortedKeys = append(sortedKeys, item.Key)
	}

	return sortedKeys
}
