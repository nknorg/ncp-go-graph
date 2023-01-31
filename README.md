# ncp-go-graph

This is a wrap of github.com/wcharczuk/go-chart/v2 to create multi-lines graph.
It is useful to show the result of multi-slice data in a chart.

## Create Multi-Line Chart
Use this function to create multi-line chart:
```
func CreateLineGraph(xname, yname string, xval []float64, yval [][]float64, legend []string, fileName string)
```
### Params
- xname: X axis name
- yname: Y axis name
- xval: X axis values slice
- yval: Y values of multi lines, each line has a value slice.
- legend: A legend name list of each line. It will be shown on top of the chart.
- filename: The image name which is created in the subfolder **images**

### View images by browser
Under some circumstances, if you can not view the images locally, you can use this function to start a web service:
```
func StartWebService(port string, dirPath string)
```
### Params
- port: The web service port. If input empty string, it will use default port ":8080"
- dirPath: The directory of images. If input empty string, it will use default directory "./images"

After starting web service, you may view images remotely by browsing URL: http://host-ip:port
