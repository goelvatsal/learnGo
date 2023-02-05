package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
}

func fieldParse(fields []string, p parser) bool {
	if len(fields) != 2 {
		fmt.Printf("wrong input: %v (line #%d)\n", fields, p.lines)
		return false
	}
	return false
}

func printTotal(p parser) bool {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	return false
}

func handleError(in *bufio.Scanner) bool {
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
	return false
}

func printTotalVisits(p parser) bool {
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
	return false

}

func newVisitCopy(p parser, domain string, visits int) bool {
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
	return false
}
