package main

import (
	colors "calc/calc/Colors"
    "calc/calc/Eval"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 280
	screenHeight = 500
	buttonSize   = 70
	gapSize      = 2
	numButtons   = 19
)

type Game struct {
	buttonPositions [numButtons][2]int
	buttonColors    [numButtons]color.RGBA
	characterPrinted bool       
}

var characters = []string{"C", "Del", "+/-", "%", "7", "8", "9", "/", "4", "5", "6", "+", "1", "2", "3", "*", "0", ".", "="}

var query = ""
var result = ""

func NewGame() *Game {
	g := &Game{}
	g.initializeButtons()
	return g
}

func (g *Game) initializeButtons() {
	cols := 4
    startX := 0
    startY := 150

	for i := 0; i < numButtons; i++ {
		row := i / cols
		col := i % cols
		x := startX + col*(buttonSize+gapSize)
		y := startY + row*(buttonSize+gapSize)
        if characters[i] == "=" || characters[i] == "."{
		    g.buttonPositions[i] = [2]int{x+72, y}
        }else {
		   g.buttonPositions[i] = [2]int{x, y}
        }
		g.buttonColors[i] = colors.LightGray
	}
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()

	for i, pos := range g.buttonPositions {
		bx, by := pos[0], pos[1]
		if (x >= bx && x <= bx+buttonSize && y >= by && y <= by+buttonSize){
			g.buttonColors[i] = colors.DarkGray
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
				if !g.characterPrinted {
                    if characters[i] != "="{
                        if query == "Division by zero error" || query == "Invalid operator" || characters[i] == "C"{
                            query = ""
					        g.characterPrinted = true
                            continue
                        }

                        if characters[i] == "Del" && len(query) > 0{
                            query = query[:len(query) - 1]
					        g.characterPrinted = true
                            continue
                        } else if characters[i] == "Del" && len(query) == 0{
                            query = ""
					        g.characterPrinted = true
                            continue
                        }
                        query += characters[i]
                    } else {
                        query = eval.Eval(query)
                    }
					g.characterPrinted = true
				}
			}
		} else {
			g.buttonColors[i] = colors.LightGray
		}
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.characterPrinted = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, pos := range g.buttonPositions {
		bx, by := pos[0], pos[1]
		buttonImage := ebiten.NewImage(buttonSize, buttonSize)
        if characters[i] == "0"{
		    buttonImage = ebiten.NewImage(buttonSize+72, buttonSize)
        }
		buttonImage.Fill(g.buttonColors[i])
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(bx), float64(by))
		screen.DrawImage(buttonImage, op)

		textX := bx + 27
		textY := by + 27
        if characters[i] == "0"{
            textX = bx + 65
        }

		ebitenutil.DebugPrintAt(screen, characters[i], textX, textY)

        ebitenutil.DebugPrintAt(screen, query, 220, 80)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 280, 500
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

