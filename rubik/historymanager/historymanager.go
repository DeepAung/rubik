package historymanager

import (
	"fmt"

	"github.com/DeepAung/rubik/rubik/types"
)

type IHistoryManager interface {
	UpdateRotate(notation *types.Notation)
	UpdateSet(setFrom *[6][3][3]uint8, setTo *[6][3][3]uint8)
	CanUndo() bool
	CanRedo() bool
	Undo(times int, r myIRubik)
	Redo(times int, r myIRubik)
}

// to prevent cycle import
type myIRubik interface {
	SetState(state *[6][3][3]uint8, saveHistory bool)
	Rotate(notation *types.Notation, saveHistory bool) error
	RotateInverse(notation *types.Notation, saveHistory bool) error
}

type historyManager struct {
	histories     [256]history
	historyIdx    uint8
	maxHistoryIdx uint8
}

type history struct {
	HistoryType    historyType
	NotationChange types.Notation
	SetFrom        [6][3][3]uint8
	SetTo          [6][3][3]uint8
}

type historyType uint8

const (
	None   historyType = 0
	Rotate historyType = 1
	Set    historyType = 2
)

func New() IHistoryManager {
	return &historyManager{
		histories:     [256]history{},
		historyIdx:    0,
		maxHistoryIdx: 0,
	}
}

func (h *historyManager) UpdateRotate(notation *types.Notation) {
	h.historyIdx++
	h.maxHistoryIdx = h.historyIdx
	h.histories[h.historyIdx].HistoryType = Rotate
	h.histories[h.historyIdx].NotationChange = *notation
}

func (h *historyManager) UpdateSet(setFrom *[6][3][3]uint8, setTo *[6][3][3]uint8) {
	h.historyIdx++
	h.maxHistoryIdx = h.historyIdx
	h.histories[h.historyIdx].HistoryType = Set
	h.histories[h.historyIdx].SetFrom = *setFrom
	h.histories[h.historyIdx].SetTo = *setTo
}

func (h *historyManager) CanUndo() bool {
	return h.histories[h.historyIdx].HistoryType != None && h.historyIdx-1 != h.maxHistoryIdx
}

func (h *historyManager) CanRedo() bool {
	return h.historyIdx != h.maxHistoryIdx
}

func (h *historyManager) Undo(times int, r myIRubik) {
	for i := 0; i < times; i++ {
		if !h.CanUndo() {
			return
		}

		switch history := h.histories[h.historyIdx]; history.HistoryType {
		case None:
			fmt.Println("should not happen!!!")
			h.maxHistoryIdx = h.historyIdx
			return
		case Rotate:
			r.RotateInverse(&history.NotationChange, false)
		case Set:
			r.SetState(&history.SetFrom, false)
		}

		h.historyIdx--
	}
}

func (h *historyManager) Redo(times int, r myIRubik) {
	for i := 0; i < times; i++ {
		if !h.CanRedo() {
			return
		}

		switch history := h.histories[h.historyIdx+1]; history.HistoryType {
		case None:
			fmt.Println("should not happen!!!")
			h.maxHistoryIdx = h.historyIdx
			return
		case Rotate:
			r.Rotate(&history.NotationChange, false)
		case Set:
			r.SetState(&history.SetTo, false)
		}

		h.historyIdx++
	}
}
