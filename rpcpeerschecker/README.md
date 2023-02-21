# Active Peers Retriever
A Go program that retrieves the active peers of a given RPC node and displays them in a comma-separated list.

## Getting Started
1. Clone this repository to your local machine

```git clone https://github.com/nodersteam/validatortools```
2. Go to the root directory of the repository and run the program by typing `go run main.go` in your terminal.
3. Enter the RPC node address when prompted.

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
