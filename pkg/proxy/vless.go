package proxy

import (
	"encoding/json"
	"errors"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorNotVlessink = errors.New("not a correct Vless link")
)

// TODO unknown field
// Link: host, path
// Vless: Network GrpcOpts

type Vless struct {
	Base
	UUID      		string            `yaml:"uuid"`
	UDP        		bool              `yaml:"bool"`
	SNI        		string            `yaml:"sni"`
	Network    		string            `yaml:"type" `
	Flow        	string            `yaml:"flow"`
	Servername   	string            `yaml:"sni"`
	ClientFingerprint	string            `yaml:"fp"`
	TLS          	string            `yaml:"security"`
	Alpn 			string 
	WSOpts             WSOptions
	GrpcOpts		GrpcOptions
}
type WSOptions struct {
	Path    string            `yaml:"path,omitempty" json:"path,omitempty"`
	Headers map[string][]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}
type GrpcOptions struct {
	Method  string              `yaml:"method,omitempty" json:"method,omitempty"`
	Path    []string            `yaml:"path,omitempty" json:"path,omitempty"`
	Headers map[string][]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}
/**
protocol://
	$(uuid)
	@
	remote-host
	:
	remote-port
?
	<protocol-specific fields>
	<transport-specific fields>
	<tls-specific fields>
#$(descriptive-text)
*/

func (t Vless) Identifier() string {
	return net.JoinHostPort(t.Server, strconv.Itoa(t.Port)) + t.UUID
}

func (t Vless) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(data)
}

func (t Vless) ToClash() string {
	data, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return "- " + string(data)
}

func (t Vless) ToSurge() string {
	return ""
}

func (t Vless) Clone() Proxy {
	return &t
}


func (t Vless) Link() (link string) {
	query := url.Values{}
	if t.SNI != "" {
		query.Set("sni", url.QueryEscape(t.SNI))
	}

	uri := url.URL{
		Scheme:   "vless",
		User:     url.User(url.QueryEscape(t.UUID)),
		Host:     net.JoinHostPort(t.Server, strconv.Itoa(t.Port)),
		RawQuery: query.Encode(),
		Fragment: t.Name,
	}

	return uri.String()
}
// https://github.com/nitezs/sub2clash/blob/main/parser/vless.go
func ParseVlessLink(link string) (*Vless, error) {
	if !strings.HasPrefix(link, "vless://") {
		return nil, ErrorNotVlessink
	}
	// 分割
	parts := strings.SplitN(strings.TrimPrefix(link, "vless://"), "@", 2)
	if len(parts) != 2 {
		return nil, ErrorNotVlessink
	}
	
	uri, err := url.Parse(link)
	if err != nil {
		return nil, ErrorNotVlessink
	}

	password := uri.User.Username()
	password, _ = url.QueryUnescape(password)

	server := uri.Hostname()
	port, _ := strconv.Atoi(uri.Port())

	moreInfos := uri.Query()
	sni := moreInfos.Get("sni")
	sni, _ = url.QueryUnescape(sni)
	Network := moreInfos.Get("type")
	Network, _ = url.QueryUnescape(Network)
	TLS := moreInfos.Get("security")
	TLS, _ = url.QueryUnescape(TLS)

	ClientFingerprint := moreInfos.Get("fp")
	ClientFingerprint, _ = url.QueryUnescape(ClientFingerprint)
	Servername := moreInfos.Get("sni")
	Servername, _ = url.QueryUnescape(Servername)
	flow := moreInfos.Get("flow")
	

	wsops := WSOptions{}
	if Network != "" {
			switch Network {
			case "ws":
				//host := make([]string, 0)
				wsops.Path= moreInfos.Get("path")
				//wsops.Headers["host"] = append(host, moreInfos.Get("host"))
			case "tcp":
				flow, _ = url.QueryUnescape(flow)
			}
	}

	return &Vless{
		Base: Base{
			Name:   "",
			Server: server,
			Port:   port,
			Type:   "vless",
			UDP:    true,
		},
		UUID:              password,
		SNI:               sni,
		Network:           Network,
		Flow:              flow,
		Servername: 	 sni,
		ClientFingerprint:  ClientFingerprint,
		TLS:                TLS,
		WSOpts:             wsops,
		//GrpcOpts:           sni,
	}, nil
}

var (
	VlessPlainRe = regexp.MustCompile("vless://([A-Za-z0-9+/_&?=@:%.-])+")
)

func GrepVlessLinkFromString(text string) []string {
	results := make([]string, 0)
	texts := strings.Split(text, "vless://")
	for _, text := range texts {
		results = append(results, VlessPlainRe.FindAllString("vless://"+text, -1)...)
	}
	return results
}
