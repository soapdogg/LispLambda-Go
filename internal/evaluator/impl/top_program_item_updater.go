package impl

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
)

type topProgramItemUpdater struct {}

func NewTopProgramItemUpdater() *topProgramItemUpdater {
	return &topProgramItemUpdater{}
}

func (t topProgramItemUpdater) UpdateTopProgramItemToNextChild(programStack datamodels.ProgramItemStack) {
	if programStack.IsNotEmpty() {
		top := programStack.Pop()
		updatedTop := datamodels.NewProgramItemFromExisting(top)
		programStack.Push(updatedTop)
	}
}

var _ internal.TopProgramItemUpdater = &topProgramItemUpdater{}