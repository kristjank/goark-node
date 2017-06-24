package api

//PeerResponse structure for call /peer/list
type PeerResponse struct {
	Success    bool   `json:"success"`
	Peers      []Peer `json:"peers"`
	SinglePeer Peer   `json:"peer"`
}

//Peer structure to hold peer data
type Peer struct {
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Version string `json:"version,omitempty"`
	Os      string `json:"os,omitempty"`
	Height  int    `json:"height,omitempty"`
	Status  string `json:"status"`
	Delay   int    `json:"delay"`
}

// PeerQueryParams - when set, they are automatically added to get requests
type PeerQueryParams struct {
	Status  string `url:"status,omitempty"`  //State of peer. OK, ETIMEOUT,...
	Os      string `url:"os,omitempty"`      //OS of peer. (String)
	Shared  string `url:"shared,omitempty"`  //Is peer shared? Boolean: true or false. (String)
	Version string `url:"version,omitempty"` //Version of peer. (String)
	Limit   int    `url:"limit,omitempty"`   //Limit to show. Max limit is 100. (Integer)
	OrderBy string `url:"orderBy,omitempty"` //Name of column to order. After column name must go 'desc' or 'asc' to choose order type. (String)
	Offset  int    `url:"offset,omitempty"`
	IP      string `url:"ip,omitempty"`
	Port    int    `url:"port,omitempty"`
}
