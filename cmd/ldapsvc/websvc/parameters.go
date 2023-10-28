package websvc

import "fmt"

func (p *Parameters) Validate() error {
	if p == nil {
		return fmt.Errorf("empty parameters")
	}

	if p.Port == 0 {
		return fmt.Errorf("invalid port %d", p.Port)
	}
	if err := p.Certificate.Validate(); err != nil {
		return fmt.Errorf("invalid certificate configuration (%v)", err.Error())
	}
	return nil
}
