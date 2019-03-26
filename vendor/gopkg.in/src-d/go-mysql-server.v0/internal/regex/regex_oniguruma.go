package regex

import (
	rubex "github.com/src-d/go-oniguruma"
)

// Oniguruma holds a rubex regular expression Matcher.
type Oniguruma struct {
	reg *rubex.Regexp
}

// Match implements Matcher interface.
func (r *Oniguruma) Match(s string) bool {
	return r.reg.MatchString2(s)
}

// Dispose implements Disposer interface.
// The function releases resources for oniguruma's precompiled regex
func (r *Oniguruma) Dispose() {
	r.reg.Free2()
}

// NewOniguruma creates a new Matcher using oniguruma engine.
func NewOniguruma(re string) (Matcher, Disposer, error) {
	reg, err := rubex.Compile2(re)
	if err != nil {
		return nil, nil, err
	}

	r := Oniguruma{
		reg: reg,
	}
	return &r, &r, nil
}

func init() {
	err := Register("oniguruma", NewOniguruma)
	if err != nil {
		panic(err.Error())
	}
}
