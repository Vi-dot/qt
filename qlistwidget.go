package qt

//#include "qlistwidget.h"
import "C"

type qlistwidget struct {
	qlistview
}

type QListWidget interface {
	QListView
	AddItem1(label string)
	AddItem2(item QListWidgetItem)
	ClosePersistentEditor(item QListWidgetItem)
	Count() int
	CurrentItem() QListWidgetItem
	CurrentRow() int
	EditItem(item QListWidgetItem)
	InsertItem1(row int, item QListWidgetItem)
	InsertItem2(row int, label string)
	IsSortingEnabled() bool
	Item(row int) QListWidgetItem
	ItemAt(x int, y int) QListWidgetItem
	ItemWidget(item QListWidgetItem) QWidget
	OpenPersistentEditor(item QListWidgetItem)
	RemoveItemWidget(item QListWidgetItem)
	Row(item QListWidgetItem) int
	SetCurrentItem(item QListWidgetItem)
	SetCurrentRow(row int)
	SetItemWidget(item QListWidgetItem, widget QWidget)
	SetSortingEnabled(enable bool)
	SortItems(order SortOrder)
	TakeItem(row int) QListWidgetItem
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSignalCurrentItemChanged(f func())
	DisconnectSignalCurrentItemChanged()
	SignalCurrentItemChanged() func()
	ConnectSignalCurrentRowChanged(f func())
	DisconnectSignalCurrentRowChanged()
	SignalCurrentRowChanged() func()
	ConnectSignalCurrentTextChanged(f func())
	DisconnectSignalCurrentTextChanged()
	SignalCurrentTextChanged() func()
	ConnectSignalItemActivated(f func())
	DisconnectSignalItemActivated()
	SignalItemActivated() func()
	ConnectSignalItemChanged(f func())
	DisconnectSignalItemChanged()
	SignalItemChanged() func()
	ConnectSignalItemClicked(f func())
	DisconnectSignalItemClicked()
	SignalItemClicked() func()
	ConnectSignalItemDoubleClicked(f func())
	DisconnectSignalItemDoubleClicked()
	SignalItemDoubleClicked() func()
	ConnectSignalItemEntered(f func())
	DisconnectSignalItemEntered()
	SignalItemEntered() func()
	ConnectSignalItemPressed(f func())
	DisconnectSignalItemPressed()
	SignalItemPressed() func()
	ConnectSignalItemSelectionChanged(f func())
	DisconnectSignalItemSelectionChanged()
	SignalItemSelectionChanged() func()
}

func (p *qlistwidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlistwidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQListWidget(parent QWidget) QListWidget {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlistwidget = new(qlistwidget)
	qlistwidget.SetPointer(C.QListWidget_New_QWidget(parentPtr))
	qlistwidget.SetObjectName("QListWidget_" + randomIdentifier())
	return qlistwidget
}

func (p *qlistwidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QListWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlistwidget) AddItem1(label string) {
	if p.Pointer() != nil {
		C.QListWidget_AddItem_String(p.Pointer(), C.CString(label))
	}
}

func (p *qlistwidget) AddItem2(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_AddItem_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) ClosePersistentEditor(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_ClosePersistentEditor_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) Count() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListWidget_Count(p.Pointer()))
}

func (p *qlistwidget) CurrentItem() QListWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidgetitem = new(qlistwidgetitem)
		qlistwidgetitem.SetPointer(C.QListWidget_CurrentItem(p.Pointer()))
		return qlistwidgetitem
	}
}

func (p *qlistwidget) CurrentRow() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QListWidget_CurrentRow(p.Pointer()))
}

func (p *qlistwidget) EditItem(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_EditItem_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) InsertItem1(row int, item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_InsertItem_Int_QListWidgetItem(p.Pointer(), C.int(row), itemPtr)
	}
}

func (p *qlistwidget) InsertItem2(row int, label string) {
	if p.Pointer() != nil {
		C.QListWidget_InsertItem_Int_String(p.Pointer(), C.int(row), C.CString(label))
	}
}

