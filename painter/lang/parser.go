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

func parse(commandLine string) painter.Operation {
	parts := strings.Split(commandLine, " ")
	instruction := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	var iArgs []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			iArgs = append(iArgs, i)
		}
	}

	var figureOps []painter.Figure

	switch instruction {
	case "white":
		return painter.OperationFunc(painter.WhiteFill)
	case "green":
		return painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		return &painter.BgRectangle{x1: iArgs[0], y1: iArgs[1], x2: iArgs[2], y2: iArgs[3]}
	case "figure":
		clr := color.RGBA{R: 255, G: 255, B: 0, A: 1}
		figure := painter.Figure{x: iArgs[0], y: iArgs[1], c: clr}
		figureOps = append(figureOps, figure)
		return &figure
	case "move":
		return &painter.Move{x: iArgs[0], y: iArgs[1], Figures: figureOps}
	case "reset":
		figureOps = figureOps[0:0]
		return painter.OperationFunc(painter.ResetScreen)
	case "update":
		return painter.UpdateOp
	}
	return nil
}