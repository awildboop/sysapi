package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/mem"
)

type SwapStatistics struct {
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
}

type MemoryStatistics struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Cached      uint64
}

type MemoryResponse struct {
	ResponseCode    uint16
	ResponseMessage string
	Swap            SwapStatistics
	Memory          MemoryStatistics
}

func MemHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	swapMem, err := mem.SwapMemory()
	if err != nil {
		if err != nil {
			res.WriteHeader(500)
			response := struct {
				ResponseCode    int
				ResponseMessage string
			}{
				ResponseCode:    http.StatusInternalServerError,
				ResponseMessage: "Error while gathering swap statistics.",
			}
			json.NewEncoder(res).Encode(response)
			return
		}

	}

	swapStats := &SwapStatistics{
		Total:       swapMem.Total,
		Used:        swapMem.Used,
		Free:        swapMem.Free,
		UsedPercent: swapMem.UsedPercent,
	}

	virtMem, err := mem.VirtualMemory()
	if err != nil {
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			response := struct {
				ResponseCode    int
				ResponseMessage string
			}{
				ResponseCode:    http.StatusInternalServerError,
				ResponseMessage: "Error while gathering memory statistics.",
			}
			json.NewEncoder(res).Encode(response)
			return
		}

	}

	memStats := &MemoryStatistics{
		Total:       virtMem.Total,
		Available:   virtMem.Available,
		Used:        virtMem.Used,
		UsedPercent: virtMem.UsedPercent,
		Cached:      virtMem.Cached,
	}

	memResponse := &MemoryResponse{
		ResponseCode:    http.StatusOK,
		ResponseMessage: "With the exception of percents, all units are in bytes.",
		Swap:            *swapStats,
		Memory:          *memStats,
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(memResponse)
}
