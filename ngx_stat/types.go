package ngx_stat





type NgxStat struct {
	HostName     string `json:"hostName"`
	NginxVersion string `json:"nginxVersion"`
	LoadMsec     int64  `json:"loadMsec"`
	NowMsec      int64  `json:"nowMsec"`
	Connections  struct {
		Active   int `json:"active"`
		Reading  int `json:"reading"`
		Writing  int `json:"writing"`
		Waiting  int `json:"waiting"`
		Accepted int `json:"accepted"`
		Handled  int `json:"handled"`
		Requests int `json:"requests"`
	} `json:"connections"`
	SharedZones struct {
		Name     string `json:"name"`
		MaxSize  int    `json:"maxSize"`
		UsedSize int    `json:"usedSize"`
		UsedNode int    `json:"usedNode"`
	} `json:"sharedZones"`

	ServerZones []struct {
		RequestCounter int `json:"requestCounter"`
		InBytes        int `json:"inBytes"`
		OutBytes       int `json:"outBytes"`
		Responses      struct {
			OneXx       int `json:"1xx"`
			TwoXx       int `json:"2xx"`
			ThreeXx     int `json:"3xx"`
			FourXx      int `json:"4xx"`
			FiveXx      int `json:"5xx"`
			Miss        int `json:"miss"`
			Bypass      int `json:"bypass"`
			Expired     int `json:"expired"`
			Stale       int `json:"stale"`
			Updating    int `json:"updating"`
			Revalidated int `json:"revalidated"`
			Hit         int `json:"hit"`
			Scarce      int `json:"scarce"`
		} `json:"responses"`
		RequestMsec  int `json:"requestMsec"`
		RequestMsecs struct {
			Times []int64 `json:"times"`
			Msecs []int   `json:"msecs"`
		} `json:"requestMsecs"`

		OverCounts struct {
			MaxIntegerSize uint64 `json:"maxIntegerSize"`
			RequestCounter int    `json:"requestCounter"`
			InBytes        int    `json:"inBytes"`
			OutBytes       int    `json:"outBytes"`
			OneXx          int    `json:"1xx"`
			TwoXx          int    `json:"2xx"`
			ThreeXx        int    `json:"3xx"`
			FourXx         int    `json:"4xx"`
			FiveXx         int    `json:"5xx"`
			Miss           int    `json:"miss"`
			Bypass         int    `json:"bypass"`
			Expired        int    `json:"expired"`
			Stale          int    `json:"stale"`
			Updating       int    `json:"updating"`
			Revalidated    int    `json:"revalidated"`
			Hit            int    `json:"hit"`
			Scarce         int    `json:"scarce"`
		} `json:"overCounts"`
	} `json:"serverZones"`

	UpstreamZones []struct {
		Server         string `json:"server"`
		RequestCounter int    `json:"requestCounter"`
		InBytes        int    `json:"inBytes"`
		OutBytes       int    `json:"outBytes"`
		Responses      struct {
			OneXx   int `json:"1xx"`
			TwoXx   int `json:"2xx"`
			ThreeXx int `json:"3xx"`
			FourXx  int `json:"4xx"`
			FiveXx  int `json:"5xx"`
		} `json:"responses"`
		RequestMsec  int `json:"requestMsec"`
		RequestMsecs struct {
			Times []interface{} `json:"times"`
			Msecs []interface{} `json:"msecs"`
		} `json:"requestMsecs"`
		ResponseMsec  int `json:"responseMsec"`
		ResponseMsecs struct {
			Times []interface{} `json:"times"`
			Msecs []interface{} `json:"msecs"`
		} `json:"responseMsecs"`
		Weight      int  `json:"weight"`
		MaxFails    int  `json:"maxFails"`
		FailTimeout int  `json:"failTimeout"`
		Backup      bool `json:"backup"`
		Down        bool `json:"down"`
		OverCounts  struct {
			MaxIntegerSize uint64 `json:"maxIntegerSize"`
			RequestCounter int    `json:"requestCounter"`
			InBytes        int    `json:"inBytes"`
			OutBytes       int    `json:"outBytes"`
			OneXx          int    `json:"1xx"`
			TwoXx          int    `json:"2xx"`
			ThreeXx        int    `json:"3xx"`
			FourXx         int    `json:"4xx"`
			FiveXx         int    `json:"5xx"`
		} `json:"overCounts"`
	}
}
