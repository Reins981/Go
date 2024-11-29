package iomanager

type IOManager interface { // Define an interface, so that we can use both types (filemanager and cmdmanager) for the NewTaxIncludedPriceJob
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
