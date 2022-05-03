package error

type FormatError struct {

}

func (e *FormatError) Error() string {
	return "Format Error"
}
var VideoFormatError  = new(FormatError)