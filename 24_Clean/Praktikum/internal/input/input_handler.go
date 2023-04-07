package input

import (
	"bufio"
	"fmt"
	"io"
)

type InputHandler struct {
	reader *bufio.Reader
	writer io.Writer
}

func NewInputHandler(reader io.Reader, writer io.Writer) *InputHandler {
	return &InputHandler{
		reader: bufio.NewReader(reader),
		writer: writer,
	}
}

func (h *InputHandler) GetPhoneNumber() (string, error) {
	phoneNumber, err := h.readLine()
	if err != nil {
		return "", err
	}

	return phoneNumber, nil
}

func (h *InputHandler) readLine() (string, error) {
	line, err := h.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return line[:len(line)-1], nil
}

func (h *InputHandler) WriteString(s string) {
	fmt.Fprint(h.writer, s)
}
