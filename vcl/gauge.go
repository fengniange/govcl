
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

type TGauge struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewGauge(owner IComponent) *TGauge {
    g := new(TGauge)
    g.instance = Gauge_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(g, (*TGauge).Free)
    return g
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsGauge(obj interface{}) *TGauge {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TGauge{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsGauge.
func GaugeFromInst(inst uintptr) *TGauge {
    return AsGauge(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsGauge.
func GaugeFromObj(obj IObject) *TGauge {
    return AsGauge(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGauge.
func GaugeFromUnsafePointer(ptr unsafe.Pointer) *TGauge {
    return AsGauge(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (g *TGauge) Free() {
    if g.instance != 0 {
        Gauge_Free(g.instance)
        g.instance, g.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (g *TGauge) Instance() uintptr {
    return g.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (g *TGauge) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (g *TGauge) IsValid() bool {
    return g.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (g *TGauge) Is() TIs {
    return TIs(g.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (g *TGauge) As() TAs {
//    return TAs(g.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TGaugeClass() TClass {
    return Gauge_StaticClassType()
}

func (g *TGauge) AddProgress(Value int32) {
    Gauge_AddProgress(g.instance, Value)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (g *TGauge) BringToFront() {
    Gauge_BringToFront(g.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (g *TGauge) ClientToScreen(Point TPoint) TPoint {
    return Gauge_ClientToScreen(g.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (g *TGauge) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (g *TGauge) Dragging() bool {
    return Gauge_Dragging(g.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (g *TGauge) HasParent() bool {
    return Gauge_HasParent(g.instance)
}

// 隐藏控件。
//
// Hidden control.
func (g *TGauge) Hide() {
    Gauge_Hide(g.instance)
}

// 要求重绘。
//
// Redraw.
func (g *TGauge) Invalidate() {
    Gauge_Invalidate(g.instance)
}

// 发送一个消息。
//
// Send a message.
func (g *TGauge) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Gauge_Perform(g.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (g *TGauge) Refresh() {
    Gauge_Refresh(g.instance)
}

// 重绘。
//
// Repaint.
func (g *TGauge) Repaint() {
    Gauge_Repaint(g.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (g *TGauge) ScreenToClient(Point TPoint) TPoint {
    return Gauge_ScreenToClient(g.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (g *TGauge) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Gauge_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (g *TGauge) SendToBack() {
    Gauge_SendToBack(g.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (g *TGauge) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Gauge_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// 显示控件。
//
// Show control.
func (g *TGauge) Show() {
    Gauge_Show(g.instance)
}

// 控件更新。
//
// Update.
func (g *TGauge) Update() {
    Gauge_Update(g.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (g *TGauge) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return Gauge_GetTextBuf(g.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (g *TGauge) GetTextLen() int32 {
    return Gauge_GetTextLen(g.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (g *TGauge) SetTextBuf(Buffer string) {
    Gauge_SetTextBuf(g.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (g *TGauge) FindComponent(AName string) *TComponent {
    return AsComponent(Gauge_FindComponent(g.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (g *TGauge) GetNamePath() string {
    return Gauge_GetNamePath(g.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (g *TGauge) Assign(Source IObject) {
    Gauge_Assign(g.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (g *TGauge) ClassType() TClass {
    return Gauge_ClassType(g.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (g *TGauge) ClassName() string {
    return Gauge_ClassName(g.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (g *TGauge) InstanceSize() int32 {
    return Gauge_InstanceSize(g.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (g *TGauge) InheritsFrom(AClass TClass) bool {
    return Gauge_InheritsFrom(g.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (g *TGauge) Equals(Obj IObject) bool {
    return Gauge_Equals(g.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (g *TGauge) GetHashCode() int32 {
    return Gauge_GetHashCode(g.instance)
}

// 文本类信息。
//
// Text information.
func (g *TGauge) ToString() string {
    return Gauge_ToString(g.instance)
}

func (g *TGauge) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorToNeighbour(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (g *TGauge) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    Gauge_AnchorParallel(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (g *TGauge) AnchorHorizontalCenterTo(ASibling IControl) {
    Gauge_AnchorHorizontalCenterTo(g.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (g *TGauge) AnchorVerticalCenterTo(ASibling IControl) {
    Gauge_AnchorVerticalCenterTo(g.instance, CheckPtr(ASibling))
}

func (g *TGauge) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    Gauge_AnchorAsAlign(g.instance, ATheAlign , ASpace)
}

func (g *TGauge) AnchorClient(ASpace int32) {
    Gauge_AnchorClient(g.instance, ASpace)
}

func (g *TGauge) PercentDone() int32 {
    return Gauge_GetPercentDone(g.instance)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (g *TGauge) Align() TAlign {
    return Gauge_GetAlign(g.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (g *TGauge) SetAlign(value TAlign) {
    Gauge_SetAlign(g.instance, value)
}

// 获取四个角位置的锚点。
func (g *TGauge) Anchors() TAnchors {
    return Gauge_GetAnchors(g.instance)
}

// 设置四个角位置的锚点。
func (g *TGauge) SetAnchors(value TAnchors) {
    Gauge_SetAnchors(g.instance, value)
}

func (g *TGauge) BackColor() TColor {
    return Gauge_GetBackColor(g.instance)
}

func (g *TGauge) SetBackColor(value TColor) {
    Gauge_SetBackColor(g.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (g *TGauge) BorderStyle() TBorderStyle {
    return Gauge_GetBorderStyle(g.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (g *TGauge) SetBorderStyle(value TBorderStyle) {
    Gauge_SetBorderStyle(g.instance, value)
}

// 获取颜色。
//
// Get color.
func (g *TGauge) Color() TColor {
    return Gauge_GetColor(g.instance)
}

// 设置颜色。
//
// Set color.
func (g *TGauge) SetColor(value TColor) {
    Gauge_SetColor(g.instance, value)
}

// 获取约束控件大小。
func (g *TGauge) Constraints() *TSizeConstraints {
    return AsSizeConstraints(Gauge_GetConstraints(g.instance))
}

// 设置约束控件大小。
func (g *TGauge) SetConstraints(value *TSizeConstraints) {
    Gauge_SetConstraints(g.instance, CheckPtr(value))
}

// 获取控件启用。
//
// Get the control enabled.
func (g *TGauge) Enabled() bool {
    return Gauge_GetEnabled(g.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (g *TGauge) SetEnabled(value bool) {
    Gauge_SetEnabled(g.instance, value)
}

func (g *TGauge) ForeColor() TColor {
    return Gauge_GetForeColor(g.instance)
}

func (g *TGauge) SetForeColor(value TColor) {
    Gauge_SetForeColor(g.instance, value)
}

// 获取字体。
//
// Get Font.
func (g *TGauge) Font() *TFont {
    return AsFont(Gauge_GetFont(g.instance))
}

// 设置字体。
//
// Set Font.
func (g *TGauge) SetFont(value *TFont) {
    Gauge_SetFont(g.instance, CheckPtr(value))
}

func (g *TGauge) MinValue() int32 {
    return Gauge_GetMinValue(g.instance)
}

func (g *TGauge) SetMinValue(value int32) {
    Gauge_SetMinValue(g.instance, value)
}

func (g *TGauge) MaxValue() int32 {
    return Gauge_GetMaxValue(g.instance)
}

func (g *TGauge) SetMaxValue(value int32) {
    Gauge_SetMaxValue(g.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (g *TGauge) ParentColor() bool {
    return Gauge_GetParentColor(g.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (g *TGauge) SetParentColor(value bool) {
    Gauge_SetParentColor(g.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (g *TGauge) ParentFont() bool {
    return Gauge_GetParentFont(g.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (g *TGauge) SetParentFont(value bool) {
    Gauge_SetParentFont(g.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (g *TGauge) ParentShowHint() bool {
    return Gauge_GetParentShowHint(g.instance)
}

// 设置以父容器的ShowHint属性为准。
func (g *TGauge) SetParentShowHint(value bool) {
    Gauge_SetParentShowHint(g.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (g *TGauge) PopupMenu() *TPopupMenu {
    return AsPopupMenu(Gauge_GetPopupMenu(g.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (g *TGauge) SetPopupMenu(value IComponent) {
    Gauge_SetPopupMenu(g.instance, CheckPtr(value))
}

func (g *TGauge) Progress() int32 {
    return Gauge_GetProgress(g.instance)
}

func (g *TGauge) SetProgress(value int32) {
    Gauge_SetProgress(g.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (g *TGauge) ShowHint() bool {
    return Gauge_GetShowHint(g.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (g *TGauge) SetShowHint(value bool) {
    Gauge_SetShowHint(g.instance, value)
}

func (g *TGauge) ShowText() bool {
    return Gauge_GetShowText(g.instance)
}

func (g *TGauge) SetShowText(value bool) {
    Gauge_SetShowText(g.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (g *TGauge) Visible() bool {
    return Gauge_GetVisible(g.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (g *TGauge) SetVisible(value bool) {
    Gauge_SetVisible(g.instance, value)
}

func (g *TGauge) Action() *TAction {
    return AsAction(Gauge_GetAction(g.instance))
}

func (g *TGauge) SetAction(value IComponent) {
    Gauge_SetAction(g.instance, CheckPtr(value))
}

func (g *TGauge) BiDiMode() TBiDiMode {
    return Gauge_GetBiDiMode(g.instance)
}

func (g *TGauge) SetBiDiMode(value TBiDiMode) {
    Gauge_SetBiDiMode(g.instance, value)
}

func (g *TGauge) BoundsRect() TRect {
    return Gauge_GetBoundsRect(g.instance)
}

func (g *TGauge) SetBoundsRect(value TRect) {
    Gauge_SetBoundsRect(g.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (g *TGauge) ClientHeight() int32 {
    return Gauge_GetClientHeight(g.instance)
}

// 设置客户区高度。
//
// Set client height.
func (g *TGauge) SetClientHeight(value int32) {
    Gauge_SetClientHeight(g.instance, value)
}

func (g *TGauge) ClientOrigin() TPoint {
    return Gauge_GetClientOrigin(g.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (g *TGauge) ClientRect() TRect {
    return Gauge_GetClientRect(g.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (g *TGauge) ClientWidth() int32 {
    return Gauge_GetClientWidth(g.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (g *TGauge) SetClientWidth(value int32) {
    Gauge_SetClientWidth(g.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (g *TGauge) ControlState() TControlState {
    return Gauge_GetControlState(g.instance)
}

// 设置控件状态。
//
// Set control state.
func (g *TGauge) SetControlState(value TControlState) {
    Gauge_SetControlState(g.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (g *TGauge) ControlStyle() TControlStyle {
    return Gauge_GetControlStyle(g.instance)
}

// 设置控件样式。
//
// Set control style.
func (g *TGauge) SetControlStyle(value TControlStyle) {
    Gauge_SetControlStyle(g.instance, value)
}

func (g *TGauge) Floating() bool {
    return Gauge_GetFloating(g.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (g *TGauge) Parent() *TWinControl {
    return AsWinControl(Gauge_GetParent(g.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (g *TGauge) SetParent(value IWinControl) {
    Gauge_SetParent(g.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (g *TGauge) Left() int32 {
    return Gauge_GetLeft(g.instance)
}

// 设置左边位置。
//
// Set Left position.
func (g *TGauge) SetLeft(value int32) {
    Gauge_SetLeft(g.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (g *TGauge) Top() int32 {
    return Gauge_GetTop(g.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (g *TGauge) SetTop(value int32) {
    Gauge_SetTop(g.instance, value)
}

// 获取宽度。
//
// Get width.
func (g *TGauge) Width() int32 {
    return Gauge_GetWidth(g.instance)
}

// 设置宽度。
//
// Set width.
func (g *TGauge) SetWidth(value int32) {
    Gauge_SetWidth(g.instance, value)
}

// 获取高度。
//
// Get height.
func (g *TGauge) Height() int32 {
    return Gauge_GetHeight(g.instance)
}

// 设置高度。
//
// Set height.
func (g *TGauge) SetHeight(value int32) {
    Gauge_SetHeight(g.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (g *TGauge) Cursor() TCursor {
    return Gauge_GetCursor(g.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (g *TGauge) SetCursor(value TCursor) {
    Gauge_SetCursor(g.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (g *TGauge) Hint() string {
    return Gauge_GetHint(g.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (g *TGauge) SetHint(value string) {
    Gauge_SetHint(g.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (g *TGauge) ComponentCount() int32 {
    return Gauge_GetComponentCount(g.instance)
}

// 获取组件索引。
//
// Get component index.
func (g *TGauge) ComponentIndex() int32 {
    return Gauge_GetComponentIndex(g.instance)
}

// 设置组件索引。
//
// Set component index.
func (g *TGauge) SetComponentIndex(value int32) {
    Gauge_SetComponentIndex(g.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (g *TGauge) Owner() *TComponent {
    return AsComponent(Gauge_GetOwner(g.instance))
}

// 获取组件名称。
//
// Get the component name.
func (g *TGauge) Name() string {
    return Gauge_GetName(g.instance)
}

// 设置组件名称。
//
// Set the component name.
func (g *TGauge) SetName(value string) {
    Gauge_SetName(g.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (g *TGauge) Tag() int {
    return Gauge_GetTag(g.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (g *TGauge) SetTag(value int) {
    Gauge_SetTag(g.instance, value)
}

// 获取左边锚点。
func (g *TGauge) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideLeft(g.instance))
}

// 设置左边锚点。
func (g *TGauge) SetAnchorSideLeft(value *TAnchorSide) {
    Gauge_SetAnchorSideLeft(g.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (g *TGauge) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideTop(g.instance))
}

// 设置顶边锚点。
func (g *TGauge) SetAnchorSideTop(value *TAnchorSide) {
    Gauge_SetAnchorSideTop(g.instance, CheckPtr(value))
}

// 获取右边锚点。
func (g *TGauge) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideRight(g.instance))
}

// 设置右边锚点。
func (g *TGauge) SetAnchorSideRight(value *TAnchorSide) {
    Gauge_SetAnchorSideRight(g.instance, CheckPtr(value))
}

// 获取底边锚点。
func (g *TGauge) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSideBottom(g.instance))
}

// 设置底边锚点。
func (g *TGauge) SetAnchorSideBottom(value *TAnchorSide) {
    Gauge_SetAnchorSideBottom(g.instance, CheckPtr(value))
}

// 获取边框间距。
func (g *TGauge) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(Gauge_GetBorderSpacing(g.instance))
}

// 设置边框间距。
func (g *TGauge) SetBorderSpacing(value *TControlBorderSpacing) {
    Gauge_SetBorderSpacing(g.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (g *TGauge) Components(AIndex int32) *TComponent {
    return AsComponent(Gauge_GetComponents(g.instance, AIndex))
}

// 获取锚侧面。
func (g *TGauge) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(Gauge_GetAnchorSide(g.instance, AKind))
}

