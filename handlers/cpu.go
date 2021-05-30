package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
)

type CPUInformation struct {
	VendorID  string
	Family    string
	Model     string
	Cores     uint8
	ModelName string
	Mhz       float64
	CacheSize int32
	Flags     []string
}

type CPUStatistics struct {
}

type CPUResponse struct {
	ResponseCode    uint16         // Any response codes as defined by IANA.
	ResponseMessage string         // Any desired response messages, optional.
	Information     CPUInformation // Information about the cpu
	Statistics      CPUStatistics  // Statistics about the cpu
}

func CPUUsage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	info, err := cpu.Info()
	if err != nil {
		res.WriteHeader(500)
		response := struct {
			ResponseCode    int
			ResponseMessage string
		}{
			ResponseCode:    http.StatusInternalServerError,
			ResponseMessage: "Error while gathering CPU information.",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	cpu0 := info[0]
	information := &CPUInformation{
		VendorID:  cpu0.VendorID,
		Family:    cpu0.Family,
		Model:     cpu0.Model,
		Cores:     uint8(len(info)),
		ModelName: cpu0.ModelName,
		Mhz:       cpu0.Mhz,
		CacheSize: cpu0.CacheSize,
		Flags:     cpu0.Flags,
	}

	json.NewEncoder(res).Encode(information)
}
