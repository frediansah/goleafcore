package gldb

import (
	"errors"

	"github.com/frediansah/goleafcore/glutil"
)

type Qparams struct {
	counter     int
	params      []string
	paramExists map[string]bool
	values      map[string]interface{}
}

func (p *Qparams) New(key string) string {
	if p.paramExists == nil {
		p.paramExists = map[string]bool{}
	}

	if p.values == nil {
		p.values = map[string]interface{}{}
	}

	if _, paramExists := p.paramExists[key]; !paramExists {
		p.paramExists[key] = true

		p.params = append(p.params, key)
		if p.counter <= 0 {
			p.counter = 0
		}

		p.counter++
	}

	return `$` + glutil.ToString(p.counter) + ` `
}

func (p *Qparams) Set(key string, value interface{}) error {
	if _, exists := p.paramExists[key]; !exists {
		return errors.New("param not found: " + key)
	}

	p.values[key] = value
	return nil
}

func (p Qparams) GetMaps() map[string]interface{} {
	return p.values
}

func (p Qparams) GetValues() []interface{} {
	var params []interface{}
	for _, name := range p.params {
		val := p.values[name]
		params = append(params, val)
	}

	return params
}
