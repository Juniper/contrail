package collector

import (
	"net"
	"os"
	"strings"

	"github.com/Juniper/contrail/pkg/version"
)

const typeModuleCPUStateTrace = "ModuleCpuStateTrace"

type payloadModuleCPUStateTrace struct {
	Name         string                   `json:"name"`
	ConfigNodeIP []string                 `json:"config_node_ip"`
	BuildInfo    *infoModuleCPUStateTrace `json:"build_info"`
	Deleted      bool                     `json:"deleted"`
}

type infoModuleCPUStateTrace struct {
	BuildVersion  string `json:"build-version"`
	BuildTime     string `json:"build-time"`
	BuildUser     string `json:"build-user"`
	BuildHostname string `json:"build-hostname"`
	BuildID       string `json:"build-id"`
	BuildNumber   string `json:"build-number"`
}

func (p *payloadModuleCPUStateTrace) Build() *Message {
	return &Message{
		SandeshType: typeModuleCPUStateTrace,
		Payload:     p,
	}
}

func ipToFQDN(ip string) string {
	names, err := net.LookupAddr(ip)
	if err == nil && len(names) > 0 {
		return strings.TrimSuffix(names[0], ".")
	}
	return ""
}

func addrToIPsFQDN(nodeAddr string) ([]string, string) {
	host, _, err := net.SplitHostPort(nodeAddr)
	if err == nil && len(host) > 0 && net.ParseIP(host) != nil {
		// IP address specified
		if name := ipToFQDN(host); len(name) > 0 {
			return []string{host}, name
		}
		return []string{host}, "unknown"
	}
	// Test configuration may not contain IP address, try to collect node IPs
	if len(host) == 0 {
		host, _ = os.Hostname() // nolint: errcheck
	}
	ips, _ := net.LookupIP(host) // nolint: errcheck
	nodeIPs := []string{}
	for _, ip := range ips {
		nodeIPs = append(nodeIPs, ip.String())
	}
	for _, ip := range ips {
		if name := ipToFQDN(ip.String()); len(name) > 0 {
			return nodeIPs, name
		}
	}
	return nodeIPs, "unknown"
}

// ModuleCPUStateTrace sends message with type ModuleCpuStateTrace
func ModuleCPUStateTrace(nodeAddr string) MessageBuilder {
	nodeIPs, nodeName := addrToIPsFQDN(nodeAddr)
	return &payloadModuleCPUStateTrace{
		Name:         nodeName,
		ConfigNodeIP: nodeIPs,
		BuildInfo: &infoModuleCPUStateTrace{
			BuildVersion:  version.Version,
			BuildTime:     strings.TrimSuffix(version.BuildTime, "+00:00"),
			BuildUser:     version.BuildUser,
			BuildHostname: version.BuildHostname,
			BuildID:       version.BuildID,
			BuildNumber:   version.BuildNumber,
		},
	}
}
