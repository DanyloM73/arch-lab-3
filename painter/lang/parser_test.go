package lang

import (
	"image/color"
	"strings"
	"testing"

	"github.com/DanyloM73/arch-lab-3/painter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "background rectangle",
			command: "bgrect 0.1 0.1 0.9 0.9",
			op:      &painter.BgRectangle{X1: 80, Y1: 80, X2: 720, Y2: 720},
		},
		{
			name:    "figure",
			command: "figure 0.3 0.3",
			op:      &painter.Figure{X: 240, Y: 240, C: color.RGBA{R: 0, G: 0, B: 255, A: 1}},
		},
		{
			name:    "move",
			command: "move 0.2 0.2",
			op:      &painter.Move{X: 160, Y: 160},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			name:    "invalid command",
			command: "invalidcommand",
			op:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parser := &Parser{}
			ops, err := parser.Parse(strings.NewReader(tc.command))
			if tc.op == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.IsType(t, tc.op, ops[1])
				assert.Equal(t, tc.op, ops[1])
			}
		})
	}
}

func Test_parse_func(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "filling with white",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:    "filling with green",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			name:    "reseting screen",
			command: "reset",
			op:      painter.OperationFunc(painter.ResetScreen),
		},
	}

	parser := &Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))
			require.NoError(t, err)
			require.Len(t, ops, 1)
			assert.IsType(t, tc.op, ops[0])
		})
	}
}
