package main

import (
	"flag"     // for parsing command-line options
	"log"      // for logging server events
	"net/http" // to use HTTP server functionaltities
)

func main() {
	/* Define command-line flags for specifying the directory of static files and the server address.
	These flags allow for flexible server configuration. */

	// var staticDir = flag.String("static-dir", "../chatapp-frontend/html", "The directory of static files to host")
	var staticDir = flag.String("static-dir", "./html", "The directory of static files to host")

	// var addr = flag.String("addr", "0.0.0.0:8081", "The address to listen on for HTTP requests.")
	var addr = flag.String("addr", "0.0.0.0:30222", "The address to listen on for HTTP requests.")

	flag.Parse() // Parse the provided command-line options.

	/* Create a file server handler that serves static files (i.e., my HTML, CSS, JavaScript)
	from the specified directory. This enables the chat application's web user interface,
	fulfilling Part 1's requirement for a web UI where users can read and send messages, and
	having no need for nginx for part 2.
	*/

	fs := http.FileServer(http.Dir(*staticDir))

	/* Register the file server handler to respond to HTTP requests on all routes ("/").
	This *effectively* makes this Go program work/pretend to be a web server, directly addressing
	Part 2's requirement to build my own web server with Go, therefore avoiding external
	solutions like Nginx for serving static content.
	*/

	http.Handle("/", fs)

	/* Log the server's start, indicating that the static file server is operational
	and listening on the specified address. This server exclusively handles the
	delivery of the frontend assets to clients' browsers.
	*/

	log.Printf("Static file server started on %s\n", *addr)

	/*  Start the HTTP server on the specified address. This call blocks, keeping the
	server running to serve incoming HTTP requests. If the server fails to start,
	log the error and terminate the program. This component is crucial for ensuring
	the chat application is accessible to users, contributing to the distributed
	nature of the application by potentially running on separate infrastructure.
	*/
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err) // Log fatal errors that prevent server startup.
	}
}
