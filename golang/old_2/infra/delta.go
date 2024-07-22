package main

import (
	"encoding/json"
	"log"
)

type Delta struct {
	ItemId    string          `json:"ItemId"`
	ItemType  int             `json:"ItemType"`
	Delta     json.RawMessage `json:"Delta"`
	EventType int             `json:"EventType"`
}

type AutocheckExploit struct {
	Type       string      `json:"Type"`
	Text       string      `json:"Text"`
	Url        string      `json:"Url"`
	Parameters interface{} `json:"Parameters"`
}

type DeltaAutocheckExploit struct {
	EventType        int              `json:"EventType"`
	AutocheckExploit AutocheckExploit `json:"AutocheckExploit"`
}

type Statistic struct {
	TotalFileCount int    `json:"TotalFileCount"`
	ScanDuration   string `json:"ScanDuration"`
}

func main() {
	log.Printf("Start")

	body := `
{"ItemId":"d414735e-3070-4c2e-8bde-fe869a882b6b","ItemType":1,"Delta":"{\"AutocheckExploit\":{\"Type\":\"HTTP\",\"Text\":\"GET /vulnerabilities/view_source.php?security=319403616&id=%22onmouseover=confirm(53408)%20&3f1319b67f3745599a48c120c31c3c99=3f1319b67f3745599a48c120c31c3c99 HTTP/1.1\\r\\nHost: vulner-dvwa.rd.ptsecurity.ru\\r\\nAccept: */*\\r\\nAccept-Encoding: deflate, gzip\\r\\nCookie: security=low; PHPSESSID=d3f3q13egr5o337i8ucsmh9dd2\\r\\nx-forwarded-for: 127.0.0.1\\r\\nuser-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36\\r\\nreferer: http://127.0.0.1/\\r\\npragma: no-cache\\r\\ncache-control: no-cache\\r\\nx-real-ip: 127.0.0.1\\r\\n\\r\\n\",\"Url\":\"http://vulner-dvwa.rd.ptsecurity.ru/vulnerabilities/view_source.php?security=319403616&id=%22onmouseover=confirm(53408)%20&3f1319b67f3745599a48c120c31c3c99=3f1319b67f3745599a48c120c31c3c99\",\"Parameters\":null},\"ApprovalState\":4}","EventType":0}
`

	log.Printf("Body: %s", body)

	var delta Delta

	if err := json.Unmarshal([]byte(body), &delta); err != nil {
		log.Printf("err: %v", err)
		return
	}

	log.Printf("delta: %s", delta.Delta)

	var Autocheck DeltaAutocheckExploit

	var tmp string
	if err := json.Unmarshal(delta.Delta, &tmp); err != nil {
		log.Fatalf("delta: %v", err)
	}

	if err := json.Unmarshal([]byte(tmp), &Autocheck); err != nil {
		log.Fatalf("err: %v", err)
	}

	log.Printf("stat: %v", Autocheck.AutocheckExploit.Url)
}
