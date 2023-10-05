package main

// Email Domain Checker tool
// written by Hamza <cyb3rguru>
// Copyright (c) 2023 
// License for use in source and binary forms
// is licensed under the MIT License
// http://opens license.org/licenses/mit-license.php
// =================================================================


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
	fmt.Println("------------Welcome To Email Domain Checker-------------------\n")
	fmt.Printf("Enter Email Domain: ")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: This happened %v\n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Println("------------------ Checker Result ---------------------\n")
	fmt.Printf("Domain: %v \nHasMx: %v \nhasSPF: %v \nspfRecord: %v \nhasDMARC: %v \ndmarcRecord: %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Printf("\nEnter Email Domain: ")
}
