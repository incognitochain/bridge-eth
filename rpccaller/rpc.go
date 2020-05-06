package rpccaller

type RPCError struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

type RPCBaseRes struct {
	Id       int       `json:"Id"`
	RPCError *RPCError `json:"Error"`
}
