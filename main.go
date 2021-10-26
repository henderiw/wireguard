package main

import (
	"fmt"

	"golang.zx2c4.com/wireguard/wgctrl"
)

func Initwireguard(privkey string) error {
	/*
		key, err := wgtypes.ParseKey(privkey)
		if err != nil {
			return err
		}
	*/
	wgclient, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wgclient.Close()

	devices, err := wgclient.Devices()
	if err != nil {
		return err
	}
	fmt.Println("Here are your ports:")
	for _, i := range devices {
		fmt.Println(i.ListenPort)
	}
	return err

}

func ListPorts() error {
	wgclient, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wgclient.Close()

	devices, err := wgclient.Devices()
	if err != nil {
		return err
	}
	fmt.Println("Here are your ports:")
	for _, i := range devices {
		fmt.Println(i.ListenPort)
	}
	return err
}

func main() {
	if err := ListPorts(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
