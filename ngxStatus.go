package ngx_status

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func processParams(p map[string]interface{}, parent string) []string {
	var m []string = []string{}

	for k, v := range p {
		if val, ok := v.(map[string]interface{}); ok {

			par := k
			if parent != "" {
				par = parent + "_" + k
			}

			for _, r := range processParams(val, par) {
				m = append(m, r)
			}

			continue
		}

		if _, ok := v.([]interface{}); ok {
			// does not support arrays
			continue
		}

		par := k
		if parent != "" {
			par = parent + "_" + k
		}

		m = append(m, fmt.Sprintf("%s=%v", par, v))
		//fmt.Printf(" %s = %T [%v]\n", k, v, v)
	}

	//	fmt.Println("--" + strings.Join(m, ",") + "--")
	return m
}

func PrintStatus(f io.Reader) {
	var d map[string]interface{}

	dec := json.NewDecoder(f)
	dec.UseNumber()

	if err := dec.Decode(&d); err != nil {
		panic(err)
	}

	for k, v := range d["serverZones"].(map[string]interface{}) {
		metricsList := processParams(v.(map[string]interface{}), "")

		zone := k
		if zone == "*" {
			zone = "all"
		}

		fmt.Printf("nginx_vts_zone,zone=%s %s\n", zone, strings.Join(metricsList, ","))
	}

	for zname, v := range d["upstreamZones"].(map[string]interface{}) {
		for _, zone := range v.([]interface{}) {
			zn := zone.(map[string]interface{})

			server := zn["server"]
			delete(zn, "server")

			metricsList := processParams(zn, "")

			fmt.Printf("nginx_vts_upstream,zone=%s,server=%s %s\n", zname, server, strings.Join(metricsList, ","))
		}
	}

}

// TRASH
//raw := &NgxStat{}

//dec := json.NewDecoder(f)

//if err := decoder.Decode(raw); err != nil {
//	panic(err)
//}

//var inServersBlock, inMetric bool

//var metricName string

//for {
//	tok, err := dec.Token()
//	if err == io.EOF {
//		break
//	}
//
//	//fmt.Printf("%T: %v\n", tok, tok)
//
//	switch v := tok.(type) {
//	case json.Delim:
//		//fmt.Printf("Delimiter: \n")
//
//	case int, int64, int32, uint, uint32:
//		if inMetric {
//			metrics[metricName] = fmt.Sprintf("%i", v)
//		}
//
//	case string:
//		if inServersBlock {
//			//fmt.Printf("string = %s \n ", v)
//
//			if inMetric {
//				metrics[metricName] = v
//				inMetric = false
//			} else {
//				metricName = v
//				inMetric = true
//			}
//
//		}
//
//		//if inServersBlock && v == "requestCounter" {
//		//	// go further
//		//	fmt.Printf("OLOLO")
//		//}
//
//		if v == "serverZones" {
//			inServersBlock = true
//
//		MAINLOOP:
//			for {
//				tok, err = dec.Token()
//
//				if err == io.EOF {
//					break
//				}
//
//				cnt := 0
//				switch v := tok.(type) {
//				case json.Delim:
//					if v == '{' {
//						cnt++
//
//						if cnt == 2 {
//							break MAINLOOP
//						}
//					}
//				}
//			}
//		}
//	}
//
//	if dec.More() {
//		//		fmt.Printf("(more)\n")
//	}
//
//}

//str := fmt.Sprintf("%s", raw.NginxVersion)
//fmt.Println(str)
//
//for k, v := range metrics {
//	fmt.Printf("%s = %s\n", k, v)
//
//}

//	for _, v := range raw.ServerZones {
//		fmt.Println(v.OverCounts)
//	}

//type JSONStruct interface{}
//
//func parseJSON(dec *json.Decoder) JSONStruct {
//
//	if dec == nil {
//		str := `{"hulio": { "test" : "123",  "many": { "ogo" : "1", "aga" : 2 } }}`
//		dec = json.NewDecoder(strings.NewReader(str))
//	}
//
//	var key string
//
//	res := make(map[string]JSONStruct)
//
//	var pos = 0
//	var arrayBlock bool
//
//	for {
//		tok, err := dec.Token()
//		if err == io.EOF {
//			break
//		}
//
//		//fmt.Printf(" %v = %T \n",  tok, tok )
//		switch v := tok.(type) {
//		case json.Delim:
//			pos = 0
//			switch v {
//			case '{':
//				//fmt.Println("HERE")
//				if key == "" {
//					continue
//				}
//				res[key] = parseJSON(dec)
//			case '[':
//				arrayBlock = true
//			case ']':
//				arrayBlock = false
//			case '}':
//				return res
//			}
//
//			continue
//		}
//
//		if arrayBlock {
//			continue
//		}
//
//		pos++
//		switch v := tok.(type) {
//		case string:
//			if pos&1 == 1 {
//				// key
//				key = v
//			} else {
//				res[key] = v
//			}
//		default:
//			///fmt.Println("bloooo")
//			if pos&1 == 0 {
//				res[key] = v
//			}
//		}
//	}
//
//	return res
//}
//
//func printMap(res JSONStruct) {
//	// print result
//
//	r, ok := res.(map[string]JSONStruct)
//	if !ok {
//		return
//	}
//
//	fmt.Print(" { ")
//	for k, v := range r {
//		if _, ok := v.(map[string]JSONStruct); ok {
//			fmt.Printf("%s = ", k)
//			printMap(v)
//			continue
//		}
//
//		fmt.Printf("%s = %v, ", k, v)
//	}
//
//	fmt.Print("}, ")
//
//}
