package input_choose_test

import (
	"bytes"
	// "errors"
	"reflect"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/input_choose"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := input_choose.New(
		[]string{""},
		input_choose.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

// func TestWithValidateFunc(t *testing.T) {
// 	var in bytes.Buffer
// 	var out bytes.Buffer
//
// 	validateErr := errors.New("validation error")
// 	validateFunc := func(s string) error {
// 		if s != "test1" {
// 			return validateErr
// 		}
// 		return nil
// 	}
//
// 	in.Write([]byte("test\r\n"))
//
// 	model := input_choose.New([]string{"test1", "test2", "test3", "test4", "test5"}, input_choose.WithValidateFunc(validateFunc))
//
// 	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
// 	require.Nil(t, err)
//
// 	m, ok := tm.(input_choose.Model)
// 	require.True(t, ok)
//
// 	require.Equal(t, m.Error(), validateErr)
// }

// func TestDefaultValueWithValidateFunc(t *testing.T) {
// 	var in bytes.Buffer
// 	var out bytes.Buffer
//
// 	validateErr := errors.New("validation error")
// 	validateFunc := func(s string) error {
// 		if s != "test1" {
// 			return validateErr
// 		}
// 		return nil
// 	}
//
// 	in.Write([]byte("\r\n"))
//
// 	model := input_choose.New([]string{"test1", "test2"}, input_choose.WithValidateFunc(validateFunc))
//
// 	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
// 	require.Nil(t, err)
//
// 	m, ok := tm.(input_choose.Model)
// 	require.True(t, ok)
//
// 	require.Nil(t, m.Error())
// }

func TestSortedValue(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte("test\r\n"))

	model := input_choose.New([]string{"test1", "test2"})

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input_choose.Model)
	require.True(t, ok)

	require.Equal(t, "test1", m.Data())
	require.Nil(t, m.Error())
}

func TestWithCharLimit(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	inputString := []byte(strings.Repeat("a", 400) + "\r\n")

	in.Write(inputString)

	model := input_choose.New(
		[]string{"test1", "test2"},
		input_choose.WithCharLimit(400),
	)

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input_choose.Model)
	print("HELLLOOO")
	print(m.Data())
	require.True(t, ok)

	require.Equal(t, 400, len(m.Data()))
}

func TestWithSmallCharLimitError(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	inputString := []byte(strings.Repeat("a", 50) + "\r\n")
	in.Write(inputString)

	model := input_choose.New(
		[]string{"test1", "test2"},
		input_choose.WithCharLimit(10),
	)

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input_choose.Model)
	require.True(t, ok)

	require.Equal(t, 10, len(m.Data()))
}
