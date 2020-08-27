package main

import (
	"fmt"

	"tinkitty.co.uk/gencloudinitfiles/iso9660"
)

// GenCloudInitFiles ...
func GenCloudInitFiles() {
	fmt.Printf("Cloud-Init file generator.\n")
	iso := iso9660.New()
	iso.Test()
}

func main() {
	GenCloudInitFiles()
}
