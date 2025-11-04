/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"go_ndm.github.com/cmd"
	"go_ndm.github.com/device"
	_ "go_ndm.github.com/setup"
)

func main() {
	dm := device.GetManager()
	for _, d := range dm.GetDeviceList() {
		fmt.Println(d.Name + " - " + d.IPAddress)
	}
	cmd.Execute()

}
