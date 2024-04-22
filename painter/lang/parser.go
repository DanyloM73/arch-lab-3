package lang

import (
	""bufio"
	"fmt"
	"image/color""
	"strconv"
	"strings"
	"io"

	"github.com/DanyloM73/arch-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	var result []painter.Operation
	for scanner.Scan() { 
		commandLine := scanner.Text()
		oprtn := parse(commandLine) 
		if oprtn == nil {
			return nil, fmt.Errorf("Failed to parse this command: %s", commandLine)
		}

		// Replace any previous BgRectangle operation with the new one
		if bgRect, ok := oprtn.(*painter.BgRectangle); ok {
			for i, oldOp := range result {
				if _, ok := oldOp.(*painter.BgRectangle); ok {
					result[i] = bgRect
					break
				}
			}
		}
		result = append(result, oprtn)
	}
	return result, nil
}
