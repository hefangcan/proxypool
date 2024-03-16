package provider

import (
	"strings"

	"github.com/bh-qt/proxypool/pkg/tool"
)

type VlessSub struct {
	Base
}

func (sub VlessSub) Provide() string {
	sub.Types = "Vless"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
