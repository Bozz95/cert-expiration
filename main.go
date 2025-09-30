package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/url"
	"time"
)

func checkCertExpiration(u string) (int, error) {
	fmt.Println("Checking Certificate Expiration")

	if u == "" {
		return 0, fmt.Errorf("URL is required")
	}
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return 0, fmt.Errorf("invalid URL: %s", err)
	}

	if parsedURL.Scheme != "https" {
		return 0, fmt.Errorf("invalid URL scheme (only https supported): %s", parsedURL.Scheme)
	}

	p := parsedURL.Port()
	addr := parsedURL.Host
	if p == "" {
		fmt.Println("Port is empty, using the default 443")
		addr = addr + ":443"
	}

	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to connect: %s", err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
	if daysLeft < 0 {
		return 0, fmt.Errorf("certificate has expired")
	}
	return daysLeft, nil
}

func main() {
	const version = "v0.0.1"

	url := flag.String("url", "", "Https URL which certification expiration will be checked. Required")
	versionFlag := flag.Bool("version", false, "Print version and exit")

	flag.Parse()

	if *versionFlag {
		fmt.Println("cert-expiration version:", version)
		return
	}

	if *url == "" {
		fmt.Println("URL is required")
		flag.PrintDefaults()
		return
	}

	exDays, err := checkCertExpiration(*url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Checking Certificate Expiration of: %s\n", *url)
	fmt.Printf("Expiration in %d days\n", exDays)
}
