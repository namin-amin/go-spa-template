package customerros

import "slices"

type Errors struct {
	errs []error
}

func (e *Errors) AddNewError(err error)  {
	e.errs = append(e.errs, err)
}

func (e *Errors) DoErrorExists()bool  {
	return slices.ContainsFunc(e.errs,func(er error) bool {
		return er !=nil
	})
}

func NewErrors() *Errors  {
	return &Errors{}
}