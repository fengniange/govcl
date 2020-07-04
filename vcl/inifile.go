
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    "time"
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TIniFile struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewIniFile(filename string) *TIniFile {
    i := new(TIniFile)
    i.instance = IniFile_Create(GoStrToDStr(filename))
    i.ptr = unsafe.Pointer(i.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(i, (*TIniFile).Free)
    return i
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsIniFile(obj interface{}) *TIniFile {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TIniFile{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsIniFile.
func IniFileFromInst(inst uintptr) *TIniFile {
    return AsIniFile(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsIniFile.
func IniFileFromObj(obj IObject) *TIniFile {
    return AsIniFile(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsIniFile.
func IniFileFromUnsafePointer(ptr unsafe.Pointer) *TIniFile {
    return AsIniFile(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (i *TIniFile) Free() {
    if i.instance != 0 {
        IniFile_Free(i.instance)
        i.instance, i.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (i *TIniFile) Instance() uintptr {
    return i.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (i *TIniFile) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (i *TIniFile) IsValid() bool {
    return i.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (i *TIniFile) Is() TIs {
    return TIs(i.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (i *TIniFile) As() TAs {
//    return TAs(i.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TIniFileClass() TClass {
    return IniFile_StaticClassType()
}

func (i *TIniFile) ReadString(Section string, Ident string, Default string) string {
    return IniFile_ReadString(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteString(Section string, Ident string, Value string) {
    IniFile_WriteString(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadSections(Strings IObject) {
    IniFile_ReadSections(i.instance, CheckPtr(Strings))
}

func (i *TIniFile) ReadSectionValues(Section string, Strings IObject) {
    IniFile_ReadSectionValues(i.instance, Section , CheckPtr(Strings))
}

func (i *TIniFile) EraseSection(Section string) {
    IniFile_EraseSection(i.instance, Section)
}

func (i *TIniFile) DeleteKey(Section string, Ident string) {
    IniFile_DeleteKey(i.instance, Section , Ident)
}

func (i *TIniFile) UpdateFile() {
    IniFile_UpdateFile(i.instance)
}

func (i *TIniFile) SectionExists(Section string) bool {
    return IniFile_SectionExists(i.instance, Section)
}

func (i *TIniFile) ReadInteger(Section string, Ident string, Default int32) int32 {
    return IniFile_ReadInteger(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteInteger(Section string, Ident string, Value int32) {
    IniFile_WriteInteger(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadBool(Section string, Ident string, Default bool) bool {
    return IniFile_ReadBool(i.instance, Section , Ident , Default)
}

func (i *TIniFile) WriteBool(Section string, Ident string, Value bool) {
    IniFile_WriteBool(i.instance, Section , Ident , Value)
}

func (i *TIniFile) ReadDate(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDate(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadDateTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadDateTime(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadFloat(Section string, Name string, Default float64) float64 {
    return IniFile_ReadFloat(i.instance, Section , Name , Default)
}

func (i *TIniFile) ReadTime(Section string, Name string, Default time.Time) time.Time {
    return IniFile_ReadTime(i.instance, Section , Name , Default)
}

func (i *TIniFile) WriteDate(Section string, Name string, Value time.Time) {
    IniFile_WriteDate(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteDateTime(Section string, Name string, Value time.Time) {
    IniFile_WriteDateTime(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteFloat(Section string, Name string, Value float64) {
    IniFile_WriteFloat(i.instance, Section , Name , Value)
}

func (i *TIniFile) WriteTime(Section string, Name string, Value time.Time) {
    IniFile_WriteTime(i.instance, Section , Name , Value)
}

func (i *TIniFile) ValueExists(Section string, Ident string) bool {
    return IniFile_ValueExists(i.instance, Section , Ident)
}

// 获取类的类型信息。
//
// Get class type information.
func (i *TIniFile) ClassType() TClass {
    return IniFile_ClassType(i.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (i *TIniFile) ClassName() string {
    return IniFile_ClassName(i.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (i *TIniFile) InstanceSize() int32 {
    return IniFile_InstanceSize(i.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (i *TIniFile) InheritsFrom(AClass TClass) bool {
    return IniFile_InheritsFrom(i.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (i *TIniFile) Equals(Obj IObject) bool {
    return IniFile_Equals(i.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (i *TIniFile) GetHashCode() int32 {
    return IniFile_GetHashCode(i.instance)
}

// 文本类信息。
//
// Text information.
func (i *TIniFile) ToString() string {
    return IniFile_ToString(i.instance)
}

func (i *TIniFile) FileName() string {
    return IniFile_GetFileName(i.instance)
}

