
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

type TControl struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewControl(owner IComponent) *TControl {
    c := new(TControl)
    c.instance = Control_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(c, (*TControl).Free)
    return c
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsControl(obj interface{}) *TControl {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TControl{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsControl.
func ControlFromInst(inst uintptr) *TControl {
    return AsControl(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsControl.
func ControlFromObj(obj IObject) *TControl {
    return AsControl(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsControl.
func ControlFromUnsafePointer(ptr unsafe.Pointer) *TControl {
    return AsControl(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (c *TControl) Free() {
    if c.instance != 0 {
        Control_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TControl) Instance() uintptr {
    return c.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TControl) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TControl) IsValid() bool {
    return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TControl) Is() TIs {
    return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TControl) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TControlClass() TClass {
    return Control_StaticClassType()
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TControl) BringToFront() {
    Control_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TControl) ClientToScreen(Point TPoint) TPoint {
    return Control_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TControl) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Control_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TControl) Dragging() bool {
    return Control_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TControl) HasParent() bool {
    return Control_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TControl) Hide() {
    Control_Hide(c.instance)
}

// 要求重绘。
//
// Redraw.
func (c *TControl) Invalidate() {
    Control_Invalidate(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Control_Perform(c.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TControl) Refresh() {
    Control_Refresh(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TControl) Repaint() {
    Control_Repaint(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TControl) ScreenToClient(Point TPoint) TPoint {
    return Control_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TControl) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Control_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TControl) SendToBack() {
    Control_SendToBack(c.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Control_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// 显示控件。
//
// Show control.
func (c *TControl) Show() {
    Control_Show(c.instance)
}

// 控件更新。
//
// Update.
func (c *TControl) Update() {
    Control_Update(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TControl) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Control_GetTextBuf(c.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TControl) GetTextLen() int32 {
    return Control_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TControl) SetTextBuf(Buffer string) {
    Control_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TControl) FindComponent(AName string) *TComponent {
    return AsComponent(Control_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TControl) GetNamePath() string {
    return Control_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TControl) Assign(Source IObject) {
    Control_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TControl) ClassType() TClass {
    return Control_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TControl) ClassName() string {
    return Control_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TControl) InstanceSize() int32 {
    return Control_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TControl) InheritsFrom(AClass TClass) bool {
    return Control_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TControl) Equals(Obj IObject) bool {
    return Control_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TControl) GetHashCode() int32 {
    return Control_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TControl) ToString() string {
    return Control_ToString(c.instance)
}

func (c *TControl) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Control_AnchorToNeighbour(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TControl) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Control_AnchorParallel(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TControl) AnchorHorizontalCenterTo(ASibling IControl) {
    Control_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TControl) AnchorVerticalCenterTo(ASibling IControl) {
    Control_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TControl) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Control_AnchorAsAlign(c.instance, ATheAlign , ASpace)
}

func (c *TControl) AnchorClient(ASpace int32) {
    Control_AnchorClient(c.instance, ASpace)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TControl) Enabled() bool {
    return Control_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TControl) SetEnabled(value bool) {
    Control_SetEnabled(c.instance, value)
}

func (c *TControl) Action() *TAction {
    return AsAction(Control_GetAction(c.instance))
}

func (c *TControl) SetAction(value IComponent) {
    Control_SetAction(c.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TControl) Align() TAlign {
    return Control_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TControl) SetAlign(value TAlign) {
    Control_SetAlign(c.instance, value)
}

// 获取四个角位置的锚点。
func (c *TControl) Anchors() TAnchors {
    return Control_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TControl) SetAnchors(value TAnchors) {
    Control_SetAnchors(c.instance, value)
}

func (c *TControl) BiDiMode() TBiDiMode {
    return Control_GetBiDiMode(c.instance)
}

func (c *TControl) SetBiDiMode(value TBiDiMode) {
    Control_SetBiDiMode(c.instance, value)
}

func (c *TControl) BoundsRect() TRect {
    return Control_GetBoundsRect(c.instance)
}

func (c *TControl) SetBoundsRect(value TRect) {
    Control_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TControl) ClientHeight() int32 {
    return Control_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TControl) SetClientHeight(value int32) {
    Control_SetClientHeight(c.instance, value)
}

func (c *TControl) ClientOrigin() TPoint {
    return Control_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TControl) ClientRect() TRect {
    return Control_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TControl) ClientWidth() int32 {
    return Control_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TControl) SetClientWidth(value int32) {
    Control_SetClientWidth(c.instance, value)
}

// 获取约束控件大小。
func (c *TControl) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Control_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TControl) SetConstraints(value *TSizeConstraints) {
    Control_SetConstraints(c.instance, CheckPtr(value))
}

// 获取控件状态。
//
// Get control state.
func (c *TControl) ControlState() TControlState {
    return Control_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TControl) SetControlState(value TControlState) {
    Control_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TControl) ControlStyle() TControlStyle {
    return Control_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TControl) SetControlStyle(value TControlStyle) {
    Control_SetControlStyle(c.instance, value)
}

func (c *TControl) Floating() bool {
    return Control_GetFloating(c.instance)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TControl) ShowHint() bool {
    return Control_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TControl) SetShowHint(value bool) {
    Control_SetShowHint(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TControl) Visible() bool {
    return Control_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TControl) SetVisible(value bool) {
    Control_SetVisible(c.instance, value)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TControl) Parent() *TWinControl {
    return AsWinControl(Control_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TControl) SetParent(value IWinControl) {
    Control_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TControl) Left() int32 {
    return Control_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TControl) SetLeft(value int32) {
    Control_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TControl) Top() int32 {
    return Control_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TControl) SetTop(value int32) {
    Control_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TControl) Width() int32 {
    return Control_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TControl) SetWidth(value int32) {
    Control_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TControl) Height() int32 {
    return Control_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TControl) SetHeight(value int32) {
    Control_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TControl) Cursor() TCursor {
    return Control_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TControl) SetCursor(value TCursor) {
    Control_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TControl) Hint() string {
    return Control_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TControl) SetHint(value string) {
    Control_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TControl) ComponentCount() int32 {
    return Control_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TControl) ComponentIndex() int32 {
    return Control_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TControl) SetComponentIndex(value int32) {
    Control_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TControl) Owner() *TComponent {
    return AsComponent(Control_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TControl) Name() string {
    return Control_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TControl) SetName(value string) {
    Control_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TControl) Tag() int {
    return Control_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TControl) SetTag(value int) {
    Control_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TControl) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Control_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TControl) SetAnchorSideLeft(value *TAnchorSide) {
    Control_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TControl) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Control_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TControl) SetAnchorSideTop(value *TAnchorSide) {
    Control_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TControl) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Control_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TControl) SetAnchorSideRight(value *TAnchorSide) {
    Control_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TControl) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Control_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TControl) SetAnchorSideBottom(value *TAnchorSide) {
    Control_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TControl) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Control_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TControl) SetBorderSpacing(value *TControlBorderSpacing) {
    Control_SetBorderSpacing(c.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TControl) Components(AIndex int32) *TComponent {
    return AsComponent(Control_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TControl) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Control_GetAnchorSide(c.instance, AKind))
}

