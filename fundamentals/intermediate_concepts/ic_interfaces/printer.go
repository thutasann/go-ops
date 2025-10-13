package icinterfaces

import "fmt"

type Printer interface {
	Print() string
}

type LaserPrinter struct {
	Brand string
}

type InkjetPrinter struct {
	Brand string
}

type ThreeDPrinter struct {
	Brand string
}

func (lp LaserPrinter) Print() string {
	return fmt.Sprintf("LaserPrinter (%s): Printing with laser precision!", lp.Brand)
}

func (ip InkjetPrinter) Print() string {
	return fmt.Sprintf("InkjetPrinter (%s): Printing with liquid ink!", ip.Brand)
}

func (tp ThreeDPrinter) Print() string {
	return fmt.Sprintf("ThreeDPrinter (%s): Printing with ThreeD!", tp.Brand)
}

// A function that accepts any Printer interface
func PrintDocument(p Printer) {
	fmt.Println(p.Print())
}

func Printer_Interface_Sample() {
	fmt.Println("\n==> Printer Interface Sample")

	laser := LaserPrinter{Brand: "HP"}
	inkjet := InkjetPrinter{Brand: "Canon"}
	tdPrinter := ThreeDPrinter{Brand: "Formlabs"}

	PrintDocument(laser)
	PrintDocument(inkjet)
	PrintDocument(tdPrinter)

	printers := []Printer{laser, inkjet, tdPrinter}
	fmt.Println("\nPrinting from slice of interfaces:")
	for _, p := range printers {
		fmt.Println(p.Print())
	}
}
