package zabbix

// For `HostinterfaceObject` field: `Main`
const (
	HostinterfaceMainNotDefault = 0
	HostinterfaceMainDefault    = 1
)

// For `HostinterfaceObject` field: `Type`
const (
	HostinterfaceTypeAgent = 1
	HostinterfaceTypeSNMP  = 2
	HostinterfaceTypeIPMI  = 3
	HostinterfaceTypeJMX   = 4
)

// For `HostinterfaceObject` field: `UseIP`
const (
	HostinterfaceUseipDNS = 0
	HostinterfaceUseipIP  = 1
)

// For `HostinterfaceObject` field: `Bulk`
const (
	HostinterfaceBulkDontUse = 0
	HostinterfaceBulkUse     = 1
)

// HostinterfaceObject struct is used to store hostinterface operations results
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostinterface/object#hostinterface
type HostinterfaceObject struct {
	InterfaceID int    `json:"interfaceid,omitempty"`
	DNS         string `json:"dns"`
	HostID      int    `json:"hostid,omitempty"`
	IP          string `json:"ip"`
	Main        int    `json:"main"` // has defined consts, see above
	Port        string `json:"port"`
	Type        int    `json:"type"`           // has defined consts, see above
	UseIP       int    `json:"useip"`          // has defined consts, see above
	Bulk        int    `json:"bulk,omitempty"` // has defined consts, see above

	// Items []ItemObject `json:"items,omitempty"` // not implemented yet
	Hosts []HostObject `json:"hosts,omitempty"`
}

// HostinterfaceGetParams struct is used for hostinterface get requests
//
// see: https://www.zabbix.com/documentation/2.4/manual/api/reference/hostinterface/get#parameters
type HostinterfaceGetParams struct {
	GetParameters

	HostIDs      []int `json:"hostids,omitempty"`
	InterfaceIDs []int `json:"interfaceids,omitempty"`
	ItemIDs      []int `json:"itemids,omitempty"`
	TriggerIDs   []int `json:"triggerids,omitempty"`

	// SelectItems SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

// Structure to store creation result
type hostinterfaceCreateResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// Structure to store deletion result
type hostinterfaceDeleteResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// HostinterfaceGet gets hostinterfaces
func (z *Context) HostinterfaceGet(params HostinterfaceGetParams) ([]HostinterfaceObject, int, error) {

	var result []HostinterfaceObject

	status, err := z.request("hostinterface.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostinterfaceCreate creates hostinterfaces
func (z *Context) HostinterfaceCreate(params []HostinterfaceObject) ([]int, int, error) {

	var result hostinterfaceCreateResult

	status, err := z.request("hostinterface.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}

// HostinterfaceDelete deletes hostinterfaces
func (z *Context) HostinterfaceDelete(hostinterfaceIDs []int) ([]int, int, error) {

	var result hostinterfaceDeleteResult

	status, err := z.request("hostinterface.delete", hostinterfaceIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}
