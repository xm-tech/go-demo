package main

import "fmt"

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string
	Connecter
}

// try to implement USB interface
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect: ", pc.name)
}

type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println("Connected:", tv.name)
}

func DisConnect(usb interface{}) {
	// if pc, ok := usb.(PhoneConnecter); ok {
	// 	fmt.Println("DisConnected:", pc.name)
	// 	return
	// }
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("DisConnected: ", v.name)
	default:
		fmt.Println("UnKnow device")
	}
}

func main() {
	pc := PhoneConnecter{name: "PhoneConnecter"}
	var u Connecter
	u = Connecter(pc)
	u.Connect()

	DisConnect(u)
}
