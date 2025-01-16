//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

// Provide msgp for external types.
// If updating packages breaks this, update structs below.

//msgp:clearomitted
//msgp:tag json
//go:generate msgp -unexported

type cpuTimesStat struct {
	CPU       string  `json:"cpu"`
	User      float64 `json:"user"`
	System    float64 `json:"system"`
	Idle      float64 `json:"idle"`
	Nice      float64 `json:"nice"`
	Iowait    float64 `json:"iowait"`
	Irq       float64 `json:"irq"`
	Softirq   float64 `json:"softirq"`
	Steal     float64 `json:"steal"`
	Guest     float64 `json:"guest"`
	GuestNice float64 `json:"guestNice"`
}

type loadAvgStat struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

// NetDevLine is single line parsed from /proc/net/dev or /proc/[pid]/net/dev.
type procfsNetDevLine struct {
	Name         string `json:"name"`          // The name of the interface.
	RxBytes      uint64 `json:"rx_bytes"`      // Cumulative count of bytes received.
	RxPackets    uint64 `json:"rx_packets"`    // Cumulative count of packets received.
	RxErrors     uint64 `json:"rx_errors"`     // Cumulative count of receive errors encountered.
	RxDropped    uint64 `json:"rx_dropped"`    // Cumulative count of packets dropped while receiving.
	RxFIFO       uint64 `json:"rx_fifo"`       // Cumulative count of FIFO buffer errors.
	RxFrame      uint64 `json:"rx_frame"`      // Cumulative count of packet framing errors.
	RxCompressed uint64 `json:"rx_compressed"` // Cumulative count of compressed packets received by the device driver.
	RxMulticast  uint64 `json:"rx_multicast"`  // Cumulative count of multicast frames received by the device driver.
	TxBytes      uint64 `json:"tx_bytes"`      // Cumulative count of bytes transmitted.
	TxPackets    uint64 `json:"tx_packets"`    // Cumulative count of packets transmitted.
	TxErrors     uint64 `json:"tx_errors"`     // Cumulative count of transmit errors encountered.
	TxDropped    uint64 `json:"tx_dropped"`    // Cumulative count of packets dropped while transmitting.
	TxFIFO       uint64 `json:"tx_fifo"`       // Cumulative count of FIFO buffer errors.
	TxCollisions uint64 `json:"tx_collisions"` // Cumulative count of collisions detected on the interface.
	TxCarrier    uint64 `json:"tx_carrier"`    // Cumulative count of carrier losses detected by the device driver.
	TxCompressed uint64 `json:"tx_compressed"` // Cumulative count of compressed packets transmitted by the device driver.
}
