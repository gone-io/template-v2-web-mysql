package dependent

import "github.com/gone-io/gone"

type iDependent struct {
	gone.Flag
}

func (s *iDependent) DoSomething() error {
	return nil
}
