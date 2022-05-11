package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

func Meeting(s string) string {
	dataList := strings.Split(strings.ToUpper(s), ";")

	for i, p := range dataList {
		person := strings.Split(string(p), ":")
		dataList[i] = fmt.Sprintf("(%s, %s)", person[1], person[0])
	}
	sort.Strings(dataList)
	return strings.Join(dataList, "")
}

func main() {
	start := time.Now()
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))

	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
}
