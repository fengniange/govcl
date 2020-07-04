
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TRadioButton struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewRadioButton(owner IComponent) *TRadioButton {
    r := new(TRadioButton)
    r.instance = RadioButton_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(r, (*TRadioButton).Free)
    return r
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsRadioButton(obj interface{}) *TRadioButton {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TRadioButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsRadioButton.
func RadioButtonFromInst(inst uintptr) *TRadioButton {
    return AsRadioButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsRadioButton.
func RadioButtonFromObj(obj IObject) *TRadioButton {
    return AsRadioButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsRadioButton.
func RadioButtonFromUnsafePointer(ptr unsafe.Pointer) *TRadioButton {
    return AsRadioButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (r *TRadioButton) Free() {
    if r.instance != 0 {
        RadioButton_Free(r.instance)
        r.instance, r.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (r *TRadioButton) Instance() uintptr {
    return r.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (r *TRadioButton) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (r *TRadioButton) IsValid() bool {
    return r.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (r *TRadioButton) Is() TIs {
    return TIs(r.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (r *TRadioButton) As() TAs {
//    return TAs(r.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TRadioButtonClass() TClass {
    return RadioButton_StaticClassType()
}

// 是否可以获得焦点。
func (r *TRadioButton) CanFocus() bool {
    return RadioButton_CanFocus(r.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (r *TRadioButton) ContainsControl(Control IControl) bool {
    return RadioButton_ContainsControl(r.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (r *TRadioButton) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(RadioButton_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (r *TRadioButton) DisableAlign() {
    RadioButton_DisableAlign(r.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (r *TRadioButton) EnableAlign() {
    RadioButton_EnableAlign(r.instance)
}

// 查找子控件。
//
// Find sub controls.
func (r *TRadioButton) FindChildControl(ControlName string) *TControl {
    return AsControl(RadioButton_FindChildControl(r.instance, ControlName))
}

func (r *TRadioButton) FlipChildren(AllLevels bool) {
    RadioButton_FlipChildren(r.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (r *TRadioButton) Focused() bool {
    return RadioButton_Focused(r.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (r *TRadioButton) HandleAllocated() bool {
    return RadioButton_HandleAllocated(r.instance)
}

// 插入一个控件。
//
// Insert a control.
func (r *TRadioButton) InsertControl(AControl IControl) {
    RadioButton_InsertControl(r.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (r *TRadioButton) Invalidate() {
    RadioButton_Invalidate(r.instance)
}

// 移除一个控件。
//
// Remove a control.
func (r *TRadioButton) RemoveControl(AControl IControl) {
    RadioButton_RemoveControl(r.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (r *TRadioButton) Realign() {
    RadioButton_Realign(r.instance)
}

// 重绘。
//
// Repaint.
func (r *TRadioButton) Repaint() {
    RadioButton_Repaint(r.instance)
}

// 按比例缩放。
//
// Scale by.
func (r *TRadioButton) ScaleBy(M int32, D int32) {
    RadioButton_ScaleBy(r.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (r *TRadioButton) ScrollBy(DeltaX int32, DeltaY int32) {
    RadioButton_ScrollBy(r.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (r *TRadioButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RadioButton_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (r *TRadioButton) SetFocus() {
    RadioButton_SetFocus(r.instance)
}

// 控件更新。
//
// Update.
func (r *TRadioButton) Update() {
    RadioButton_Update(r.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (r *TRadioButton) BringToFront() {
    RadioButton_BringToFront(r.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (r *TRadioButton) ClientToScreen(Point TPoint) TPoint {
    return RadioButton_ClientToScreen(r.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (r *TRadioButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (r *TRadioButton) Dragging() bool {
    return RadioButton_Dragging(r.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (r *TRadioButton) HasParent() bool {
    return RadioButton_HasParent(r.instance)
}

// 隐藏控件。
//
// Hidden control.
func (r *TRadioButton) Hide() {
    RadioButton_Hide(r.instance)
}

// 发送一个消息。
//
// Send a message.
func (r *TRadioButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RadioButton_Perform(r.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (r *TRadioButton) Refresh() {
    RadioButton_Refresh(r.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (r *TRadioButton) ScreenToClient(Point TPoint) TPoint {
    return RadioButton_ScreenToClient(r.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (r *TRadioButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RadioButton_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (r *TRadioButton) SendToBack() {
    RadioButton_SendToBack(r.instance)
}

// 显示控件。
//
// Show control.
func (r *TRadioButton) Show() {
    RadioButton_Show(r.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (r *TRadioButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return RadioButton_GetTextBuf(r.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (r *TRadioButton) GetTextLen() int32 {
    return RadioButton_GetTextLen(r.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (r *TRadioButton) SetTextBuf(Buffer string) {
    RadioButton_SetTextBuf(r.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (r *TRadioButton) FindComponent(AName string) *TComponent {
    return AsComponent(RadioButton_FindComponent(r.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (r *TRadioButton) GetNamePath() string {
    return RadioButton_GetNamePath(r.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (r *TRadioButton) Assign(Source IObject) {
    RadioButton_Assign(r.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (r *TRadioButton) ClassType() TClass {
    return RadioButton_ClassType(r.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TRadioButton) ClassName() string {
    return RadioButton_ClassName(r.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TRadioButton) InstanceSize() int32 {
    return RadioButton_InstanceSize(r.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TRadioButton) InheritsFrom(AClass TClass) bool {
    return RadioButton_InheritsFrom(r.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (r *TRadioButton) Equals(Obj IObject) bool {
    return RadioButton_Equals(r.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TRadioButton) GetHashCode() int32 {
    return RadioButton_GetHashCode(r.instance)
}

// 文本类信息。
//
// Text information.
func (r *TRadioButton) ToString() string {
    return RadioButton_ToString(r.instance)
}

func (r *TRadioButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RadioButton_AnchorToNeighbour(r.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (r *TRadioButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RadioButton_AnchorParallel(r.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (r *TRadioButton) AnchorHorizontalCenterTo(ASibling IControl) {
    RadioButton_AnchorHorizontalCenterTo(r.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (r *TRadioButton) AnchorVerticalCenterTo(ASibling IControl) {
    RadioButton_AnchorVerticalCenterTo(r.instance, CheckPtr(ASibling))
}

func (r *TRadioButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    RadioButton_AnchorAsAlign(r.instance, ATheAlign , ASpace)
}

func (r *TRadioButton) AnchorClient(ASpace int32) {
    RadioButton_AnchorClient(r.instance, ASpace)
}

// 设置改变事件。
//
// Set changed event.
func (r *TRadioButton) SetOnChange(fn TNotifyEvent) {
    RadioButton_SetOnChange(r.instance, fn)
}

func (r *TRadioButton) Action() *TAction {
    return AsAction(RadioButton_GetAction(r.instance))
}

func (r *TRadioButton) SetAction(value IComponent) {
    RadioButton_SetAction(r.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (r *TRadioButton) Align() TAlign {
    return RadioButton_GetAlign(r.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (r *TRadioButton) SetAlign(value TAlign) {
    RadioButton_SetAlign(r.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (r *TRadioButton) Alignment() TLeftRight {
    return RadioButton_GetAlignment(r.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (r *TRadioButton) SetAlignment(value TLeftRight) {
    RadioButton_SetAlignment(r.instance, value)
}

// 获取四个角位置的锚点。
func (r *TRadioButton) Anchors() TAnchors {
    return RadioButton_GetAnchors(r.instance)
}

// 设置四个角位置的锚点。
func (r *TRadioButton) SetAnchors(value TAnchors) {
    RadioButton_SetAnchors(r.instance, value)
}

func (r *TRadioButton) BiDiMode() TBiDiMode {
    return RadioButton_GetBiDiMode(r.instance)
}

func (r *TRadioButton) SetBiDiMode(value TBiDiMode) {
    RadioButton_SetBiDiMode(r.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (r *TRadioButton) Caption() string {
    return RadioButton_GetCaption(r.instance)
}

// 设置控件标题。
//
// Set the control title.
func (r *TRadioButton) SetCaption(value string) {
    RadioButton_SetCaption(r.instance, value)
}

// 获取是否选中。
func (r *TRadioButton) Checked() bool {
    return RadioButton_GetChecked(r.instance)
}

// 设置是否选中。
func (r *TRadioButton) SetChecked(value bool) {
    RadioButton_SetChecked(r.instance, value)
}

// 获取颜色。
//
// Get color.
func (r *TRadioButton) Color() TColor {
    return RadioButton_GetColor(r.instance)
}

// 设置颜色。
//
// Set color.
func (r *TRadioButton) SetColor(value TColor) {
    RadioButton_SetColor(r.instance, value)
}

// 获取约束控件大小。
func (r *TRadioButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(RadioButton_GetConstraints(r.instance))
}

// 设置约束控件大小。
func (r *TRadioButton) SetConstraints(value *TSizeConstraints) {
    RadioButton_SetConstraints(r.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (r *TRadioButton) DoubleBuffered() bool {
    return RadioButton_GetDoubleBuffered(r.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (r *TRadioButton) SetDoubleBuffered(value bool) {
    RadioButton_SetDoubleBuffered(r.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (r *TRadioButton) DragCursor() TCursor {
    return RadioButton_GetDragCursor(r.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (r *TRadioButton) SetDragCursor(value TCursor) {
    RadioButton_SetDragCursor(r.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (r *TRadioButton) DragKind() TDragKind {
    return RadioButton_GetDragKind(r.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (r *TRadioButton) SetDragKind(value TDragKind) {
    RadioButton_SetDragKind(r.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (r *TRadioButton) DragMode() TDragMode {
    return RadioButton_GetDragMode(r.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (r *TRadioButton) SetDragMode(value TDragMode) {
    RadioButton_SetDragMode(r.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (r *TRadioButton) Enabled() bool {
    return RadioButton_GetEnabled(r.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (r *TRadioButton) SetEnabled(value bool) {
    RadioButton_SetEnabled(r.instance, value)
}

// 获取字体。
//
// Get Font.
func (r *TRadioButton) Font() *TFont {
    return AsFont(RadioButton_GetFont(r.instance))
}

// 设置字体。
//
// Set Font.
func (r *TRadioButton) SetFont(value *TFont) {
    RadioButton_SetFont(r.instance, CheckPtr(value))
}

// 获取使用父容器颜色。
//
// Get parent color.
func (r *TRadioButton) ParentColor() bool {
    return RadioButton_GetParentColor(r.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (r *TRadioButton) SetParentColor(value bool) {
    RadioButton_SetParentColor(r.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (r *TRadioButton) ParentDoubleBuffered() bool {
    return RadioButton_GetParentDoubleBuffered(r.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (r *TRadioButton) SetParentDoubleBuffered(value bool) {
    RadioButton_SetParentDoubleBuffered(r.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (r *TRadioButton) ParentFont() bool {
    return RadioButton_GetParentFont(r.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (r *TRadioButton) SetParentFont(value bool) {
    RadioButton_SetParentFont(r.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (r *TRadioButton) ParentShowHint() bool {
    return RadioButton_GetParentShowHint(r.instance)
}

// 设置以父容器的ShowHint属性为准。
func (r *TRadioButton) SetParentShowHint(value bool) {
    RadioButton_SetParentShowHint(r.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (r *TRadioButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(RadioButton_GetPopupMenu(r.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (r *TRadioButton) SetPopupMenu(value IComponent) {
    RadioButton_SetPopupMenu(r.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (r *TRadioButton) ShowHint() bool {
    return RadioButton_GetShowHint(r.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (r *TRadioButton) SetShowHint(value bool) {
    RadioButton_SetShowHint(r.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (r *TRadioButton) TabOrder() TTabOrder {
    return RadioButton_GetTabOrder(r.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (r *TRadioButton) SetTabOrder(value TTabOrder) {
    RadioButton_SetTabOrder(r.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (r *TRadioButton) TabStop() bool {
    return RadioButton_GetTabStop(r.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (r *TRadioButton) SetTabStop(value bool) {
    RadioButton_SetTabStop(r.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (r *TRadioButton) Visible() bool {
    return RadioButton_GetVisible(r.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (r *TRadioButton) SetVisible(value bool) {
    RadioButton_SetVisible(r.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (r *TRadioButton) SetOnClick(fn TNotifyEvent) {
    RadioButton_SetOnClick(r.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (r *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
    RadioButton_SetOnContextPopup(r.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (r *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
    RadioButton_SetOnDragDrop(r.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (r *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
    RadioButton_SetOnDragOver(r.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (r *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
    RadioButton_SetOnEndDrag(r.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (r *TRadioButton) SetOnEnter(fn TNotifyEvent) {
    RadioButton_SetOnEnter(r.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (r *TRadioButton) SetOnExit(fn TNotifyEvent) {
    RadioButton_SetOnExit(r.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (r *TRadioButton) SetOnKeyDown(fn TKeyEvent) {
    RadioButton_SetOnKeyDown(r.instance, fn)
}

// 设置键键下事件。
func (r *TRadioButton) SetOnKeyPress(fn TKeyPressEvent) {
    RadioButton_SetOnKeyPress(r.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (r *TRadioButton) SetOnKeyUp(fn TKeyEvent) {
    RadioButton_SetOnKeyUp(r.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (r *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
    RadioButton_SetOnMouseDown(r.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (r *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
    RadioButton_SetOnMouseEnter(r.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (r *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
    RadioButton_SetOnMouseLeave(r.instance, fn)
}

// 设置鼠标移动事件。
func (r *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
    RadioButton_SetOnMouseMove(r.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (r *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
    RadioButton_SetOnMouseUp(r.instance, fn)
}

// 获取依靠客户端总数。
func (r *TRadioButton) DockClientCount() int32 {
    return RadioButton_GetDockClientCount(r.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (r *TRadioButton) DockSite() bool {
    return RadioButton_GetDockSite(r.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (r *TRadioButton) SetDockSite(value bool) {
    RadioButton_SetDockSite(r.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (r *TRadioButton) MouseInClient() bool {
    return RadioButton_GetMouseInClient(r.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (r *TRadioButton) VisibleDockClientCount() int32 {
    return RadioButton_GetVisibleDockClientCount(r.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (r *TRadioButton) Brush() *TBrush {
    return AsBrush(RadioButton_GetBrush(r.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (r *TRadioButton) ControlCount() int32 {
    return RadioButton_GetControlCount(r.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (r *TRadioButton) Handle() HWND {
    return RadioButton_GetHandle(r.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (r *TRadioButton) ParentWindow() HWND {
    return RadioButton_GetParentWindow(r.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (r *TRadioButton) SetParentWindow(value HWND) {
    RadioButton_SetParentWindow(r.instance, value)
}

func (r *TRadioButton) Showing() bool {
    return RadioButton_GetShowing(r.instance)
}

// 获取使用停靠管理。
func (r *TRadioButton) UseDockManager() bool {
    return RadioButton_GetUseDockManager(r.instance)
}

// 设置使用停靠管理。
func (r *TRadioButton) SetUseDockManager(value bool) {
    RadioButton_SetUseDockManager(r.instance, value)
}

func (r *TRadioButton) BoundsRect() TRect {
    return RadioButton_GetBoundsRect(r.instance)
}

func (r *TRadioButton) SetBoundsRect(value TRect) {
    RadioButton_SetBoundsRect(r.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (r *TRadioButton) ClientHeight() int32 {
    return RadioButton_GetClientHeight(r.instance)
}

// 设置客户区高度。
//
// Set client height.
func (r *TRadioButton) SetClientHeight(value int32) {
    RadioButton_SetClientHeight(r.instance, value)
}

func (r *TRadioButton) ClientOrigin() TPoint {
    return RadioButton_GetClientOrigin(r.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (r *TRadioButton) ClientRect() TRect {
    return RadioButton_GetClientRect(r.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (r *TRadioButton) ClientWidth() int32 {
    return RadioButton_GetClientWidth(r.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (r *TRadioButton) SetClientWidth(value int32) {
    RadioButton_SetClientWidth(r.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (r *TRadioButton) ControlState() TControlState {
    return RadioButton_GetControlState(r.instance)
}

// 设置控件状态。
//
// Set control state.
func (r *TRadioButton) SetControlState(value TControlState) {
    RadioButton_SetControlState(r.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (r *TRadioButton) ControlStyle() TControlStyle {
    return RadioButton_GetControlStyle(r.instance)
}

// 设置控件样式。
//
// Set control style.
func (r *TRadioButton) SetControlStyle(value TControlStyle) {
    RadioButton_SetControlStyle(r.instance, value)
}

func (r *TRadioButton) Floating() bool {
    return RadioButton_GetFloating(r.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (r *TRadioButton) Parent() *TWinControl {
    return AsWinControl(RadioButton_GetParent(r.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (r *TRadioButton) SetParent(value IWinControl) {
    RadioButton_SetParent(r.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (r *TRadioButton) Left() int32 {
    return RadioButton_GetLeft(r.instance)
}

// 设置左边位置。
//
// Set Left position.
func (r *TRadioButton) SetLeft(value int32) {
    RadioButton_SetLeft(r.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (r *TRadioButton) Top() int32 {
    return RadioButton_GetTop(r.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (r *TRadioButton) SetTop(value int32) {
    RadioButton_SetTop(r.instance, value)
}

// 获取宽度。
//
// Get width.
func (r *TRadioButton) Width() int32 {
    return RadioButton_GetWidth(r.instance)
}

// 设置宽度。
//
// Set width.
func (r *TRadioButton) SetWidth(value int32) {
    RadioButton_SetWidth(r.instance, value)
}

// 获取高度。
//
// Get height.
func (r *TRadioButton) Height() int32 {
    return RadioButton_GetHeight(r.instance)
}

// 设置高度。
//
// Set height.
func (r *TRadioButton) SetHeight(value int32) {
    RadioButton_SetHeight(r.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (r *TRadioButton) Cursor() TCursor {
    return RadioButton_GetCursor(r.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (r *TRadioButton) SetCursor(value TCursor) {
    RadioButton_SetCursor(r.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (r *TRadioButton) Hint() string {
    return RadioButton_GetHint(r.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (r *TRadioButton) SetHint(value string) {
    RadioButton_SetHint(r.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (r *TRadioButton) ComponentCount() int32 {
    return RadioButton_GetComponentCount(r.instance)
}

// 获取组件索引。
//
// Get component index.
func (r *TRadioButton) ComponentIndex() int32 {
    return RadioButton_GetComponentIndex(r.instance)
}

// 设置组件索引。
//
// Set component index.
func (r *TRadioButton) SetComponentIndex(value int32) {
    RadioButton_SetComponentIndex(r.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (r *TRadioButton) Owner() *TComponent {
    return AsComponent(RadioButton_GetOwner(r.instance))
}

// 获取组件名称。
//
// Get the component name.
func (r *TRadioButton) Name() string {
    return RadioButton_GetName(r.instance)
}

// 设置组件名称。
//
// Set the component name.
func (r *TRadioButton) SetName(value string) {
    RadioButton_SetName(r.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (r *TRadioButton) Tag() int {
    return RadioButton_GetTag(r.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (r *TRadioButton) SetTag(value int) {
    RadioButton_SetTag(r.instance, value)
}

// 获取左边锚点。
func (r *TRadioButton) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(RadioButton_GetAnchorSideLeft(r.instance))
}

// 设置左边锚点。
func (r *TRadioButton) SetAnchorSideLeft(value *TAnchorSide) {
    RadioButton_SetAnchorSideLeft(r.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (r *TRadioButton) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(RadioButton_GetAnchorSideTop(r.instance))
}

// 设置顶边锚点。
func (r *TRadioButton) SetAnchorSideTop(value *TAnchorSide) {
    RadioButton_SetAnchorSideTop(r.instance, CheckPtr(value))
}

// 获取右边锚点。
func (r *TRadioButton) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(RadioButton_GetAnchorSideRight(r.instance))
}

// 设置右边锚点。
func (r *TRadioButton) SetAnchorSideRight(value *TAnchorSide) {
    RadioButton_SetAnchorSideRight(r.instance, CheckPtr(value))
}

// 获取底边锚点。
func (r *TRadioButton) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(RadioButton_GetAnchorSideBottom(r.instance))
}

// 设置底边锚点。
func (r *TRadioButton) SetAnchorSideBottom(value *TAnchorSide) {
    RadioButton_SetAnchorSideBottom(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(RadioButton_GetChildSizing(r.instance))
}

func (r *TRadioButton) SetChildSizing(value *TControlChildSizing) {
    RadioButton_SetChildSizing(r.instance, CheckPtr(value))
}

// 获取边框间距。
func (r *TRadioButton) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(RadioButton_GetBorderSpacing(r.instance))
}

// 设置边框间距。
func (r *TRadioButton) SetBorderSpacing(value *TControlBorderSpacing) {
    RadioButton_SetBorderSpacing(r.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (r *TRadioButton) DockClients(Index int32) *TControl {
    return AsControl(RadioButton_GetDockClients(r.instance, Index))
}

// 获取指定索引子控件。
func (r *TRadioButton) Controls(Index int32) *TControl {
    return AsControl(RadioButton_GetControls(r.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (r *TRadioButton) Components(AIndex int32) *TComponent {
    return AsComponent(RadioButton_GetComponents(r.instance, AIndex))
}

// 获取锚侧面。
func (r *TRadioButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(RadioButton_GetAnchorSide(r.instance, AKind))
}

