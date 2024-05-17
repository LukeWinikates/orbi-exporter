package orbi

import "strconv"

type Device struct {
	IP                string `json:"ip"`
	MAC               string `json:"mac"`
	ConnectionType    string `json:"contype"`
	AttachType        string `json:"attachtype"`
	DeviceType        string `json:"devtype"`
	DeviceTypeName    string `json:"devtype_name"`
	Model             string `json:"model"`
	Name              string `json:"name"`
	ActiveStatus      string `json:"accsta"`
	ConnectedOrbiName string `json:"conn_orbi_name"`
	ConnectedOrbiMAC  string `json:"conn_orbi_mac"`
	BackhaulStatus    string `json:"backhaul_sta"`
	LedState          string `json:"ledstate"`
	LedFunc           string `json:"led_func"`
	SyncBtn           string `json:"sync_btn"`
	Uprate            string `json:"uprate"`
	Downrate          string `json:"downrate"`
	ModuleName        string `json:"module_name"`
}

func (d Device) DownrateFloat() float64 {
	return floatOrNegative(d.Downrate)
}

func floatOrNegative(decimalString string) float64 {
	float, err := strconv.ParseFloat(decimalString, 64)
	if err != nil {
		return -1
	}
	return float
}
func (d Device) UprateFloat() float64 {
	return floatOrNegative(d.Uprate)
}

func (d Device) ActiveStatusFloat() float64 {
	return floatOrNegative(d.ActiveStatus)
}
