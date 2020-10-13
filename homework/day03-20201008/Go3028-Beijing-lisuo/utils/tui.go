// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table3 := widgets.NewTable()
	table3.Rows = [][]string{
		[]string{"ID", "NAME", "CONTACT", "ADDRESS"},
		[]string{"AAA", "BBB", "CCC", "ddd"},
		[]string{"DDD", "EEE", "FFF", "ddd"},
		[]string{"GGG", "HHH", "III", "dddddddddddddddddddddddddd"},
	}
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	table3.RowSeparator = true
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(0, 30, 70, 20)
	table3.FillRow = true
	//table3.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	//table3.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	//table3.RowStyles[3] = ui.NewStyle(ui.ColorYellow)

	ui.Render(table3)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
