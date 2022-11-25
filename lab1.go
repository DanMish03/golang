package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Started")
	s := bufio.NewScanner(os.Stdin)

	col:= 0

	content := [][]string{}

	for s.Scan() {
		line := s.Text()
		row := strings.Split(line, ",")
		if col == 0{
			col = len(row)
		}
		if line == "" {
			break
		}
		if col != len(row){
			 log.Fatalf("Row has %d columns, but must have %d\n ", len(row), col)
		}
		content = append(content,row)
		sort.Slice(content, func(i, j int) bool { return content[i][0] < content[j][0]})
		
		
	}

	

	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	fmt.Println(content)
	fmt.Println("Finished")
}