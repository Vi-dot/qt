package qt

//#include "qcombobox.h"
import "C"

import "strings"

type qcombobox struct {
	qwidget
}

type QComboBox interface {
	QWidget
	AddItems(texts []string)
	Count() int
	CurrentIndex() int
	CurrentText() string
	DuplicatesEnabled() bool
	HasFrame() bool
	HidePopup()
	InsertSeparator(index int)
	IsEditable() bool
	ItemText(index int) string
	LineEdit() QLineEdit
	MaxCount() int
	MaxVisibleItems() int
	MinimumContentsLength() int
	ModelColumn() int
	RemoveItem(index int)
	SetDuplicatesEnabled(enable bool)
	SetEditable(editable bool)
	SetFrame(frame bool)
	SetItemText(index int, text string)
	SetLineEdit(edit QLineEdit)
	SetMaxCount(max int)
	SetMaxVisibleItems(maxItems int)
	SetMinimumContentsLength(characters int)
	SetModelColumn(visibleColumn int)
	ShowPopup()
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotClearEditText()
	DisconnectSlotClearEditText()
	SlotClearEditText()
	ConnectSlotSetCurrentIndex()
	DisconnectSlotSetCurrentIndex()
	SlotSetCurrentIndex(index int)
	ConnectSignalCurrentTextChanged(f func())
	DisconnectSignalCurrentTextChanged()
	SignalCurrentTextChanged() func()
	ConnectSignalEditTextChanged(f func())
	DisconnectSignalEditTextChanged()
	SignalEditTextChanged() func()
}

func (p *qcombobox) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qcombobox) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQComboBox(parent QWidget) QComboBox {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qcombobox = new(qcombobox)
	qcombobox.SetPointer(C.QComboBox_New_QWidget(parentPtr))
	qcombobox.SetObjectName("QComboBox_" + randomIdentifier())
	return qcombobox
}

func (p *qcombobox) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QComboBox_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qcombobox) AddItems(texts []string) {
	if p.Pointer() != nil {
		C.QComboBox_AddItems_QStringList(p.Pointer(), C.CString(strings.Join(texts, "|")))
	}
}

func (p *qcombobox) Count() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_Count(p.Pointer()))
}

func (p *qcombobox) CurrentIndex() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_CurrentIndex(p.Pointer()))
}

func (p *qcombobox) CurrentText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QComboBox_CurrentText(p.Pointer()))
}

func (p *qcombobox) DuplicatesEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QComboBox_DuplicatesEnabled(p.Pointer()) != 0
}

func (p *qcombobox) HasFrame() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QComboBox_HasFrame(p.Pointer()) != 0
}

func (p *qcombobox) HidePopup() {
	if p.Pointer() != nil {
		C.QComboBox_HidePopup(p.Pointer())
	}
}

func (p *qcombobox) InsertSeparator(index int) {
	if p.Pointer() != nil {
		C.QComboBox_InsertSeparator_Int(p.Pointer(), C.int(index))
	}
}

func (p *qcombobox) IsEditable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QComboBox_IsEditable(p.Pointer()) != 0
}

func (p *qcombobox) ItemText(index int) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QComboBox_ItemText_Int(p.Pointer(), C.int(index)))
}

func (p *qcombobox) LineEdit() QLineEdit {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlineedit = new(qlineedit)
		qlineedit.SetPointer(C.QComboBox_LineEdit(p.Pointer()))
		if qlineedit.ObjectName() == "" {
			qlineedit.SetObjectName("QLineEdit_" + randomIdentifier())
		}
		return qlineedit
	}
}

func (p *qcombobox) MaxCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_MaxCount(p.Pointer()))
}

func (p *qcombobox) MaxVisibleItems() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_MaxVisibleItems(p.Pointer()))
}

func (p *qcombobox) MinimumContentsLength() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_MinimumContentsLength(p.Pointer()))
}

func (p *qcombobox) ModelColumn() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QComboBox_ModelColumn(p.Pointer()))
}

func (p *qcombobox) RemoveItem(index int) {
	if p.Pointer() != nil {
		C.QComboBox_RemoveItem_Int(p.Pointer(), C.int(index))
	}
}

func (p *qcombobox) SetDuplicatesEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QComboBox_SetDuplicatesEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qcombobox) SetEditable(editable bool) {
	if p.Pointer() != nil {
		C.QComboBox_SetEditable_Bool(p.Pointer(), goBoolToCInt(editable))
	}
}

