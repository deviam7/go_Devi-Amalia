package output

type OutputHandler interface {
	WriteString(s string)
}
