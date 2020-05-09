package controllers

type DeviceVersionMap struct {
	Device string
	Versions []string
}

func QueryDeviceInfo()  []DeviceVersionMap{
	return []DeviceVersionMap{
		{"鸿蒙",[]string{"0.1","0.2","0.3"}},
		{"伏羲",[]string{"1.1","1.2","1.3"}},
		{"欧拉",[]string{"2.1","2.2"}},
		{"VT200",[]string{}},
		{"VCT200",[]string{}},
	}
}
