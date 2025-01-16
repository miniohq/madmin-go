//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build linux
// +build linux

package madmin

import (
	"github.com/prometheus/procfs/sysfs"
)

func getCPUFreqStats() ([]CPUFreqStats, error) {
	fs, err := sysfs.NewFS("/sys")
	if err != nil {
		return nil, err
	}

	stats, err := fs.SystemCpufreq()
	if err != nil {
		return nil, err
	}

	out := make([]CPUFreqStats, 0, len(stats))
	for _, stat := range stats {
		out = append(out, CPUFreqStats{
			Name:                     stat.Name,
			CpuinfoCurrentFrequency:  stat.CpuinfoCurrentFrequency,
			CpuinfoMinimumFrequency:  stat.CpuinfoMinimumFrequency,
			CpuinfoMaximumFrequency:  stat.CpuinfoMaximumFrequency,
			CpuinfoTransitionLatency: stat.CpuinfoTransitionLatency,
			ScalingCurrentFrequency:  stat.ScalingCurrentFrequency,
			ScalingMinimumFrequency:  stat.ScalingMinimumFrequency,
			ScalingMaximumFrequency:  stat.ScalingMaximumFrequency,
			AvailableGovernors:       stat.AvailableGovernors,
			Driver:                   stat.Driver,
			Governor:                 stat.Governor,
			RelatedCpus:              stat.RelatedCpus,
			SetSpeed:                 stat.SetSpeed,
		})
	}
	return out, nil
}
