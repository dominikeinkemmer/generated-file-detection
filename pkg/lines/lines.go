package lines

type Line interface {
	HasLines() bool
	Len() int
	Last(n int) []string
	First(n int) []string
	All() []string
}

// HasLines checks if the file we read is empty or not
func (l *lines) HasLines() bool {
	return len(l.lines) > 0
}

// Len returns the number of lines
func (l *lines) Len() int {
	return len(l.lines)
}

// All returns all lines of the file
func (l *lines) All() []string {
	return l.lines
}

// Last returns the last n number of lines
func (l *lines) Last(n int) []string {
	lines := make([]string, n)

	lineLength := len(l.lines)
	for i := lineLength; i > (lineLength - n); i -= 1 {
		lines = append(lines, l.lines[i])
	}

	return lines
}

// First returns the first n number of lines
func (l *lines) First(n int) []string {
	lines := make([]string, n)

	for i := 0; i < n; i += 1 {
		lines = append(lines, l.lines[i])
	}

	return lines
}
