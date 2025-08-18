package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type addTxRequest struct {
	RawTx string `json:"rawTx"`
}

type jsonRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type jsonRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	Result  string        `json:"result"`
	Error   *jsonRPCError `json:"error,omitempty"`
	ID      int           `json:"id"`
}

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req addTxRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("json.NewDecoder.Decode(): %v", err), http.StatusBadRequest)
		return
	}

	if req.RawTx == "" || !strings.HasPrefix(req.RawTx, "0x") {
		http.Error(w, "invalid rawTx: must be a non-empty hex string starting with 0x", http.StatusBadRequest)
		return
	}

	rpcURL := os.Getenv("CLIENT_ETH")
	if rpcURL == "" {
		http.Error(w, "CLIENT_ETH not configured", http.StatusInternalServerError)
		return
	}

	payload := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_sendRawTransaction",
		Params:  []interface{}{req.RawTx},
		ID:      1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("json.Marshal(payload): %v", err), http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewReader(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("http.Post(rpcURL): %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("io.ReadAll(resp.Body): %v", err), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("upstream RPC error: status=%d body=%s", resp.StatusCode, string(respBody)), http.StatusBadGateway)
		return
	}

	var rpcResp jsonRPCResponse
	if err := json.Unmarshal(respBody, &rpcResp); err != nil {
		http.Error(w, fmt.Sprintf("json.Unmarshal(respBody): %v", err), http.StatusInternalServerError)
		return
	}

	if rpcResp.Error != nil {
		http.Error(w, fmt.Sprintf("eth_sendRawTransaction error: code=%d message=%s", rpcResp.Error.Code, rpcResp.Error.Message), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"txHash": rpcResp.Result})
}


