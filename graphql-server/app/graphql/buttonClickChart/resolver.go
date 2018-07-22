package buttonClickChart

import (
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/app/graphql/common"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/app/core"
)

type ButtonClickChart struct {
	GetButtonClickCounts core.ServiceGetButtonClickCounts
}

func (b *ButtonClickChart) ButtonClickCounts() []ButtonClickCounts {
	// todo make core.GetButtonClickCounts a injected dependency
	// todo write integration tests for it
	clickCounts := b.GetButtonClickCounts()
	buttonCLickCounts := make([]ButtonClickCounts, len(clickCounts))
	i := 0
	for k, v := range clickCounts {
		buttonCLickCounts[i] = ButtonClickCounts{
			buttonName: k,
			count: v,
		}
		i++
	}
	return buttonCLickCounts
}

func (b *ButtonClickChart) TimeRange() common.TimeRange {
	return common.TimeRange{}
}

type ButtonClickCounts struct {
	buttonName string
	count int32
}

func (b ButtonClickCounts) ButtonName() string {
	return b.buttonName
}

func (b ButtonClickCounts) Count() int32 {
	return b.count
}
