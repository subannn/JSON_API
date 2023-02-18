package main

import (
	pkgs "./devices"
)

func main() {

	pkgs.OpenDB()

	pkgs.RequestHTTP()
}

