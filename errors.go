package prompt

import (
	"errors"

	"github.com/DaDevFox/prompt/constants"
)

var (
	ErrModelConversion = errors.New("model conversion failed")
	ErrUserQuit        = constants.ErrUserQuit
)
