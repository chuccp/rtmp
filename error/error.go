package error

import "errors"

type FormatError struct {

}

func (e *FormatError) Error() string {
	return "Format Error"
}
var VideoFormatError  = new(FormatError)

var UnknownFormatError = errors.New("UnknownFormatError")