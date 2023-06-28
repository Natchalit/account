package actions

import (
	"net/http"
	"os"

	"github.com/Natchalit/account/src/functions"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func XGetCurrency(c *gin.Context) {
	stringObj, ex := functions.GetCurrency()
	if ex != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": ex.Error()})
	}

	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Currency",
		Subtitle: "All Currency in the world",
	}))
	// Put data into instance

	units := []string{}
	for k := range stringObj.Rates {
		units = append(units, k)
	}

	bar.SetXAxis(units).
		AddSeries("Category A", generateBarItems(stringObj.Rates))
	// Where the magic happens
	f, _ := os.Create("src/pages/chart/bar.tmpl")
	bar.Render(f)

	c.JSON(http.StatusOK, bar)
}

func generateBarItems(rates map[string]float64) []opts.BarData {
	items := []opts.BarData{}
	for _, data := range rates {
		items = append(items, opts.BarData{Value: data})
	}
	return items

}
