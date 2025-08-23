package loadbalancer

import "fmt"

type LoadBalancer interface {
	Configure(servers []string) error
	GetServer() (string, error)
}

const (
	AlgorithmRoundRobin    = "roundrobin"
	AlgorithmLowestLatency = "lowestlatency"
)

func GetLoadBalancer(algorithm string, servers []string) (LoadBalancer, error) {
	switch algorithm {
	case AlgorithmRoundRobin:
		return &RoundRobin{currentIndex: 0, servers: servers}, nil
	case AlgorithmLowestLatency:
		return &LowestLatency{servers: servers}, nil
	default:
		return nil, fmt.Errorf("unknown load balancing algorithm: %s", algorithm)
	}
}