func (p *qcombobox) SetFrame(frame bool) {
	if p.Pointer() != nil {
		C.QComboBox_SetFrame_Bool(p.Pointer(), goBoolToCInt(frame))
	}
}

func (p *qcombobox) SetItemText(index int, text string) {
	if p.Pointer() != nil {
		C.QComboBox_SetItemText_Int_String(p.Pointer(), C.int(index), C.CString(text))
	}
}

func (p *qcombobox) SetLineEdit(edit QLineEdit) {
	if p.Pointer() != nil {
		var editPtr C.QtObjectPtr
		if edit != nil {
			editPtr = edit.Pointer()
		}
		C.QComboBox_SetLineEdit_QLineEdit(p.Pointer(), editPtr)
	}
}

func (p *qcombobox) SetMaxCount(max int) {
	if p.Pointer() != nil {
		C.QComboBox_SetMaxCount_Int(p.Pointer(), C.int(max))
	}
}

func (p *qcombobox) SetMaxVisibleItems(maxItems int) {
	if p.Pointer() != nil {
		C.QComboBox_SetMaxVisibleItems_Int(p.Pointer(), C.int(maxItems))
	}
}

func (p *qcombobox) SetMinimumContentsLength(characters int) {
	if p.Pointer() != nil {
		C.QComboBox_SetMinimumContentsLength_Int(p.Pointer(), C.int(characters))
	}
}

func (p *qcombobox) SetModelColumn(visibleColumn int) {
	if p.Pointer() != nil {
		C.QComboBox_SetModelColumn_Int(p.Pointer(), C.int(visibleColumn))
	}
}

func (p *qcombobox) ShowPopup() {
	if p.Pointer() != nil {
		C.QComboBox_ShowPopup(p.Pointer())
	}
}

func (p *qcombobox) ConnectSlotClear() {
	C.QComboBox_ConnectSlotClear(p.Pointer())
}

func (p *qcombobox) DisconnectSlotClear() {
	C.QComboBox_DisconnectSlotClear(p.Pointer())
}

func (p *qcombobox) SlotClear() {
	if p.Pointer() != nil {
		C.QComboBox_Clear(p.Pointer())
	}
}

func (p *qcombobox) ConnectSlotClearEditText() {
	C.QComboBox_ConnectSlotClearEditText(p.Pointer())
}

func (p *qcombobox) DisconnectSlotClearEditText() {
	C.QComboBox_DisconnectSlotClearEditText(p.Pointer())
}

func (p *qcombobox) SlotClearEditText() {
	if p.Pointer() != nil {
		C.QComboBox_ClearEditText(p.Pointer())
	}
}

func (p *qcombobox) ConnectSlotSetCurrentIndex() {
	C.QComboBox_ConnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qcombobox) DisconnectSlotSetCurrentIndex() {
	C.QComboBox_DisconnectSlotSetCurrentIndex(p.Pointer())
}

func (p *qcombobox) SlotSetCurrentIndex(index int) {
	if p.Pointer() != nil {
		C.QComboBox_SetCurrentIndex_Int(p.Pointer(), C.int(index))
	}
}

func (p *qcombobox) ConnectSignalCurrentTextChanged(f func()) {
	C.QComboBox_ConnectSignalCurrentTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentTextChanged", f)
}

func (p *qcombobox) DisconnectSignalCurrentTextChanged() {
	C.QComboBox_DisconnectSignalCurrentTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentTextChanged")
}

func (p *qcombobox) SignalCurrentTextChanged() func() {
	return getSignal(p.ObjectName(), "currentTextChanged")
}

func (p *qcombobox) ConnectSignalEditTextChanged(f func()) {
	C.QComboBox_ConnectSignalEditTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "editTextChanged", f)
}

func (p *qcombobox) DisconnectSignalEditTextChanged() {
	C.QComboBox_DisconnectSignalEditTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "editTextChanged")
}

func (p *qcombobox) SignalEditTextChanged() func() {
	return getSignal(p.ObjectName(), "editTextChanged")
}

//export callbackQComboBox
func callbackQComboBox(ptr C.QtObjectPtr, msg *C.char) {
	var qcombobox = new(qcombobox)
	qcombobox.SetPointer(ptr)
	getSignal(qcombobox.ObjectName(), C.GoString(msg))()
}
