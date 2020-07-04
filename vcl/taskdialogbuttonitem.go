
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

type TTaskDialogButtonItem struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewTaskDialogButtonItem(AOwner *TCollection) *TTaskDialogButtonItem {
    t := new(TTaskDialogButtonItem)
    t.instance = TaskDialogButtonItem_Create(CheckPtr(AOwner))
    t.ptr = unsafe.Pointer(t.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(t, (*TTaskDialogButtonItem).Free)
    return t
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsTaskDialogButtonItem(obj interface{}) *TTaskDialogButtonItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TTaskDialogButtonItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsTaskDialogButtonItem.
func TaskDialogButtonItemFromInst(inst uintptr) *TTaskDialogButtonItem {
    return AsTaskDialogButtonItem(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsTaskDialogButtonItem.
func TaskDialogButtonItemFromObj(obj IObject) *TTaskDialogButtonItem {
    return AsTaskDialogButtonItem(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsTaskDialogButtonItem.
func TaskDialogButtonItemFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogButtonItem {
    return AsTaskDialogButtonItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (t *TTaskDialogButtonItem) Free() {
    if t.instance != 0 {
        TaskDialogButtonItem_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (t *TTaskDialogButtonItem) Instance() uintptr {
    return t.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (t *TTaskDialogButtonItem) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (t *TTaskDialogButtonItem) IsValid() bool {
    return t.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (t *TTaskDialogButtonItem) Is() TIs {
    return TIs(t.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (t *TTaskDialogButtonItem) As() TAs {
//    return TAs(t.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TTaskDialogButtonItemClass() TClass {
    return TaskDialogButtonItem_StaticClassType()
}

// 获取类名路径。
//
// Get the class name path.
func (t *TTaskDialogButtonItem) GetNamePath() string {
    return TaskDialogButtonItem_GetNamePath(t.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (t *TTaskDialogButtonItem) Assign(Source IObject) {
    TaskDialogButtonItem_Assign(t.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (t *TTaskDialogButtonItem) ClassType() TClass {
    return TaskDialogButtonItem_ClassType(t.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (t *TTaskDialogButtonItem) ClassName() string {
    return TaskDialogButtonItem_ClassName(t.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (t *TTaskDialogButtonItem) InstanceSize() int32 {
    return TaskDialogButtonItem_InstanceSize(t.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (t *TTaskDialogButtonItem) InheritsFrom(AClass TClass) bool {
    return TaskDialogButtonItem_InheritsFrom(t.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (t *TTaskDialogButtonItem) Equals(Obj IObject) bool {
    return TaskDialogButtonItem_Equals(t.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (t *TTaskDialogButtonItem) GetHashCode() int32 {
    return TaskDialogButtonItem_GetHashCode(t.instance)
}

// 文本类信息。
//
// Text information.
func (t *TTaskDialogButtonItem) ToString() string {
    return TaskDialogButtonItem_ToString(t.instance)
}

// 获取模态对话框显示结果。
func (t *TTaskDialogButtonItem) ModalResult() TModalResult {
    return TaskDialogButtonItem_GetModalResult(t.instance)
}

// 设置模态对话框显示结果。
func (t *TTaskDialogButtonItem) SetModalResult(value TModalResult) {
    TaskDialogButtonItem_SetModalResult(t.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (t *TTaskDialogButtonItem) Caption() string {
    return TaskDialogButtonItem_GetCaption(t.instance)
}

// 设置控件标题。
//
// Set the control title.
func (t *TTaskDialogButtonItem) SetCaption(value string) {
    TaskDialogButtonItem_SetCaption(t.instance, value)
}

func (t *TTaskDialogButtonItem) Default() bool {
    return TaskDialogButtonItem_GetDefault(t.instance)
}

func (t *TTaskDialogButtonItem) SetDefault(value bool) {
    TaskDialogButtonItem_SetDefault(t.instance, value)
}

func (t *TTaskDialogButtonItem) Collection() *TCollection {
    return AsCollection(TaskDialogButtonItem_GetCollection(t.instance))
}

func (t *TTaskDialogButtonItem) SetCollection(value *TCollection) {
    TaskDialogButtonItem_SetCollection(t.instance, CheckPtr(value))
}

func (t *TTaskDialogButtonItem) Index() int32 {
    return TaskDialogButtonItem_GetIndex(t.instance)
}

func (t *TTaskDialogButtonItem) SetIndex(value int32) {
    TaskDialogButtonItem_SetIndex(t.instance, value)
}

func (t *TTaskDialogButtonItem) DisplayName() string {
    return TaskDialogButtonItem_GetDisplayName(t.instance)
}

func (t *TTaskDialogButtonItem) SetDisplayName(value string) {
    TaskDialogButtonItem_SetDisplayName(t.instance, value)
}

