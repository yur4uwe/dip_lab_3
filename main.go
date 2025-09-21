package main

import (
	"fmt"
	"graph"
	"lab/correlation"
)

func main() {
	signalX := graph.IntLinearArray(0, 8)

	signalex1Y := []float64{4, 2, -1, 3, -2, -6, -5, 4, 5}
	signalex2Y := []float64{-4, 1, 3, 7, 4, -2, -8, -2, 1}

	fmt.Println("Correlation between signalex1 and signalex2:", correlation.Correlation(signalex1Y, signalex2Y, 0))
	fmt.Println("Normalized correlation between signalex1 and signalex2:", correlation.NormalizedCorrelation(signalex1Y, signalex2Y, 0))
	fmt.Println("Auto-correlation of signalex1:", correlation.AutoCorrelation(signalex1Y, 0))
	fmt.Println("Auto-correlation of signalex2:", correlation.AutoCorrelation(signalex2Y, 0))

	signal1Y := []float64{0, 3, 5, 5, 5, 2, 0.5, 0.25, 0}
	signal2Y := []float64{1, 1, 1, 1, 1, 0, 0, 0, 0}
	signal3Y := []float64{0, 9, 15, 15, 15, 6, 1.5, 0.75, 0}
	signal4Y := []float64{2, 2, 2, 2, 2, 0, 0, 0, 0}

	fmt.Println("Correlation between signal 1 and signal 3:", correlation.Correlation(signal1Y, signal3Y, 0))
	fmt.Println("Normalized correlation between signal 1 and signal 3:", correlation.NormalizedCorrelation(signal1Y, signal2Y, 0))
	fmt.Println()
	fmt.Println("Correlation between signal 2 and signal 4:", correlation.Correlation(signal2Y, signal4Y, 0))
	fmt.Println("Normalized correlation between signal 2 and signal 4:", correlation.NormalizedCorrelation(signal3Y, signal4Y, 0))
	fmt.Println()
	fmt.Println("Auto-correlation of signal 1:", correlation.AutoCorrelation(signal1Y, 0))
	fmt.Println("Auto-correlation of signal 2:", correlation.AutoCorrelation(signal2Y, 0))
	fmt.Println("Auto-correlation of signal 3:", correlation.AutoCorrelation(signal3Y, 0))
	fmt.Println("Auto-correlation of signal 4:", correlation.AutoCorrelation(signal4Y, 0))

	ls := graph.NewLS()
	ls.Solid()

	g := graph.NewGraph()
	g.SetLineStyle(ls)

	g.Plot(signalX, signalex1Y)
	g.Plot(signalX, signalex2Y)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/signal_plot_ex.png", true); err != nil {
		panic(err)
	}

	g.Clear()

	g.Plot(signalX, signal1Y)
	g.Plot(signalX, signal3Y)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/signal_plot13.png", true); err != nil {
		panic(err)
	}

	g.Clear()

	g.Plot(signalX, signal2Y)
	g.Plot(signalX, signal4Y)

	g.SetLineStyle(ls)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/signal_plot24.png", true); err != nil {
		panic(err)
	}

	g.Clear()

	phase := graph.LinearArray(0, 5, 10)
	phaseCorr13 := make([]float64, len(phase))
	phaseCorr24 := make([]float64, len(phase))
	for i, p := range phase {
		phaseCorr13[i] = correlation.Correlation(signal1Y, signal3Y, p)
		phaseCorr24[i] = correlation.Correlation(signal2Y, signal4Y, p)
	}

	g.Plot(phase, phaseCorr13)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/phase_corr_13.png", true); err != nil {
		panic(err)
	}

	g.Clear()

	g.Plot(phase, phaseCorr24)

	if err := g.Draw(); err != nil {
		panic(err)
	}

	if err := g.SavePNG("images/phase_corr_24.png", true); err != nil {
		panic(err)
	}

	g.Clear()
}