func (p *qlistwidget) IsSortingEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QListWidget_IsSortingEnabled(p.Pointer()) != 0
}

func (p *qlistwidget) Item(row int) QListWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidgetitem = new(qlistwidgetitem)
		qlistwidgetitem.SetPointer(C.QListWidget_Item_Int(p.Pointer(), C.int(row)))
		return qlistwidgetitem
	}
}

func (p *qlistwidget) ItemAt(x int, y int) QListWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidgetitem = new(qlistwidgetitem)
		qlistwidgetitem.SetPointer(C.QListWidget_ItemAt_Int_Int(p.Pointer(), C.int(x), C.int(y)))
		return qlistwidgetitem
	}
}

func (p *qlistwidget) ItemWidget(item QListWidgetItem) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QListWidget_ItemWidget_QListWidgetItem(p.Pointer(), itemPtr))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlistwidget) OpenPersistentEditor(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_OpenPersistentEditor_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) RemoveItemWidget(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_RemoveItemWidget_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) Row(item QListWidgetItem) int {
	if p.Pointer() == nil {
		return 0
	} else {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		return int(C.QListWidget_Row_QListWidgetItem(p.Pointer(), itemPtr))
	}
}

func (p *qlistwidget) SetCurrentItem(item QListWidgetItem) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		C.QListWidget_SetCurrentItem_QListWidgetItem(p.Pointer(), itemPtr)
	}
}

func (p *qlistwidget) SetCurrentRow(row int) {
	if p.Pointer() != nil {
		C.QListWidget_SetCurrentRow_Int(p.Pointer(), C.int(row))
	}
}

func (p *qlistwidget) SetItemWidget(item QListWidgetItem, widget QWidget) {
	if p.Pointer() != nil {
		var itemPtr C.QtObjectPtr
		if item != nil {
			itemPtr = item.Pointer()
		}
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QListWidget_SetItemWidget_QListWidgetItem_QWidget(p.Pointer(), itemPtr, widgetPtr)
	}
}

func (p *qlistwidget) SetSortingEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QListWidget_SetSortingEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qlistwidget) SortItems(order SortOrder) {
	if p.Pointer() != nil {
		C.QListWidget_SortItems_SortOrder(p.Pointer(), C.int(order))
	}
}

func (p *qlistwidget) TakeItem(row int) QListWidgetItem {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlistwidgetitem = new(qlistwidgetitem)
		qlistwidgetitem.SetPointer(C.QListWidget_TakeItem_Int(p.Pointer(), C.int(row)))
		return qlistwidgetitem
	}
}

func (p *qlistwidget) ConnectSlotClear() {
	C.QListWidget_ConnectSlotClear(p.Pointer())
}

func (p *qlistwidget) DisconnectSlotClear() {
	C.QListWidget_DisconnectSlotClear(p.Pointer())
}

func (p *qlistwidget) SlotClear() {
	if p.Pointer() != nil {
		C.QListWidget_Clear(p.Pointer())
	}
}

func (p *qlistwidget) ConnectSignalCurrentItemChanged(f func()) {
	C.QListWidget_ConnectSignalCurrentItemChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentItemChanged", f)
}

func (p *qlistwidget) DisconnectSignalCurrentItemChanged() {
	C.QListWidget_DisconnectSignalCurrentItemChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentItemChanged")
}

func (p *qlistwidget) SignalCurrentItemChanged() func() {
	return getSignal(p.ObjectName(), "currentItemChanged")
}

func (p *qlistwidget) ConnectSignalCurrentRowChanged(f func()) {
	C.QListWidget_ConnectSignalCurrentRowChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentRowChanged", f)
}

func (p *qlistwidget) DisconnectSignalCurrentRowChanged() {
	C.QListWidget_DisconnectSignalCurrentRowChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentRowChanged")
}

func (p *qlistwidget) SignalCurrentRowChanged() func() {
	return getSignal(p.ObjectName(), "currentRowChanged")
}

func (p *qlistwidget) ConnectSignalCurrentTextChanged(f func()) {
	C.QListWidget_ConnectSignalCurrentTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentTextChanged", f)
}

