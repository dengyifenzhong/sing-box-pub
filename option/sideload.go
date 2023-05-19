package option

type SideLoadOutboundOptions struct {
	DialerOptions
	ServerOptions
	ListenPort      uint16           `json:"listen_port"`
	ListenNetwork   NetworkList      `json:"listen_network,omitempty"`
	UDPTimeout      int64            `json:"udp_timeout,omitempty"`
	Command         Listable[string] `json:"command"`
	Env             Listable[string] `json:"env,omitempty"`
	Socks5ProxyPort uint16           `json:"socks5_proxy_port"`
	Network         NetworkList      `json:"network,omitempty"`
}
