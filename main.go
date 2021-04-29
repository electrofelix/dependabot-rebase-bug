package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	config := aws.Config{}

	fmt.Println("config:", config)
}
