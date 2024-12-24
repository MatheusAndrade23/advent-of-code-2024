package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	data, err := os.Open("test.txt")

	if(err != nil){
		panic(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)

	var disk string

	for scanner.Scan() {
		line := scanner.Text()

		disk = line
	}

	originalDisk := []int{}
	originalDiskCopy := []int{}
	blockScanned := 0

	for index, bit := range disk {
		if(index % 2 != 0){
			intBit, err := strconv.Atoi(string(bit))

			if err != nil {
				panic(err)
			}

			freeSpace := make([]int, intBit)
			for i := 0; i < intBit; i++ {
    		freeSpace[i] = -1
			}

			originalDiskCopy = append(originalDiskCopy, freeSpace...)
			originalDisk = append(originalDisk, freeSpace...)
		} else {
			intBit, err := strconv.Atoi(string(bit))

			if err != nil {
				panic(err)
			}

			blockSpace := make([]int, intBit)
			for i := 0; i < intBit; i++ {
				blockSpace[i] = blockScanned
			}

			originalDiskCopy = append(originalDiskCopy, blockSpace...)
			originalDisk = append(originalDisk, blockSpace...)
			blockScanned++
		}
	}

	bitsChecked := make([]bool, len(originalDisk))
	
	
	for index, bit := range originalDisk {
		if bit == -1 {
			for i := len(originalDisk) - 1; i >= 0; i-- {
				if originalDisk[i] != -1 && !bitsChecked[i] {
					originalDisk[index] = originalDisk[i]
					originalDisk[i] = -1
					bitsChecked[index] = true
					break
				}
			}
		}
		bitsChecked[index] = true
	}

	fileSystemCheckSum := 0

	for index, bit := range originalDisk {
		if bit != -1 {
			multiplication := bit * index
			fileSystemCheckSum += multiplication
		}
	}

	secondFileSystemCheckSum := 0

	for index, bit := range originalDiskCopy {
		if bit != -1 {
			multiplication := bit * index
			secondFileSystemCheckSum += multiplication
		}
	}

	fmt.Println("First Challenge: ", fileSystemCheckSum)
}