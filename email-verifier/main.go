package main

// imports
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a domain name: ")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error: %v", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	// check if domain is valid
	mxRecords, err := net.LookupMX(domain)
	// check if domain has MX records
	if err != nil {
		log.Printf("error: %v", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	// check SPF record
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, record := range txtRecords {
		// check if domain has SPF record
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	// check DMARC record
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, record := range dmarcRecords {
		// check if domain has DMARC record
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	// print results
	fmt.Printf("%v, %v, %v, %v, %v,%v \n", domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecord)
}
