package helpers

import (
	"fmt"
)

func PrintHelp()  {
	Usage()
	Examples()
}

func Usage()  {
	fmt.Println("USAGE:")
	fmt.Println("\tfdl <comands>")
	fmt.Println("COMANDS:")
	fmt.Println("\t-url\t\tURL where file is located")
	fmt.Println("\t-help\t\tHelp for the tool")
	fmt.Println("\t-list\t\tLists all downloaded files")
	fmt.Println("\t-examples\tExamples of usage")
}

func Examples()  {
	fmt.Println("EXAMPLES:")
	fmt.Println("\tfdl -url http://localhost:8080/server/ortiz-cv-v9.pdf\t\tDownload a file from passed url")
	fmt.Println("\tfdl -list\t\t\t\t\t\t\tList all files")
	fmt.Println("\tfdl -examples\t\t\t\t\t\t\tShows examples")
}

func BasicInfo()  {
	fmt.Println("BASIC INFO:")
	fmt.Println("\tSimple CLI File Downloader")
	fmt.Println("\tAutor: Ortiz David")
	fmt.Println("\tEmail: ortizaad1994@gmail.com")
}