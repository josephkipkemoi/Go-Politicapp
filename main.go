package main

import (
	"github.com/josephkipkemoi/politicapp/database"
	"github.com/josephkipkemoi/politicapp/server"
)

func main() {
	const port = ":8080"
	// Connect to Postgresql Database
	database.ConnectDB()

	// Set up server
	p := proxies()
	r := server.SetupRouter()
	r.SetTrustedProxies(p)
	r.Run(port)
}

// Set up server proxies
func proxies() []string {
	var proxies []string
	proxies = append(proxies, "ipv4")

	return proxies
}
