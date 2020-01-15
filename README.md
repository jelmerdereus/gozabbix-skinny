# gozabbix-skinny
Skinny Go library for communicating with the Zabbix API

usage:

```go
import (
    gozabbix "github.com/jelmerdereus/gozabbix-skinny"
)

func main() {
	zabbix := gozabbix.ZabbixAPI("http://localhost")
	err := zabbix.Signin("zbx_api_user", "AP!Monkii")
	if err != nil {
		log.Fatal(err)
	}

	// get Zabbix version (withAuthentication: false)
	APIVersion, err := zabbix.Call("apiinfo.version", []string{}, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Zabbix version: %s\n\n", APIVersion.Result)

	// get Zabbix host details
	zabbixHosts, err := zabbix.Call("host.get", []string{}, true)
	if err != nil {
		log.Fatal(err)
	}

	if zabbixHostList, ok := zabbixHosts.Result.([]interface{}); ok {
		fmt.Println("Hosts in Zabbix:")

		for _, host := range zabbixHostList {
			if hostMap, ok := host.(map[string]interface{}); ok {
				fmt.Printf("id: %s | name: %s | status: %s\n", hostMap["hostid"], hostMap["host"], hostMap["status"])
			}
		}
	}
}
```
