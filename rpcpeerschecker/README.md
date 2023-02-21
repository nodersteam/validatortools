# Active Peers Retriever
A Go program that retrieves the active peers of a given RPC node and displays them in a comma-separated list.

This code is a Go program that retrieves active peer information through an RPC node. It allows the user to input the RPC node address via the terminal. It then calls the retrieveActivePeers function to retrieve active peer information through an HTTP request to the specified RPC node.

The retrieveActivePeers function returns a list of Peer structures containing the identifier and IP address of each peer. After receiving this information, it outputs a list of active peers separated by commas. If the IP address contains the string "tcp", it is not displayed. If the IP address contains invalid characters, it is also not displayed.

## Getting Started
1. Clone this repository to your local machine
`git clone https://github.com/nodersteam/validatortools`
2. Go to the root directory of the repository `cd $HOME/validatortools/rpcpeerschecker` and run the program by typing `go run main.go` in your terminal.
3. Enter the RPC node address when prompted.
![Alt text](https://github.com/nodersteam/picture/blob/main/%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%202023-02-21%20153907.png)

## Prerequisites
- Go installed on your local machine.

## Libraries Used
- fmt - for reading user input and printing results
- net/http - for making HTTP requests
- io/ioutil - for reading the HTTP response body
- regexp - for validating the IP address of a peer
- encoding/json - for unmarshalling the JSON response from the RPC node

## Author
[NODERS]TEAM
