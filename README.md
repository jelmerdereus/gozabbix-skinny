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
    
    // get the Zabbix version
	zabbixResponse, err := zabbix.Call("apiinfo.version", []string{}, false)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("Version of Zabbix: %s\n", zabbixResponse.Result)
    
    // get details about the hosts
	zabbixResponse, err = zabbix.Call("host.get", []string{}, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hosts in Zabbix: %#v\n", zabbixResponse.Result)
}
```
