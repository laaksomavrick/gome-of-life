package gol

import (
	"github.com/gdamore/tcell"
	"log"
	"os"
	"time"
)

const (
	cellRune = rune('.')
	tickRate = time.Second * 1
)

type Renderer struct {
	screen tcell.Screen
	queue chan[][]bool
	terminalEventChan chan tcell.Event
	backgroundStyle tcell.Style
	cellStyle tcell.Style
}

func NewRenderer(queue chan[][]bool) Renderer {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	terminalEventChan := make(chan tcell.Event)

	backgroundStyle := tcell.StyleDefault.Foreground(
		tcell.ColorBlack).Background(tcell.ColorBlack)

	cellStyle := tcell.StyleDefault.Foreground(
		tcell.ColorWhiteSmoke).Background(tcell.ColorWhiteSmoke)

	return Renderer{
		screen: screen,
		queue: queue,
		terminalEventChan: terminalEventChan,
		backgroundStyle: backgroundStyle,
		cellStyle: cellStyle,
	}
}

func (r *Renderer) Init() {
	err := r.screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	r.pumpEvents()
	r.setStyle()
	r.setKeyEventHandlers()

}

func (r *Renderer) Read() {
	for {
		select {
		case newState := <-r.queue:
			r.render(newState)
			r.screen.Show()
			time.Sleep(tickRate)
		default:
			r.screen.Show()
			time.Sleep(tickRate)
		}
	}
}

func (r *Renderer) render(state [][]bool) {
	for y := range state {
		for x := range state[y] {
			var style tcell.Style
			value := state[y][x]
			if value {
				style = r.cellStyle
			} else {
				style = r.backgroundStyle
			}
			r.screen.SetCell(x, y, style, cellRune)
		}
	}
}

func (r *Renderer) Size() (x int, y int) {
	y, x = r.screen.Size()
	return x, y
}

func (r *Renderer) setStyle() {
	r.screen.HideCursor()
	r.screen.SetStyle(r.backgroundStyle)
	r.screen.Clear()
}

func (r *Renderer) pumpEvents() {
	go func() {
		for {
			event := r.screen.PollEvent()
			r.terminalEventChan <- event
		}
	}()
}

func (r *Renderer) setKeyEventHandlers() {
	go func() {
		for {
			select {
			case event := <-r.terminalEventChan:
				switch ev := event.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyCtrlZ, tcell.KeyCtrlC:
						r.close()
					}
				}
			default:
				continue
			}
		}
	}()
}

func (r *Renderer) close() {
	r.screen.Fini()
	os.Exit(0)
}
