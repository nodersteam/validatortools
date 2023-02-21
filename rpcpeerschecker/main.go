package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("Enter the RPC node address: ")
	var rpcNode string
	fmt.Scanln(&rpcNode)

	peers := retrieveActivePeers(rpcNode)

	fmt.Println("Active peers:")

	comma := ""
	for _, peer := range peers {
		if !strings.Contains(peer.IP, "tcp") {
			if match, _ := regexp.MatchString("^[0-9.:A-Za-Z]+$", peer.IP); match {
				fmt.Print(comma + peer.ID + "@" + peer.IP)
				comma = ","
			}
		}
	}
	fmt.Println()
}

type Peer struct {
	ID string
	IP string
}

func retrieveActivePeers(rpcNode string) []Peer {
	var peers []Peer

	res, err := http.Get(rpcNode + "/net_info")
	if err != nil {
		return peers
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return peers
	}

	var netInfo map[string]interface{}
	err = json.Unmarshal(body, &netInfo)
	if err != nil {
		// Handle error
	}

	result := netInfo["result"].(map[string]interface{})
	peersArray := result["peers"].([]interface{})

	for _, peer := range peersArray {
		peerMap := peer.(map[string]interface{})
		listenAddr := peerMap["node_info"].(map[string]interface{})["listen_addr"].(string)

		ip := listenAddr

		peers = append(peers, Peer{
			ID: peerMap["node_info"].(map[string]interface{})["id"].(string),
			IP: ip,
		})
	}

	return peers
}
