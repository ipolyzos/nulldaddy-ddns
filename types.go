package main

// wrap ipify response data
type IPAddr struct {
	IP string `json:"ip"`
}

// wrap godaddy request data
type RecordUpdate struct {
	Data string `json:"data"`
	TTL  int    `json:"ttl"`
}