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
	UUID      		string            `yaml:"uuid" json:"uuid"`
	SNI        		string            `yaml:"sni" json:"sni"`
	Network    		string            `yaml:"network,omitempty" json:"network,omitempty"`
	Flow        	string            	  `yaml:"flow,omitempty" json:"flow,omitempty"`
	Servername   	string            	  `yaml:"servername,omitempty" json:"servername,omitempty"`
	ClientFingerprint	string            `yaml:"client-fingerprint,omitempty" json:"client-fingerprint,omitempty"`
	TLS            bool              	  `yaml:"tls,omitempty" json:"tls,omitempty"`
	Tfo 			bool 		  `yaml:"tfo" json:"tfo"`
	SkipCertVerify bool              	  `yaml:"skip-cert-verify,omitempty" json:"skip-cert-verify,omitempty"`
	Wsopts             WSOptions              `yaml:"ws-opts,omitempty" json:"ws-opts,omitempty"`
	GrpcOpts	   GrpcOptions		  `yaml:"grpc-opts,omitempty" json:"grpc-opts,omitempty"`
	RealOpts	   RealOptions		  `yaml:"reality-opts,omitempty" json:"reality-opts,omitempty"`
}
type WSOptions struct {
	Path    string            `yaml:"Path,omitempty" json:"Path,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`
}
type GrpcOptions struct {
	Grpcname  string              `yaml:"grpc-service-name,omitempty" json:"grpc-service-name,omitempty"`
}
type RealOptions struct {
	Pbk    string            `yaml:"public-key,omitempty" json:"public-key,omitempty"`
	Sid    string            `yaml:"short-id,omitempty" json:"short-id,omitempty"`
}
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
	tls := moreInfos.Get("security")
	tls, _ = url.QueryUnescape(tls)
	tlss := false
	if tls == "tls" {
		tlss = true
	}
	ClientFingerprint := moreInfos.Get("fp")
	ClientFingerprint, _ = url.QueryUnescape(ClientFingerprint)
	Servername := moreInfos.Get("sni")
	Servername, _ = url.QueryUnescape(Servername)
	flow := moreInfos.Get("flow")
	wsops := WSOptions{}
	grpcops := GrpcOptions{}
	realops := RealOptions{}
	if tls == "reality" {
		realops.Pbk= moreInfos.Get("pbk")
		realops.Sid= moreInfos.Get("sid")
	}
	if Network != "" {
			switch Network {
			case "ws":
				
				wsHeaders := make(map[string]string)
				host := moreInfos.Get("host")
				wsHeaders["Host"] = host
				wsops.Headers = wsHeaders
				wsops.Path= moreInfos.Get("path")
				if wsops.Path == ""{
					wsops.Path = "/"
				}
			case "tcp":
				flow, _ = url.QueryUnescape(flow)
			case "grpc":
				grpcNmae := moreInfos.Get("serviceNmae")
				grpcops.Grpcname, _ = url.QueryUnescape(grpcNmae)
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
		Servername: 	   sni,
		ClientFingerprint:  ClientFingerprint,
		TLS:                tlss,
		Wsopts:             wsops,
		GrpcOpts:           grpcops,
		RealOpts:           realops,
		SkipCertVerify:    false,
		Tfo:    false,
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
