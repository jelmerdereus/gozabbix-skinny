# gozabbix-skinny
Skinny Go library for communicating with the Zabbix API

usage:

```go
import (
	"fmt"
	"log"

	gozabbix "github.com/jelmerdereus/gozabbix-skinny"
)

func main() {
	zabbix := gozabbix.ZabbixAPI("http://localhost")
	err := zabbix.Signin("zbx_api_user", "AP!Monkii")
	if err != nil {
		log.Fatal(err)
	}

	// get Zabbix version
	resp, err := zabbix.Call("apiinfo.version", []string{}, false)
	if err != nil {
		log.Fatal(err)
	}

	if APIVersion, ok := resp.Result.(string); ok {
		fmt.Printf("Version of Zabbix: %s\n\n", APIVersion)
	}

	// get Zabbix host details
	resp, err = zabbix.Call("host.get", []string{}, true)
	if err != nil {
		log.Fatal(err)
	}

	if hostList, ok := resp.Result.([]interface{}); ok {
		fmt.Println("Hosts in Zabbix:")

		for _, host := range hostList {
			if details, ok := host.(map[string]interface{}); ok {
				fmt.Printf("| id: %s | name: %s | status: %s |\n", details["hostid"], details["host"], details["status"])
			}
		}
	}
}
```
