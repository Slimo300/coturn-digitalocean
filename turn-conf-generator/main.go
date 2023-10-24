package main

import (
	"embed"
	"log"
	"os"
	"strconv"
	"text/template"
)

//go:embed turnserver.conf
var FS embed.FS

func main() {

	publicIP := os.Args[1]
	user := os.Args[2]

	certFile := os.Getenv("TURNS_CERT_FILE")
	keyFile := os.Getenv("TURNS_KEY_FILE")
	realm := os.Getenv("REALM")
	domain := os.Getenv("TURN_DOMAIN")
	tlsPort := os.Getenv("TLS_PORT")

	if len(publicIP) == 0 {
		log.Fatal("public-ip flag must be specified")
	}
	if len(user) == 0 {
		log.Fatal("user flag must be specified")
	}

	if len(certFile) == 0 {
		log.Fatal("TURNS_CERT_FILE not specified")
	}
	if len(keyFile) == 0 {
		log.Fatal("TURNS_KEY_FILE not specified")
	}
	if len(realm) == 0 {
		log.Fatal("REALM not specified")
	}
	if len(domain) == 0 {
		log.Fatal("TURN_DOMAIN not specified")
	}

	_, err := strconv.Atoi(tlsPort)
	if err != nil {
		log.Fatalf("Error converting TLS_PORT to int: %v", err)
	}

	tmpl, err := template.ParseFS(FS, "turnserver.conf")
	if err != nil {
		log.Fatalf("Error creating a template: %v", err)
	}

	if err := tmpl.Execute(os.Stdout, map[string]string{
		"IP_ADDRESS":  publicIP,
		"CREDENTIALS": user,
		"CERT_FILE":   certFile,
		"KEY_FILE":    keyFile,
		"REALM":       realm,
		"SERVER_NAME": domain,
		"TLS_PORT":    tlsPort,
	}); err != nil {
		log.Fatalf("Error executing a template")
	}
}
