package proxy

import (
	"fmt"
	"testing"

	"github.com/bh-qt/proxypool/pkg/tool"
	"github.com/ghodss/yaml"
)

func TestSSLink(t *testing.T) {
	ss, err := ParseSSLink("ss://YWVzLTI1Ni1jZmI6ZUlXMERuazY5NDU0ZTZuU3d1c3B2OURtUzIwMXRRMERAMTcyLjEwNC4xNjEuNTQ6ODA5OQ==#翻墙党223.13新加坡")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ss)
	fmt.Println(ss.Link())
	ss, err = ParseSSLink(ss.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ss)
}

func TestSSRLink(t *testing.T) {
	ssr, err := ParseSSRLink("ssr://MTcyLjEwNC4xNjEuNTQ6ODA5OTpvcmlnaW46YWVzLTI1Ni1jZmI6cGxhaW46WlVsWE1FUnVhelk1TkRVMFpUWnVVM2QxYzNCMk9VUnRVekl3TVhSUk1FUT0vP29iZnNwYXJhbT0mcHJvdG9wYXJhbT0mcmVtYXJrcz01Ny03NWFLWjVZV2FNakl6TGpFejVwYXc1WXFnNVoyaCZncm91cD01cGF3NVlxZzVaMmg=")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ssr)
	fmt.Println(ssr.Link())
	ssr, err = ParseSSRLink(ssr.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ssr)
	fmt.Println(ssr.ToClash())
}

