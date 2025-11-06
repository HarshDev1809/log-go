package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	args := os.Args[1:]
	var filePath string
	lineCounts := 0
	if len(args) == 0 {
		for {
			fmt.Println("Kindly enter a valid file path :")
			argsCount, err := fmt.Scanf("%s", &filePath)
			if err == nil && argsCount > 0 {
				break
			}

		}
	}else{
		filePath = args[0]
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error while opening file : ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCounts ++
		line := scanner.Text()
		fmt.Printf("\033[90m%d\033[0m %s\n", lineCounts, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines : ",err)
	}

	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("Total Lines: ",lineCounts)

}
