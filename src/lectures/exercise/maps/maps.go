//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status displaying:
//   - number of servers
//   - number of servers for each status (Online, Offline, Maintenance, Retired)
func printServerInfo(serverMap map[string]int) {
	var serverCounter = make(map[int]int)
	for _, status := range serverMap {
		serverCounter[status]++
	}

	var serverStatusTitles = []string{"Online", "Offline", "Maintenance", "Retired"}

	fmt.Println("total servers:", len(serverMap))
	fmt.Println("server statuses:")
	for status, serverCount := range serverCounter {
		fmt.Println(serverCount, "servers in state", serverStatusTitles[status])
	}
}

func main() {
	var servers = []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map
	var serverMap = make(map[string]int)
	for _, serverName := range servers {
		serverMap[serverName] = Online
	}

	//  - call display server info function
	printServerInfo(serverMap)

	//  - change server status of `darkstar` to `Retired`
	serverMap["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverMap["aiur"] = Offline
	//  - call display server info function
	printServerInfo(serverMap)

	//  - change server status of all servers to `Maintenance`
	for _, serverName := range servers {
		serverMap[serverName] = Maintenance
	}
	//  - call display server info function
	printServerInfo(serverMap)
}