func TestTrojanLink(t *testing.T) {
	trojan, err := ParseTrojanLink("trojan://65474277@sqcu.hostmsu.ru:55551?allowinsecure=0&peer=mza.hkfq.xyz&mux=1&ws=0&wspath=&wshost=&ss=0&ssmethod=aes-128-gcm&sspasswd=&group=#%E9%A6%99%E6%B8%AFCN2-MZA%E8%8A%82%E7%82%B9-%E5%AE%BF%E8%BF%81%E8%81%94%E9%80%9A%E4%B8%AD%E8%BD%AC")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(trojan)
	fmt.Println(trojan.Link())
	trojan, err = ParseTrojanLink(trojan.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(trojan)
}

func TestVmessLink(t *testing.T) {
	//v, err := ParseVmessLink("vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIuW+ruS/oeWFrOS8l+WPtyDlpJrlvannmoTlpKfljYPkuJbnlYwiLA0KICAiYWRkIjogInMyNzEuc25vZGUueHl6IiwNCiAgInBvcnQiOiAiNDQzIiwNCiAgImlkIjogIjZhOTAwZDYzLWNiOTItMzVhMC1hZWYwLTNhMGMxMWFhODUyMyIsDQogICJhaWQiOiAiMSIsDQogICJuZXQiOiAid3MiLA0KICAidHlwZSI6ICJub25lIiwNCiAgImhvc3QiOiAiczI3MS5zbm9kZS54eXoiLA0KICAicGF0aCI6ICIvcGFuZWwiLA0KICAidGxzIjogInRscyINCn0=")
	//v, err := ParseVmessLink("vmess://YXV0bzphMjA1ZjRiNi0xMzg2LTQ3NjUtYjQ0YS02YjFiYmE0N2Q1MzdAMTQyLjQuMTA0LjIyNjo0NDM?remarks=%F0%9F%87%BA%F0%9F%87%B8%20US_616%20caicai&obfsParam=www.036452916.xyz&path=/footers&obfs=websocket&tls=1&allowInsecure=1&alterId=64")
	v, err := ParseVmessLink("vmess://YXV0bzo1YjQ1ZjQ2Yi1iNTVmLTRkNWQtOGJjOS1jZjY1MzZlZjkyMzhAMTM3LjE3NS4zNS4xMzo0NDM?remarks=%F0%9F%87%BA%F0%9F%87%B8%20US_480%20caicai&obfsParam=www.4336705.xyz&path=/footers&obfs=websocket&tls=1&allowInsecure=1&alterId=64")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	fmt.Println(v.Link())
	v, err = ParseVmessLink(v.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestNewVmessParser(t *testing.T) {
	linkPayload := "ew0KICAidiI6ICIyIiwNCiAgInBzIjogIuW+ruS/oeWFrOS8l+WPtyDlpJrlvannmoTlpKfljYPkuJbnlYwiLA0KICAiYWRkIjogInMyNzEuc25vZGUueHl6IiwNCiAgInBvcnQiOiAiNDQzIiwNCiAgImlkIjogIjZhOTAwZDYzLWNiOTItMzVhMC1hZWYwLTNhMGMxMWFhODUyMyIsDQogICJhaWQiOiAiMSIsDQogICJuZXQiOiAid3MiLA0KICAidHlwZSI6ICJub25lIiwNCiAgImhvc3QiOiAiczI3MS5zbm9kZS54eXoiLA0KICAicGF0aCI6ICIvcGFuZWwiLA0KICAidGxzIjogInRscyINCn0="
	payload, err := tool.Base64DecodeString(linkPayload)
	if err != nil {
		fmt.Println("vmess link payload parse failed")
		return
	}
	jsonMap, err := str2jsonDynaUnmarshal(payload)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	vmessJson, err := mapStrInter2VmessLinkJson(jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(vmessJson)
}

func TestSSRClashYaml(t *testing.T) {
	str := "{\"name\":\"JP_609\",\"server\":\"13.231.143.248\",\"ip\":\"13.231.143.248\",\"outip\":\"\",\"port\":857,\"type\":\"ssr\",\"country\":\"JP\",\"flag\":\"🇯🇵\",\"usable\":true,\"delay\":847,\"Download\":0,\"Upload\":0,\"password\":\"CF5IKQ\",\"cipher\":\"chacha20-ietf\",\"protocol\":\"auth_aes128_sha1\",\"protocol-param\":\"45063:tyaGuO\",\"obfs\":\"tls1.2_ticket_auth\",\"obfs-param\":\"ffb1945063.microsoft.com\",\"group\":\"proxycrawler-clash\"}"
	var ssr ShadowsocksR
	err := yaml.Unmarshal([]byte(str), &ssr)
	if err != nil {
		panic(err)
	}
	fmt.Println(ssr)
	fmt.Println(ssr.Link())
	fmt.Println(ssr.ToClash())
	ssrp, err := ParseSSRLink(ssr.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ssrp)
	fmt.Println(ssrp.ToClash())

}
func main() {
	//v, err := ParseVmessLink("vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIuW+ruS/oeWFrOS8l+WPtyDlpJrlvannmoTlpKfljYPkuJbnlYwiLA0KICAiYWRkIjogInMyNzEuc25vZGUueHl6IiwNCiAgInBvcnQiOiAiNDQzIiwNCiAgImlkIjogIjZhOTAwZDYzLWNiOTItMzVhMC1hZWYwLTNhMGMxMWFhODUyMyIsDQogICJhaWQiOiAiMSIsDQogICJuZXQiOiAid3MiLA0KICAidHlwZSI6ICJub25lIiwNCiAgImhvc3QiOiAiczI3MS5zbm9kZS54eXoiLA0KICAicGF0aCI6ICIvcGFuZWwiLA0KICAidGxzIjogInRscyINCn0=")
	//v, err := ParseVmessLink("vmess://YXV0bzphMjA1ZjRiNi0xMzg2LTQ3NjUtYjQ0YS02YjFiYmE0N2Q1MzdAMTQyLjQuMTA0LjIyNjo0NDM?remarks=%F0%9F%87%BA%F0%9F%87%B8%20US_616%20caicai&obfsParam=www.036452916.xyz&path=/footers&obfs=websocket&tls=1&allowInsecure=1&alterId=64")
	v, err := ParseVlessLink("vless://e4e7b149-8bd8-4266-8cbb-0987604dbc87@104.16.222.204:443?encryption=none&security=tls&sni=toujiediansigema.5112233.xyz&type=ws&host=toujiediansigema.5112233.xyz&path=tg%40juzibaipiao%E5%81%B7%E8%8A%82%E7%82%B9%E6%AD%BB%E5%90%97#%F0%9F%87%BA%F0%9F%87%B8.US.%E7%BE%8E%E5%9B%BD.3")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
	fmt.Println(v.Link())
	v, err = ParseVlessLink(v.Link())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
