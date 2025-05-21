package main

/*
  The simple idea is to have the Clone() method in the interface which will be implemented by the the class implementing the interface.
  This will make sure the programmer is able to call this easily and this method just creates a new object of the same class with the same
  data and returns it.
*/

import "fmt"

// Cloneable defines the interface for objects that can be cloned
type Cloneable interface {
	Clone() Cloneable
	Print()
}

// Document represents a basic document template
type Document struct {
	Title     string
	Content   string
	Footer    string
	CreatedBy string
}

// Clone creates a copy of the Document
func (d *Document) Clone() Cloneable {
	// Return a new Document with the same properties
	return &Document{
		Title:     d.Title,
		Content:   d.Content,
		Footer:    d.Footer,
		CreatedBy: d.CreatedBy,
	}
}

// Print displays the Document details
func (d *Document) Print() {
	fmt.Printf("Document: %s\nBy: %s\nContent: %s\nFooter: %s\n\n", 
		d.Title, d.CreatedBy, d.Content, d.Footer)
}

func main() {
	// Create an original document template
	originalDoc := &Document{
		Title:     "Report Template",
		Content:   "This is a standard report.",
		Footer:    "Company Confidential",
		CreatedBy: "Template Team",
	}
	
	fmt.Println("Original template:")
	originalDoc.Print()
	
	// Clone the document for a specific report
	salesReportClone := originalDoc.Clone().(*Document)
	salesReportClone.Title = "Q2 Sales Report"
	salesReportClone.Content = "Sales increased by 20% in Q2."
	salesReportClone.CreatedBy = "Sales Department"
	
	fmt.Println("Sales report created from template:")
	salesReportClone.Print()
	
	// Clone again for another department
	marketingReportClone := originalDoc.Clone().(*Document)
	marketingReportClone.Title = "Marketing Strategy 2025"
	marketingReportClone.Content = "New campaign details for next quarter."
	marketingReportClone.CreatedBy = "Marketing Team"
	
	fmt.Println("Marketing report created from template:")
	marketingReportClone.Print()
	
	// The original template remains unchanged
	fmt.Println("Original template (unchanged):")
	originalDoc.Print()
}
