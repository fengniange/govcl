
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

type TScrollBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewScrollBox(owner IComponent) *TScrollBox {
    s := new(TScrollBox)
    s.instance = ScrollBox_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(s, (*TScrollBox).Free)
    return s
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsScrollBox(obj interface{}) *TScrollBox {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TScrollBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsScrollBox.
func ScrollBoxFromInst(inst uintptr) *TScrollBox {
    return AsScrollBox(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsScrollBox.
func ScrollBoxFromObj(obj IObject) *TScrollBox {
    return AsScrollBox(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsScrollBox.
func ScrollBoxFromUnsafePointer(ptr unsafe.Pointer) *TScrollBox {
    return AsScrollBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (s *TScrollBox) Free() {
    if s.instance != 0 {
        ScrollBox_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (s *TScrollBox) Instance() uintptr {
    return s.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (s *TScrollBox) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (s *TScrollBox) IsValid() bool {
    return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (s *TScrollBox) Is() TIs {
    return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (s *TScrollBox) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TScrollBoxClass() TClass {
    return ScrollBox_StaticClassType()
}

func (s *TScrollBox) ScrollInView(AControl IControl) {
    ScrollBox_ScrollInView(s.instance, CheckPtr(AControl))
}

// 是否可以获得焦点。
func (s *TScrollBox) CanFocus() bool {
    return ScrollBox_CanFocus(s.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (s *TScrollBox) ContainsControl(Control IControl) bool {
    return ScrollBox_ContainsControl(s.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (s *TScrollBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ScrollBox_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (s *TScrollBox) DisableAlign() {
    ScrollBox_DisableAlign(s.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (s *TScrollBox) EnableAlign() {
    ScrollBox_EnableAlign(s.instance)
}

// 查找子控件。
//
// Find sub controls.
func (s *TScrollBox) FindChildControl(ControlName string) *TControl {
    return AsControl(ScrollBox_FindChildControl(s.instance, ControlName))
}

func (s *TScrollBox) FlipChildren(AllLevels bool) {
    ScrollBox_FlipChildren(s.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (s *TScrollBox) Focused() bool {
    return ScrollBox_Focused(s.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (s *TScrollBox) HandleAllocated() bool {
    return ScrollBox_HandleAllocated(s.instance)
}

// 插入一个控件。
//
// Insert a control.
func (s *TScrollBox) InsertControl(AControl IControl) {
    ScrollBox_InsertControl(s.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (s *TScrollBox) Invalidate() {
    ScrollBox_Invalidate(s.instance)
}

// 移除一个控件。
//
// Remove a control.
func (s *TScrollBox) RemoveControl(AControl IControl) {
    ScrollBox_RemoveControl(s.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (s *TScrollBox) Realign() {
    ScrollBox_Realign(s.instance)
}

// 重绘。
//
// Repaint.
func (s *TScrollBox) Repaint() {
    ScrollBox_Repaint(s.instance)
}

// 按比例缩放。
//
// Scale by.
func (s *TScrollBox) ScaleBy(M int32, D int32) {
    ScrollBox_ScaleBy(s.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (s *TScrollBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ScrollBox_ScrollBy(s.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (s *TScrollBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBox_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (s *TScrollBox) SetFocus() {
    ScrollBox_SetFocus(s.instance)
}

// 控件更新。
//
// Update.
func (s *TScrollBox) Update() {
    ScrollBox_Update(s.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (s *TScrollBox) BringToFront() {
    ScrollBox_BringToFront(s.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TScrollBox) ClientToScreen(Point TPoint) TPoint {
    return ScrollBox_ClientToScreen(s.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TScrollBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TScrollBox) Dragging() bool {
    return ScrollBox_Dragging(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TScrollBox) HasParent() bool {
    return ScrollBox_HasParent(s.instance)
}

// 隐藏控件。
//
// Hidden control.
func (s *TScrollBox) Hide() {
    ScrollBox_Hide(s.instance)
}

// 发送一个消息。
//
// Send a message.
func (s *TScrollBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBox_Perform(s.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (s *TScrollBox) Refresh() {
    ScrollBox_Refresh(s.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TScrollBox) ScreenToClient(Point TPoint) TPoint {
    return ScrollBox_ScreenToClient(s.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TScrollBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (s *TScrollBox) SendToBack() {
    ScrollBox_SendToBack(s.instance)
}

// 显示控件。
//
// Show control.
func (s *TScrollBox) Show() {
    ScrollBox_Show(s.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TScrollBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ScrollBox_GetTextBuf(s.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TScrollBox) GetTextLen() int32 {
    return ScrollBox_GetTextLen(s.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TScrollBox) SetTextBuf(Buffer string) {
    ScrollBox_SetTextBuf(s.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TScrollBox) FindComponent(AName string) *TComponent {
    return AsComponent(ScrollBox_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TScrollBox) GetNamePath() string {
    return ScrollBox_GetNamePath(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TScrollBox) Assign(Source IObject) {
    ScrollBox_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TScrollBox) ClassType() TClass {
    return ScrollBox_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TScrollBox) ClassName() string {
    return ScrollBox_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TScrollBox) InstanceSize() int32 {
    return ScrollBox_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TScrollBox) InheritsFrom(AClass TClass) bool {
    return ScrollBox_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TScrollBox) Equals(Obj IObject) bool {
    return ScrollBox_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TScrollBox) GetHashCode() int32 {
    return ScrollBox_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TScrollBox) ToString() string {
    return ScrollBox_ToString(s.instance)
}

func (s *TScrollBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ScrollBox_AnchorToNeighbour(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (s *TScrollBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ScrollBox_AnchorParallel(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (s *TScrollBox) AnchorHorizontalCenterTo(ASibling IControl) {
    ScrollBox_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (s *TScrollBox) AnchorVerticalCenterTo(ASibling IControl) {
    ScrollBox_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TScrollBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ScrollBox_AnchorAsAlign(s.instance, ATheAlign , ASpace)
}

func (s *TScrollBox) AnchorClient(ASpace int32) {
    ScrollBox_AnchorClient(s.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TScrollBox) Align() TAlign {
    return ScrollBox_GetAlign(s.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TScrollBox) SetAlign(value TAlign) {
    ScrollBox_SetAlign(s.instance, value)
}

// 获取四个角位置的锚点。
func (s *TScrollBox) Anchors() TAnchors {
    return ScrollBox_GetAnchors(s.instance)
}

// 设置四个角位置的锚点。
func (s *TScrollBox) SetAnchors(value TAnchors) {
    ScrollBox_SetAnchors(s.instance, value)
}

func (s *TScrollBox) AutoScroll() bool {
    return ScrollBox_GetAutoScroll(s.instance)
}

func (s *TScrollBox) SetAutoScroll(value bool) {
    ScrollBox_SetAutoScroll(s.instance, value)
}

// 获取自动调整大小。
func (s *TScrollBox) AutoSize() bool {
    return ScrollBox_GetAutoSize(s.instance)
}

// 设置自动调整大小。
func (s *TScrollBox) SetAutoSize(value bool) {
    ScrollBox_SetAutoSize(s.instance, value)
}

func (s *TScrollBox) BiDiMode() TBiDiMode {
    return ScrollBox_GetBiDiMode(s.instance)
}

func (s *TScrollBox) SetBiDiMode(value TBiDiMode) {
    ScrollBox_SetBiDiMode(s.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (s *TScrollBox) BorderStyle() TBorderStyle {
    return ScrollBox_GetBorderStyle(s.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (s *TScrollBox) SetBorderStyle(value TBorderStyle) {
    ScrollBox_SetBorderStyle(s.instance, value)
}

// 获取约束控件大小。
func (s *TScrollBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ScrollBox_GetConstraints(s.instance))
}

// 设置约束控件大小。
func (s *TScrollBox) SetConstraints(value *TSizeConstraints) {
    ScrollBox_SetConstraints(s.instance, CheckPtr(value))
}

// 获取停靠站点。
//
// Get Docking site.
func (s *TScrollBox) DockSite() bool {
    return ScrollBox_GetDockSite(s.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (s *TScrollBox) SetDockSite(value bool) {
    ScrollBox_SetDockSite(s.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (s *TScrollBox) DoubleBuffered() bool {
    return ScrollBox_GetDoubleBuffered(s.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (s *TScrollBox) SetDoubleBuffered(value bool) {
    ScrollBox_SetDoubleBuffered(s.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (s *TScrollBox) DragCursor() TCursor {
    return ScrollBox_GetDragCursor(s.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (s *TScrollBox) SetDragCursor(value TCursor) {
    ScrollBox_SetDragCursor(s.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (s *TScrollBox) DragKind() TDragKind {
    return ScrollBox_GetDragKind(s.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (s *TScrollBox) SetDragKind(value TDragKind) {
    ScrollBox_SetDragKind(s.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (s *TScrollBox) DragMode() TDragMode {
    return ScrollBox_GetDragMode(s.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (s *TScrollBox) SetDragMode(value TDragMode) {
    ScrollBox_SetDragMode(s.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (s *TScrollBox) Enabled() bool {
    return ScrollBox_GetEnabled(s.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (s *TScrollBox) SetEnabled(value bool) {
    ScrollBox_SetEnabled(s.instance, value)
}

// 获取颜色。
//
// Get color.
func (s *TScrollBox) Color() TColor {
    return ScrollBox_GetColor(s.instance)
}

// 设置颜色。
//
// Set color.
func (s *TScrollBox) SetColor(value TColor) {
    ScrollBox_SetColor(s.instance, value)
}

// 获取字体。
//
// Get Font.
func (s *TScrollBox) Font() *TFont {
    return AsFont(ScrollBox_GetFont(s.instance))
}

// 设置字体。
//
// Set Font.
func (s *TScrollBox) SetFont(value *TFont) {
    ScrollBox_SetFont(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ParentBackground() bool {
    return ScrollBox_GetParentBackground(s.instance)
}

func (s *TScrollBox) SetParentBackground(value bool) {
    ScrollBox_SetParentBackground(s.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (s *TScrollBox) ParentColor() bool {
    return ScrollBox_GetParentColor(s.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (s *TScrollBox) SetParentColor(value bool) {
    ScrollBox_SetParentColor(s.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (s *TScrollBox) ParentDoubleBuffered() bool {
    return ScrollBox_GetParentDoubleBuffered(s.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (s *TScrollBox) SetParentDoubleBuffered(value bool) {
    ScrollBox_SetParentDoubleBuffered(s.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (s *TScrollBox) ParentFont() bool {
    return ScrollBox_GetParentFont(s.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (s *TScrollBox) SetParentFont(value bool) {
    ScrollBox_SetParentFont(s.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (s *TScrollBox) ParentShowHint() bool {
    return ScrollBox_GetParentShowHint(s.instance)
}

// 设置以父容器的ShowHint属性为准。
func (s *TScrollBox) SetParentShowHint(value bool) {
    ScrollBox_SetParentShowHint(s.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (s *TScrollBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ScrollBox_GetPopupMenu(s.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (s *TScrollBox) SetPopupMenu(value IComponent) {
    ScrollBox_SetPopupMenu(s.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TScrollBox) ShowHint() bool {
    return ScrollBox_GetShowHint(s.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TScrollBox) SetShowHint(value bool) {
    ScrollBox_SetShowHint(s.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (s *TScrollBox) TabOrder() TTabOrder {
    return ScrollBox_GetTabOrder(s.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (s *TScrollBox) SetTabOrder(value TTabOrder) {
    ScrollBox_SetTabOrder(s.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (s *TScrollBox) TabStop() bool {
    return ScrollBox_GetTabStop(s.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (s *TScrollBox) SetTabStop(value bool) {
    ScrollBox_SetTabStop(s.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (s *TScrollBox) Visible() bool {
    return ScrollBox_GetVisible(s.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (s *TScrollBox) SetVisible(value bool) {
    ScrollBox_SetVisible(s.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (s *TScrollBox) SetOnClick(fn TNotifyEvent) {
    ScrollBox_SetOnClick(s.instance, fn)
}

// 设置双击事件。
func (s *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
    ScrollBox_SetOnDblClick(s.instance, fn)
}

func (s *TScrollBox) SetOnDockDrop(fn TDockDropEvent) {
    ScrollBox_SetOnDockDrop(s.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (s *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
    ScrollBox_SetOnDragDrop(s.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (s *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
    ScrollBox_SetOnDragOver(s.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (s *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
    ScrollBox_SetOnEndDrag(s.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (s *TScrollBox) SetOnEnter(fn TNotifyEvent) {
    ScrollBox_SetOnEnter(s.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (s *TScrollBox) SetOnExit(fn TNotifyEvent) {
    ScrollBox_SetOnExit(s.instance, fn)
}

func (s *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    ScrollBox_SetOnGetSiteInfo(s.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (s *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
    ScrollBox_SetOnMouseDown(s.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (s *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBox_SetOnMouseEnter(s.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (s *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBox_SetOnMouseLeave(s.instance, fn)
}

// 设置鼠标移动事件。
func (s *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ScrollBox_SetOnMouseMove(s.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (s *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
    ScrollBox_SetOnMouseUp(s.instance, fn)
}

// 设置鼠标滚轮事件。
func (s *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
    ScrollBox_SetOnMouseWheel(s.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (s *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelDown(s.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (s *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelUp(s.instance, fn)
}

// 设置大小被改变事件。
func (s *TScrollBox) SetOnResize(fn TNotifyEvent) {
    ScrollBox_SetOnResize(s.instance, fn)
}

func (s *TScrollBox) SetOnUnDock(fn TUnDockEvent) {
    ScrollBox_SetOnUnDock(s.instance, fn)
}

// 设置对齐位置事件，当Align为alCustom时Parent会收到这个消息。
func (s *TScrollBox) SetOnAlignPosition(fn TAlignPositionEvent) {
    ScrollBox_SetOnAlignPosition(s.instance, fn)
}

func (s *TScrollBox) HorzScrollBar() *TControlScrollBar {
    return AsControlScrollBar(ScrollBox_GetHorzScrollBar(s.instance))
}

func (s *TScrollBox) SetHorzScrollBar(value *TControlScrollBar) {
    ScrollBox_SetHorzScrollBar(s.instance, CheckPtr(value))
}

func (s *TScrollBox) VertScrollBar() *TControlScrollBar {
    return AsControlScrollBar(ScrollBox_GetVertScrollBar(s.instance))
}

func (s *TScrollBox) SetVertScrollBar(value *TControlScrollBar) {
    ScrollBox_SetVertScrollBar(s.instance, CheckPtr(value))
}

// 获取依靠客户端总数。
func (s *TScrollBox) DockClientCount() int32 {
    return ScrollBox_GetDockClientCount(s.instance)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (s *TScrollBox) MouseInClient() bool {
    return ScrollBox_GetMouseInClient(s.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (s *TScrollBox) VisibleDockClientCount() int32 {
    return ScrollBox_GetVisibleDockClientCount(s.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (s *TScrollBox) Brush() *TBrush {
    return AsBrush(ScrollBox_GetBrush(s.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (s *TScrollBox) ControlCount() int32 {
    return ScrollBox_GetControlCount(s.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (s *TScrollBox) Handle() HWND {
    return ScrollBox_GetHandle(s.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (s *TScrollBox) ParentWindow() HWND {
    return ScrollBox_GetParentWindow(s.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (s *TScrollBox) SetParentWindow(value HWND) {
    ScrollBox_SetParentWindow(s.instance, value)
}

func (s *TScrollBox) Showing() bool {
    return ScrollBox_GetShowing(s.instance)
}

// 获取使用停靠管理。
func (s *TScrollBox) UseDockManager() bool {
    return ScrollBox_GetUseDockManager(s.instance)
}

// 设置使用停靠管理。
func (s *TScrollBox) SetUseDockManager(value bool) {
    ScrollBox_SetUseDockManager(s.instance, value)
}

func (s *TScrollBox) Action() *TAction {
    return AsAction(ScrollBox_GetAction(s.instance))
}

func (s *TScrollBox) SetAction(value IComponent) {
    ScrollBox_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBox) BoundsRect() TRect {
    return ScrollBox_GetBoundsRect(s.instance)
}

func (s *TScrollBox) SetBoundsRect(value TRect) {
    ScrollBox_SetBoundsRect(s.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (s *TScrollBox) ClientHeight() int32 {
    return ScrollBox_GetClientHeight(s.instance)
}

// 设置客户区高度。
//
// Set client height.
func (s *TScrollBox) SetClientHeight(value int32) {
    ScrollBox_SetClientHeight(s.instance, value)
}

func (s *TScrollBox) ClientOrigin() TPoint {
    return ScrollBox_GetClientOrigin(s.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (s *TScrollBox) ClientRect() TRect {
    return ScrollBox_GetClientRect(s.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (s *TScrollBox) ClientWidth() int32 {
    return ScrollBox_GetClientWidth(s.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (s *TScrollBox) SetClientWidth(value int32) {
    ScrollBox_SetClientWidth(s.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (s *TScrollBox) ControlState() TControlState {
    return ScrollBox_GetControlState(s.instance)
}

// 设置控件状态。
//
// Set control state.
func (s *TScrollBox) SetControlState(value TControlState) {
    ScrollBox_SetControlState(s.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (s *TScrollBox) ControlStyle() TControlStyle {
    return ScrollBox_GetControlStyle(s.instance)
}

// 设置控件样式。
//
// Set control style.
func (s *TScrollBox) SetControlStyle(value TControlStyle) {
    ScrollBox_SetControlStyle(s.instance, value)
}

func (s *TScrollBox) Floating() bool {
    return ScrollBox_GetFloating(s.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (s *TScrollBox) Parent() *TWinControl {
    return AsWinControl(ScrollBox_GetParent(s.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (s *TScrollBox) SetParent(value IWinControl) {
    ScrollBox_SetParent(s.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (s *TScrollBox) Left() int32 {
    return ScrollBox_GetLeft(s.instance)
}

// 设置左边位置。
//
// Set Left position.
func (s *TScrollBox) SetLeft(value int32) {
    ScrollBox_SetLeft(s.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (s *TScrollBox) Top() int32 {
    return ScrollBox_GetTop(s.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (s *TScrollBox) SetTop(value int32) {
    ScrollBox_SetTop(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TScrollBox) Width() int32 {
    return ScrollBox_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TScrollBox) SetWidth(value int32) {
    ScrollBox_SetWidth(s.instance, value)
}

// 获取高度。
//
// Get height.
func (s *TScrollBox) Height() int32 {
    return ScrollBox_GetHeight(s.instance)
}

// 设置高度。
//
// Set height.
func (s *TScrollBox) SetHeight(value int32) {
    ScrollBox_SetHeight(s.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (s *TScrollBox) Cursor() TCursor {
    return ScrollBox_GetCursor(s.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (s *TScrollBox) SetCursor(value TCursor) {
    ScrollBox_SetCursor(s.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TScrollBox) Hint() string {
    return ScrollBox_GetHint(s.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TScrollBox) SetHint(value string) {
    ScrollBox_SetHint(s.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TScrollBox) ComponentCount() int32 {
    return ScrollBox_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TScrollBox) ComponentIndex() int32 {
    return ScrollBox_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TScrollBox) SetComponentIndex(value int32) {
    ScrollBox_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TScrollBox) Owner() *TComponent {
    return AsComponent(ScrollBox_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TScrollBox) Name() string {
    return ScrollBox_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TScrollBox) SetName(value string) {
    ScrollBox_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TScrollBox) Tag() int {
    return ScrollBox_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TScrollBox) SetTag(value int) {
    ScrollBox_SetTag(s.instance, value)
}

// 获取左边锚点。
func (s *TScrollBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ScrollBox_GetAnchorSideLeft(s.instance))
}

// 设置左边锚点。
func (s *TScrollBox) SetAnchorSideLeft(value *TAnchorSide) {
    ScrollBox_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (s *TScrollBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ScrollBox_GetAnchorSideTop(s.instance))
}

// 设置顶边锚点。
func (s *TScrollBox) SetAnchorSideTop(value *TAnchorSide) {
    ScrollBox_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// 获取右边锚点。
func (s *TScrollBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ScrollBox_GetAnchorSideRight(s.instance))
}

// 设置右边锚点。
func (s *TScrollBox) SetAnchorSideRight(value *TAnchorSide) {
    ScrollBox_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// 获取底边锚点。
func (s *TScrollBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ScrollBox_GetAnchorSideBottom(s.instance))
}

// 设置底边锚点。
func (s *TScrollBox) SetAnchorSideBottom(value *TAnchorSide) {
    ScrollBox_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ScrollBox_GetChildSizing(s.instance))
}

func (s *TScrollBox) SetChildSizing(value *TControlChildSizing) {
    ScrollBox_SetChildSizing(s.instance, CheckPtr(value))
}

// 获取边框间距。
func (s *TScrollBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ScrollBox_GetBorderSpacing(s.instance))
}

// 设置边框间距。
func (s *TScrollBox) SetBorderSpacing(value *TControlBorderSpacing) {
    ScrollBox_SetBorderSpacing(s.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (s *TScrollBox) DockClients(Index int32) *TControl {
    return AsControl(ScrollBox_GetDockClients(s.instance, Index))
}

// 获取指定索引子控件。
func (s *TScrollBox) Controls(Index int32) *TControl {
    return AsControl(ScrollBox_GetControls(s.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TScrollBox) Components(AIndex int32) *TComponent {
    return AsComponent(ScrollBox_GetComponents(s.instance, AIndex))
}

// 获取锚侧面。
func (s *TScrollBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ScrollBox_GetAnchorSide(s.instance, AKind))
}

