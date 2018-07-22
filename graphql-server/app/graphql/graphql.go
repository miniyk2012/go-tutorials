package graphql

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/app/graphql/buttonClickChart"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/app/core"
)

type Query struct{}

func (q *Query) Chart1() *Chart1 { return &Chart1{} }

type Chart1 struct {}

func (m Chart1) Points() []*Point {
	return []*Point{
		{a: "C", b: 2},
		{a: "C", b: 7},
		{a: "C", b: 4},
		{a: "D", b: 1},
		{a: "D", b: 3},
	}
}

type Point struct {
	a string
	b int32
}

func (p *Point) A() string {return p.a}
func (p *Point) B() int32 {return p.b}


type ButtonClickChartArg struct{
	TimeRange *string
}

func (q *Query) ButtonClickChart(arg *ButtonClickChartArg) *buttonClickChart.ButtonClickChart {
	return &buttonClickChart.ButtonClickChart{
		GetButtonClickCounts: core.GetButtonClickCounts,
	}
}
