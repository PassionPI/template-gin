package docker

import "strings"

// Stat docker stats命令的输出解析格式
type Stat struct {
	ContainerID   string
	Name          string
	CPU           string
	Memory        string
	MemoryPercent string
	NetIO         string
	BlockIO       string
	Pid           string
}

// ParseStats 解析docker stats命令的输出
func ParseStats(stats string) []*Stat {
	raw := strings.Split(stats, "\n")
	slice := raw[1 : len(raw)-1]
	result := make([]*Stat, len(slice))

	for i, v := range slice {
		if (len(v)) < 1 {
			continue
		}
		items := strings.Split(v, "   ")
		for i, item := range items {
			items[i] = strings.TrimSpace(item)
		}
		result[i] = &Stat{
			ContainerID:   items[0],
			Name:          items[1],
			CPU:           items[2],
			Memory:        items[3],
			MemoryPercent: items[4],
			NetIO:         items[5],
			BlockIO:       items[6],
			Pid:           items[7],
		}
	}

	return result
}
