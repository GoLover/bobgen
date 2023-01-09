package domain

import "errors"

var (
	ErrNumberOfSlashesMustBeOdd = errors.New(`number of slashes must be odd`)
	ErrCreatingFile             = errors.New(`unexpected error in creating file`)
	ErrRemovingFile             = errors.New(`unexpected error in removing file`)
	ErrFileAlreadyExist         = errors.New(`file already exists`)
	ErrOpeningFile              = errors.New(`unexpected error in reading file`)
	ErrWritingFile              = errors.New(`unexpected error in writing file`)
)
