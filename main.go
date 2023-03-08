package main

import (
	"log"
	"strings"
)

const STR string = `A  B  C\K  D  E
F  G  H  I  J
L  M  N  O  P
Q  R  S  T  U
V  W  X  Y  Z`

func decryption(input string) string {
	input = strings.ToUpper(input)
	data := [][]string{}
	temp := strings.Split(STR, "\n")
	for i, _ := range temp {
		line := strings.Split(temp[i], "  ")
		lineData := []string{}
		for j, _ := range line {
			lineData = append(lineData, line[j])
		}
		data = append(data, lineData)
	}
	str := ""

	if strings.Contains(input, ".") {
		data2 := strings.Split(input, " ")
		for i := 0; i < len(data2); i = i + 2 {
			x := len(data2[i]) - 1
			y := len(data2[i+1]) - 1
			temp := data[x][y]

			str += string(temp[0])
		}
		return str
	}

	for ii, _ := range input {
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				temp := data[i]
				if strings.Contains(temp[j], string(input[ii])) {
					x := strings.Repeat(".", i+1)
					y := strings.Repeat(".", j+1)
					if len(str) > 0 {
						str += " "
					}
					str += x + " " + y
				}
			}
		}
	}
	return str
}
func main() {
	log.Println(decryption(".. ... . . . ... . ..."))
}
