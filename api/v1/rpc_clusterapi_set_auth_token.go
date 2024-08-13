package apiv1

// ClusterAPISetAuthTokenRPC is the path for the ClusterAPISetAuthToken RPC.
const ClusterAPISetAuthTokenRPC = "x/capi/set-auth-token"

// ClusterAPISetAuthTokenRequest is the request message for the ClusterAPISetAuthToken RPC.
type ClusterAPISetAuthTokenRequest struct {
	Token string `json:"token"`
}

// ClusterAPISetAuthTokenResponse is the response message for the ClusterAPISetAuthToken RPC.
type ClusterAPISetAuthTokenResponse struct{}
