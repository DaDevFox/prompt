package input_choose

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Choice struct {
	Text string
}

type Model struct {
	textInput textinput.Model
	choices   []Choice
	filtered  []Choice
	cursor    int

	validateFunc ValidateFunc
	inputMode    InputMode

	quitting       bool
	err            error
	keys           KeyMap
	showHelp       bool
	help           help.Model
	teaProgramOpts []tea.ProgramOption
}

func New(choices []string, opts ...Option) *Model {
	ti := textinput.New()
	ti.Placeholder = "Start typing..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Prompt = ""

	c := make([]Choice, len(choices))
	for i, choice := range choices {
		c[i] = Choice{Text: choice}
	}

	m := &Model{
		textInput: ti,
		choices:   c,
		filtered:  c,
		cursor:    0,

		quitting:       false,
		err:            nil,
		keys:           DefaultKeyMap,
		showHelp:       false,
		help:           help.New(),
		teaProgramOpts: make([]tea.ProgramOption, 0),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *Model) FilterChoices() {
	query := strings.ToLower(m.textInput.Value())
	var filtered []Choice
	for _, choice := range m.choices {
		if strings.Trim(query, " ") == "" || strings.Contains(strings.ToLower(choice.Text), query) {
			filtered = append(filtered, choice)
		}
	}
	m.filtered = filtered
	if m.cursor >= len(m.filtered) {
		m.cursor = len(m.filtered) - 1
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Prev):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Next):
			if m.cursor < len(m.filtered)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Confirm):
			if m.err == nil && m.validateFunc != nil {
				currVal := m.textInput.Value()
				if currVal == "" {
					currVal = m.textInput.Placeholder
				}
				m.err = m.validateFunc(currVal)
			}
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			if m.showHelp {
				m.help.ShowAll = !m.help.ShowAll
			}
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	m.FilterChoices()

	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return "Selected: " + m.filtered[m.cursor].Text + "\n"
	}

	view := m.textInput.View() + "\n"
	for i, choice := range m.filtered {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		view += cursor + " " + choice.Text + "\n"
	}
	return view
}

func (m Model) Data() string {
	return m.choices[m.cursor].Text
}

func (m Model) DataString() string {
	return m.Data()
}

func (m Model) Quitting() bool {
	return m.quitting
}

func (m Model) Error() error {
	return m.err
}

func (m *Model) WithInputMode(mode InputMode) *Model {
	m.inputMode = mode
	return m
}

func (m *Model) WithEchoMode(mode EchoMode) *Model {
	m.textInput.EchoMode = mode
	return m
}

func (m *Model) WithValidateFunc(vf ValidateFunc) *Model {
	m.validateFunc = vf
	return m
}

func (m Model) TeaProgramOpts() []tea.ProgramOption {
	return m.teaProgramOpts
}

// package input
//
// import (
// 	"strings"
// 	"unicode"
//
// 	"github.com/charmbracelet/bubbles/help"
// 	"github.com/charmbracelet/bubbles/key"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/cqroot/prompt/constants"
// )
//
// type Choice struct {
// 	Text string
// 	Note string
// }
//
// type Model struct {
// 	choices      []Choice
// 	df           string
// 	textInput    textinput.Model
// 	validateFunc ValidateFunc
// 	inputMode    InputMode
//
// 	quitting       bool
// 	err            error
// 	keys           KeyMap
// 	showHelp       bool
// 	help           help.Model
// 	teaProgramOpts []tea.ProgramOption
// }
//
// func NewWithStrings(defaultValue string, choices []string, opts ...Option) *Model {
// 	_choices := make([]Choice, 0, len(choices))
// 	for _, choice := range choices {
// 		_choices = append(_choices, Choice{Text: choice})
// 	}
// 	return New(defaultValue, _choices, opts...)
// }
//
// func New(defaultValue string, choices []Choice, opts ...Option) *Model {
// 	ti := textinput.New()
// 	ti.Placeholder = defaultValue
// 	ti.Focus()
// 	ti.CharLimit = 156
// 	ti.Width = 40
// 	ti.Prompt = ""
//
// 	m := &Model{
// 		textInput:      ti,
// 		df:             defaultValue,
// 		inputMode:      InputAll,
// 		quitting:       false,
// 		err:            nil,
// 		keys:           DefaultKeyMap,
// 		showHelp:       false,
// 		help:           help.New(),
// 		teaProgramOpts: make([]tea.ProgramOption, 0),
// 	}
//
// 	for _, opt := range opts {
// 		opt(m)
// 	}
//
// 	return m
// }
//
// func (m Model) Data() string {
// 	if m.textInput.Value() == "" {
// 		return m.textInput.Placeholder
// 	} else {
// 		return m.textInput.Value()
// 	}
// }
//
// func (m Model) DataString() string {
// 	if m.textInput.EchoMode == EchoNormal {
// 		return m.Data()
// 	}
// 	m.textInput.Blur()
// 	str := m.textInput.View()
// 	m.textInput.Focus()
// 	return str
// }
//
// func (m Model) Quitting() bool {
// 	return m.quitting
// }
//
// func (m Model) Error() error {
// 	return m.err
// }
//
// func (m Model) TeaProgramOpts() []tea.ProgramOption {
// 	return m.teaProgramOpts
// }
//
// func (m *Model) WithInputMode(mode InputMode) *Model {
// 	m.inputMode = mode
// 	return m
// }
//
// func (m *Model) WithEchoMode(mode EchoMode) *Model {
// 	m.textInput.EchoMode = mode
// 	return m
// }
//
// func (m *Model) WithValidateFunc(vf ValidateFunc) *Model {
// 	m.validateFunc = vf
// 	return m
// }
//
// func (m Model) Init() tea.Cmd {
// 	return textinput.Blink
// }
//
// func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.help.Width = msg.Width
//
// 	case tea.KeyMsg:
// 		switch {
// 		case key.Matches(msg, m.keys.Confirm):
// 			if m.err == nil && m.validateFunc != nil {
// 				currVal := m.textInput.Value()
// 				if currVal == "" {
// 					currVal = m.textInput.Placeholder
// 				}
// 				m.err = m.validateFunc(currVal)
// 			}
// 			m.quitting = true
// 			return m, tea.Quit
//
// 		case key.Matches(msg, m.keys.Quit):
// 			m.quitting = true
// 			m.err = constants.ErrUserQuit
// 			return m, tea.Quit
// 		}
//
// 		if m.inputMode == InputNumber || m.inputMode == InputInteger {
// 			keypress := msg.String()
// 			if len(keypress) == 1 {
// 				if keypress == "." {
// 					if m.inputMode != InputNumber ||
// 						strings.Contains(m.textInput.Value(), ".") {
// 						return m, nil
// 					}
// 				} else {
// 					if !unicode.IsNumber([]rune(keypress)[0]) {
// 						return m, nil
// 					}
// 				}
// 			}
// 		}
// 	}
//
// 	var cmd tea.Cmd
// 	m.textInput, cmd = m.textInput.Update(msg)
// 	return m, cmd
// }
//
// func (m Model) View() string {
// 	view := m.textInput.View()
//
// 	if m.textInput.Value() != "" && m.validateFunc != nil {
// 		err := m.validateFunc(m.textInput.Value())
// 		if err != nil {
// 			view = view + constants.DefaultErrorPromptPrefixStyle.Render("\n✖  ") +
// 				constants.DefaultNoteStyle.Render(err.Error())
// 		}
// 	}
//
// 	if m.showHelp {
// 		view += "\n\n"
// 		view += m.help.View(m.keys)
// 	}
//
// 	return view
// }
