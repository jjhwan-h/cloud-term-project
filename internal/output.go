package internal

import "fmt"

type option int

const (
	listInstance option = iota + 1
	availableZones
	startInstance
	availableRegions
	stopInstance
	createInstance
	rebootInstance
	listImages
	quit = 99
)

var menu = map[string]string{
	"listInstance":     "1. list instance",
	"availableZones":   "2. available zones",
	"startInstance":    "3. start instance",
	"availableRegions": "4. available regions",
	"stopInstance":     "5. stop instance",
	"createInstance":   "6. create instance",
	"rebootInstance":   "7. reboot instance",
	"listImages":       "8. list images",
	"quit":             "99. quit",
	"line":             "------------------------------------------------------------",
}

func printMenu() {
	str := fmt.Sprintf("%s\n%s\t\t%s\n%s\t\t%s\n%s\t\t%s\n%s\t\t%s\n%s\n%s",
		menu["line"],
		menu["listInstance"],
		menu["availableZones"],
		menu["startInstance"],
		menu["availableRegions"],
		menu["stopInstance"],
		menu["createInstance"],
		menu["rebootInstance"],
		menu["listImages"],
		menu["quit"],
		menu["line"])
	fmt.Println(str)
}
