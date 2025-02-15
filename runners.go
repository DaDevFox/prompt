package prompt

import (
	"github.com/DaDevFox/prompt/choose"
	"github.com/DaDevFox/prompt/input"
	"github.com/DaDevFox/prompt/input_choose"
	"github.com/DaDevFox/prompt/multichoose"
	"github.com/DaDevFox/prompt/write"
)

// Choose lets the user choose one of the given choices.
func (p Prompt) Choose(choices []string, opts ...choose.Option) (string, error) {
	pm := choose.NewWithStrings(choices, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return "", err
	}
	return m.(choose.Model).Data(), nil
}

// Choose lets the user choose one of the given choices.
func (p Prompt) AdvancedChoose(choices []choose.Choice, opts ...choose.Option) (string, error) {
	pm := choose.New(choices, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return "", err
	}
	return m.(choose.Model).Data(), nil
}

// MultiChoose lets the user choose multiples from the given choices.
func (p Prompt) MultiChoose(choices []string, opts ...multichoose.Option) ([]string, error) {
	pm := multichoose.New(choices, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return nil, err
	}
	return m.(multichoose.Model).Data(), nil
}

// Input asks the user to enter a string.
func (p Prompt) InputChoose(choices []string, opts ...input_choose.Option) (string, error) {
	pm := input_choose.New(choices, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return "", err
	}
	return m.(input.Model).Data(), nil
}

// Input asks the user to enter a string.
func (p Prompt) Input(defaultValue string, opts ...input.Option) (string, error) {
	pm := input.New(defaultValue, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return "", err
	}
	return m.(input.Model).Data(), nil
}

func (p Prompt) Write(defaultValue string, opts ...write.Option) (string, error) {
	pm := write.New(defaultValue, opts...)

	m, err := p.Run(*pm, append(p.teaProgramOpts, pm.TeaProgramOpts()...)...)
	if err != nil {
		return "", err
	}
	return m.(write.Model).Data(), nil
}
