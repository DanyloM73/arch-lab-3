package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRectangle struct {
	x1, y1, x2, y2 int
}

func (op *BgRectangle) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.x1, op.y1, op.x2, op.y2), color.Black, screen.Src)
	return false
}

type Figure struct {
	x, y int
	c    color.RGBA
}

func (op *Figure) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.x-150, op.y-100, op.x+150, op.y), op.c, draw.Src)
	t.Fill(image.Rect(op.x-50, op.y, op.x+50, op.y+100), op.c, draw.Src)
	return false
}

type Move struct {
	X, Y    int
	Figures []*Figure
}

func ResetScreen(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].x += op.x
		op.Figures[i].y += op.y
		op.Figures[i].Do(t)
	}
	return false
}
