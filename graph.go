package ncpgraph

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	chart "github.com/wcharczuk/go-chart/v2"
)

func CreateLineGraph(xname, yname string, xval []float64, yval [][]float64, legend []string, fileName string) error {
	start := xval[0]
	xlen := len(xval)
	end := xval[xlen-1]

	cs := make([]chart.Series, 0)
	for i, y := range yval {
		s := chart.ContinuousSeries{ //TimeSeries{
			Style:   chart.Style{StrokeColor: chart.GetDefaultColor(i)},
			Name:    legend[i],
			XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(start).WithEnd(end)}.Values(),
			YValues: y,
		}
		cs = append(cs, s)
	}
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{Top: 40, Left: 20, Right: 20, Bottom: 20},
		},
		XAxis: chart.XAxis{Name: xname,
			ValueFormatter: chart.IntValueFormatter, // TimeHourValueFormatter,
		},
		YAxis:  chart.YAxis{Name: yname},
		Series: cs,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendThin(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("images/" + fileName + ".png")
	if err != nil {
		fmt.Println(err)
	}

	if _, err := f.Write(buffer.Bytes()); err != nil {
		fmt.Println(err)
	}

	return err
}

func StartWebService(port string, dirPath string) {
	servPort := port
	if len(port) == 0 {
		servPort = ":8080"
	}
	servPath := dirPath
	if len(dirPath) == 0 {
		servPath = "./images"
	}
	fs := http.FileServer(http.Dir(servPath))
	fmt.Printf("Please visit http://localhost%s or http://hostip%s in browser to view image files in %s\n", servPort, servPort, servPath)
	http.ListenAndServe(servPort, fs) // ":8080"
}
