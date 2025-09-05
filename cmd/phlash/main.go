package main

import (
	"fmt"
	"phlash/internal/device"
)

func main() {
	fmt.Println("phlash: ISO flashing utility")

	drives, err := device.GetDrives()
	if err != nil {
		fmt.Printf("Error detecting drives: %v\n", err)
		return
	}

	if len(drives) == 0 {
		fmt.Println("No USB drives detected.")
		return
	}

	fmt.Println("Available drives:")
	for _, drive := range drives {
		fmt.Println(drive)
	}
}
