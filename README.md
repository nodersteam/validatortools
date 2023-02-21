RPC Peer List
This program allows you to retrieve a list of active peers from a specified RPC node and display them in the terminal.

Getting Started
To run the program, enter the RPC node address when prompted in the terminal. The program will then retrieve the list of active peers and display them, separated by commas.

Implementation Details
The program uses the http package from the Go standard library to make an HTTP GET request to the specified RPC node and retrieve information on the network. The returned data is in JSON format and is parsed and stored as a Go struct. A regular expression is used to validate the IP addresses of the peers.

Built With
Go programming language
