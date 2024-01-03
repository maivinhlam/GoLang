package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("dict.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	maps := make(map[string][][]string, 0)

	// could have multiple sheets
	sheets := f.GetSheetList()
	for _, sheetName := range sheets {
		d, err := f.GetRows(sheetName)
		if err != nil {
			fmt.Println("error reading sheet", sheetName, ":", err)
			return
		}
		maps[sheetName] = d
	}
	saveAsJSON(maps, "dict.json")
}

func saveAsJSON(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	empty := ""
	tab := "\t"

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	log.Println(buffer.String())

	file.WriteString(buffer.String())
	return nil

}