func (p *qlistwidget) DisconnectSignalCurrentTextChanged() {
	C.QListWidget_DisconnectSignalCurrentTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentTextChanged")
}

func (p *qlistwidget) SignalCurrentTextChanged() func() {
	return getSignal(p.ObjectName(), "currentTextChanged")
}

func (p *qlistwidget) ConnectSignalItemActivated(f func()) {
	C.QListWidget_ConnectSignalItemActivated(p.Pointer())
	connectSignal(p.ObjectName(), "itemActivated", f)
}

func (p *qlistwidget) DisconnectSignalItemActivated() {
	C.QListWidget_DisconnectSignalItemActivated(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemActivated")
}

func (p *qlistwidget) SignalItemActivated() func() {
	return getSignal(p.ObjectName(), "itemActivated")
}

func (p *qlistwidget) ConnectSignalItemChanged(f func()) {
	C.QListWidget_ConnectSignalItemChanged(p.Pointer())
	connectSignal(p.ObjectName(), "itemChanged", f)
}

func (p *qlistwidget) DisconnectSignalItemChanged() {
	C.QListWidget_DisconnectSignalItemChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemChanged")
}

func (p *qlistwidget) SignalItemChanged() func() {
	return getSignal(p.ObjectName(), "itemChanged")
}

func (p *qlistwidget) ConnectSignalItemClicked(f func()) {
	C.QListWidget_ConnectSignalItemClicked(p.Pointer())
	connectSignal(p.ObjectName(), "itemClicked", f)
}

func (p *qlistwidget) DisconnectSignalItemClicked() {
	C.QListWidget_DisconnectSignalItemClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemClicked")
}

func (p *qlistwidget) SignalItemClicked() func() {
	return getSignal(p.ObjectName(), "itemClicked")
}

func (p *qlistwidget) ConnectSignalItemDoubleClicked(f func()) {
	C.QListWidget_ConnectSignalItemDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "itemDoubleClicked", f)
}

func (p *qlistwidget) DisconnectSignalItemDoubleClicked() {
	C.QListWidget_DisconnectSignalItemDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemDoubleClicked")
}

func (p *qlistwidget) SignalItemDoubleClicked() func() {
	return getSignal(p.ObjectName(), "itemDoubleClicked")
}

func (p *qlistwidget) ConnectSignalItemEntered(f func()) {
	C.QListWidget_ConnectSignalItemEntered(p.Pointer())
	connectSignal(p.ObjectName(), "itemEntered", f)
}

func (p *qlistwidget) DisconnectSignalItemEntered() {
	C.QListWidget_DisconnectSignalItemEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemEntered")
}

func (p *qlistwidget) SignalItemEntered() func() {
	return getSignal(p.ObjectName(), "itemEntered")
}

func (p *qlistwidget) ConnectSignalItemPressed(f func()) {
	C.QListWidget_ConnectSignalItemPressed(p.Pointer())
	connectSignal(p.ObjectName(), "itemPressed", f)
}

func (p *qlistwidget) DisconnectSignalItemPressed() {
	C.QListWidget_DisconnectSignalItemPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemPressed")
}

func (p *qlistwidget) SignalItemPressed() func() {
	return getSignal(p.ObjectName(), "itemPressed")
}

func (p *qlistwidget) ConnectSignalItemSelectionChanged(f func()) {
	C.QListWidget_ConnectSignalItemSelectionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "itemSelectionChanged", f)
}

func (p *qlistwidget) DisconnectSignalItemSelectionChanged() {
	C.QListWidget_DisconnectSignalItemSelectionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemSelectionChanged")
}

func (p *qlistwidget) SignalItemSelectionChanged() func() {
	return getSignal(p.ObjectName(), "itemSelectionChanged")
}

//export callbackQListWidget
func callbackQListWidget(ptr C.QtObjectPtr, msg *C.char) {
	var qlistwidget = new(qlistwidget)
	qlistwidget.SetPointer(ptr)
	getSignal(qlistwidget.ObjectName(), C.GoString(msg))()
}
