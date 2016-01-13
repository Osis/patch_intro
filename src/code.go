package main

import (
	"fmt"
)

func main() {
	validate()
	deploy()
	smokeTests()
}

func validate() {
	fmt.Println("Validating...")
}

func deploy() {
	fmt.Println("Deploying...")
}

func smokeTests() {
	fmt.Println("Running Smoke Tests...")
}
