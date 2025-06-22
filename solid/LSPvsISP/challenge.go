/*
🧪 Challenge: Real-World Print System
You're building a system for different types of office machines.

Requirements:
You have multiple types of machines:

BasicPrinter – Can only print

SmartPrinter – Can print, scan, and fax

PhotoPrinter – Can print and scan, but not fax

You need to:

Design a set of interfaces and structs.

Follow Interface Segregation Principle so no struct is forced to implement extra methods.

Follow Liskov Substitution Principle — if a function expects a printer/scanner/fax, it should not panic or behave incorrectly.

🎯 Expected Usage Example (you define how it works):

basic := BasicPrinter{}
smart := SmartPrinter{}
photo := PhotoPrinter{}

// These functions should work without crashing
PrintSomething(basic)
ScanSomething(photo)
FaxSomething(smart)
*/

package main

import "fmt"

type Printer interface {
	Print(document string)
}

type Scanner interface {
	Scan(document string)
}

type Faxer interface {
	Fax(document string)
}

type BasicPrinter struct{}
type SmartPrinter struct{}
type PhotoPrinter struct{}

var _ Printer = (*BasicPrinter)(nil)
var _ Printer = (*SmartPrinter)(nil)
var _ Scanner = (*SmartPrinter)(nil)
var _ Faxer = (*SmartPrinter)(nil)
var _ Printer = (*PhotoPrinter)(nil)
var _ Scanner = (*PhotoPrinter)(nil)

func (bp BasicPrinter) String() string {
	return "BasicPrinter"
}

func (sp SmartPrinter) String() string {
	return "SmartPrinter"
}

func (pp PhotoPrinter) String() string {
	return "PhotoPrinter"
}

func (bp BasicPrinter) Print(document string) {
	fmt.Printf("[%s] printing document: %s", bp, document)
}

func (sp SmartPrinter) Print(document string) {
	fmt.Printf("[%s] printing document: %s", sp, document)
}

func (sp SmartPrinter) Scan(document string) {
	fmt.Printf("[%s] Scanning document: %s", sp, document)
}

func (sp SmartPrinter) Fax(document string) {
	fmt.Printf("[%s] Faxing document: %s", sp, document)
}

func (pp PhotoPrinter) Print(document string) {
	fmt.Printf("[%s] printing document: %s", pp, document)
}

func (pp PhotoPrinter) Scan(document string) {
	fmt.Printf("[%s] Scanning document: %s", pp, document)
}

func PrintSomething(printer Printer, document string) {
	printer.Print(document)
}

func ScanSomething(scanner Scanner, document string) {
	scanner.Scan(document)
}

func FaxSomething(faxer Faxer, document string) {
	faxer.Fax(document)
}

func main() {
	basic := BasicPrinter{}
	smart := SmartPrinter{}
	photo := PhotoPrinter{}

	document := "confidential.docx"

	PrintSomething(basic, document)
	ScanSomething(photo, document)
	FaxSomething(smart, document)
}
