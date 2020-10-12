package bytecode

import (
	"encoding/binary"
	"fmt"
	"math"
)

/*
Default Java major versions
Currently there's no stack frame calculation, therefore Java5 should be used at default.
*/
const Java5 = 49
const Java6 = 50
const Java7 = 51
const Java8 = 52
const Java9 = 53
const Java10 = 54
const Java11 = 55
const Java12 = 56
const Java13 = 57
const Java14 = 58
const Java15 = 59

/*
Access modifiers
*/
const AccPublic = 0x0001
const AccFinal = 0x0010
const AccSuper = 0x0020
const AccInterface = 0x0200
const AccAbstract = 0x0400
const AccSynthetic = 0x1000
const AccAnnotation = 0x2000
const AccEnum = 0x4000

const AccPrivate = 0x0002
const AccProtected = 0x0004
const AccStatic = 0x0008
const AccVolatile = 0x0040
const AccTransient = 0x0080

const AccSynchronized = 0x0020
const AccBridge = 0x0040
const AccVarargs = 0x0080
const AccNative = 0x0100
const AccStrict = 0x0800

/*
Constant pool tags
*/
const Class = 7
const Fieldref = 9
const Methodref = 10
const InterfaceMethodref = 11
const String = 8
const Integer = 3
const Float = 4
const Long = 5
const Double = 6
const NameAndType = 12
const Utf8 = 1
const MethodHandle = 15
const MethodType = 16
const InvokeDynamic = 18 // Java 7+ only

/*
Instructions
*/
const Aaload = 0x32
const Aastore = 0x53
const AconstNull = 0x01
const Aload = 0x19
const Aload0 = 0x2a
const Aload1 = 0x2b
const Aload2 = 0x2c
const Aload3 = 0x2d
const Anewarray = 0xbd
const Areturn = 0xb0
const Arraylength = 0xbe
const Astore = 0x3a
const Astore0 = 0x4b
const Astore1 = 0x4c
const Astore2 = 0x4d
const Astore3 = 0x4e
const Athrow = 0xbf
const Baload = 0x33
const Bastore = 0x54
const Bipush = 0x10
const Breakpoint = 0xca
const Caload = 0x34
const Castore = 0x55
const Checkcast = 0xc0
const D2f = 0x90
const D2i = 0x8e
const D2l = 0x8f
const Dadd = 0x64
const Daload = 0x31
const Dastore = 0x52
const Dcmpg = 0x98
const Dcmpl = 0x97
const Dconst0 = 0x0e
const Dconst1 = 0x0f
const Ddiv = 0x6f
const Dload = 0x18
const Dload0 = 0x26
const Dload1 = 0x27
const Dload2 = 0x28
const Dload3 = 0x29
const Dmul = 0x6b
const Dneg = 0x77
const Drem = 0x73
const Dreturn = 0xaf
const Dstore = 0x39
const Dstore0 = 0x47
const Dstore1 = 0x48
const Dstore2 = 0x49
const Dstore3 = 0x4a
const Dsub = 0x67
const Dup = 0x59
const Dupx1 = 0x5a
const Dupx2 = 0x5b
const Dup2 = 0x5c
const Dup2x1 = 0x5d
const Dup2x2 = 0x5e
const F2d = 0x8d
const F2i = 0x8b
const F2l = 0x8c
const Fadd = 0x62
const Faload = 0x30
const Fastore = 0x51
const Fcmpg = 0x96
const Fcmpl = 0x95
const Fconst0 = 0x0b
const Fconst1 = 0x0c
const Fconst2 = 0x0d
const Fdiv = 0x6e
const Fload = 0x17
const Fload0 = 0x22
const Fload1 = 0x23
const Fload2 = 0x24
const Fload3 = 0x25
const Fmul = 0x6a
const Fneg = 0x76
const Frem = 0x72
const Freturn = 0xae
const Fstore = 0x38
const Fstore0 = 0x43
const Fstore1 = 0x44
const Fstore2 = 0x45
const Fstore3 = 0x46
const Fsub = 0x66
const Getfield = 0xb4
const Getstatic = 0xb2
const Goto = 0xa7
const Gotow = 0xc8
const I2b = 0x91
const I2c = 0x92
const I2d = 0x87
const I2f = 0x86
const I2l = 0x85
const I2s = 0x93
const Iadd = 0x60
const Iaload = 0x2e
const Iand = 0x7e
const Iastore = 0x4f
const Iconstm1 = 0x02
const Iconst0 = 0x03
const Iconst1 = 0x04
const Iconst2 = 0x05
const Iconst3 = 0x06
const Iconst4 = 0x07
const Iconst5 = 0x08
const Idiv = 0x6c
const Ifacmpeq = 0xa5
const Ifacmpne = 0xa6
const Ificmpeq = 0x9f
const Ificmpge = 0xa2
const Ificmpgt = 0xa3
const Ificmple = 0xa4
const Ificmplt = 0xa1
const Ificmpne = 0xa0
const Ifeq = 0x99
const Ifge = 0x9c
const Ifgt = 0x9d
const Ifle = 0x9e
const Iflt = 0x9b
const Ifne = 0x9a
const Ifnonnull = 0xc7
const Ifnull = 0xc6
const Iinc = 0x84
const Iload = 0x15
const Iload0 = 0x1a
const Iload1 = 0x1b
const Iload2 = 0x1c
const Iload3 = 0x1d
const Impdep1 = 0xfe
const Impdep2 = 0xff
const Imul = 0x68
const Ineg = 0x74
const Instanceof = 0xc1
const Invokedynamic = 0xba
const Invokeinterface = 0xb9
const Invokespecial = 0xb7
const Invokestatic = 0xb8
const Invokevirtual = 0xb6
const Ior = 0x80
const Irem = 0x70
const Ireturn = 0xac
const Ishl = 0x78
const Ishr = 0x7a
const Istore = 0x36
const Istore0 = 0x3b
const Istore1 = 0x3c
const Istore2 = 0x3d
const Istore3 = 0x3e
const Isub = 0x64
const Iushr = 0x7c
const Ixor = 0x82
const Jsr = 0xa8
const Jsrw = 0xc9
const L2d = 0x8a
const L2f = 0x89
const L2i = 0x88
const Ladd = 0x61
const Laload = 0x2f
const Land = 0x7f
const Lastore = 0x50
const Lcmp = 0x94
const Lconst0 = 0x09
const Lconst1 = 0x0a
const Ldc = 0x12
const Ldcw = 0x13
const Ldc2w = 0x14
const Ldiv = 0x6d
const Lload = 0x16
const Lload0 = 0x1e
const Lload1 = 0x1f
const Lload2 = 0x20
const Lload3 = 0x21
const Lmul = 0x69
const Lneg = 0x75
const Lookupswitch = 0xab
const Lor = 0x81
const Lrem = 0x71
const Lreturn = 0xad
const Lshl = 0x79
const Lshr = 0x7b
const Lstore = 0x37
const Lstore0 = 0x3f
const Lstore1 = 0x40
const Lstore2 = 0x41
const Lstore3 = 0x42
const Lsub = 0x65
const Lushr = 0x7d
const Lxor = 0x83
const Monitorenter = 0xc2
const Monitorexit = 0xc3
const Multianewarray = 0xc5
const New = 0xbb
const Newarray = 0xbc
const Nop = 0x00
const Pop = 0x57
const Pop2 = 0x58
const Putfield = 0xb5
const Putstatic = 0xb3
const Ret = 0xa9
const Return = 0xb1
const Saload = 0x35
const Sastore = 0x56
const Sipush = 0x11
const Swap = 0x5f
const Tableswitch = 0xaa
const Wide = 0xc4

/*
Java types
*/
const TypeInt = 0
const TypeByte = 1
const TypeShort = 2
const TypeDouble = 3
const TypeLong = 4
const TypeString = 5
const TypeObject = 6
const TypeBoolean = 7
const TypeFloat = 8
const TypeChar = 9

/*
Bytecode types
*/
type U1 struct {
	V int
}

func (u *U1) Get() int {
	return u.V
}

type U2 struct {
	V int
}

func (u *U2) Get() []int {
	buf := make([]int, 2)
	buf[1] = u.V & 0xff
	buf[0] = u.V >> 8
	return buf
}

type U4 struct {
	V1 U2
	V2 U2
}

func (u *U4) Get() []int {
	return append(u.V1.Get(), u.V2.Get()...)
}

type ConstantPoolData struct {
	Tag  int
	Info []int
}

type ClassVisitor struct {
	MagicNumber      U4
	MinorVersion     U2
	MajorVersion     U2
	ConstantPoolSize U2
	ConstantPool     []ConstantPoolData
	AccessFlags      U2
	ClassIndex       U2
	SuperClassIndex  U2
	InterfaceLength  U2
	// interfaces
	FieldLength U2
	// fields
	MethodLength    U2
	Methods         []*MethodVisitor
	AttributeLength U2
	// attributes
}

func NewClass(version int, name string, accessFlags int) *ClassVisitor {
	visitor := ClassVisitor{}
	visitor.MagicNumber = U4{U2{0xcafe}, U2{0xbabe}}
	visitor.MinorVersion = U2{0x0000}
	visitor.MajorVersion = U2{version}
	visitor.ConstantPoolSize = U2{0x0001}

	visitor.AddConstantPoolData(Class, 0x0000, 0x0002)
	visitor.AddUtf8Data(name)

	visitor.ClassIndex = U2{0x0001}
	visitor.SuperClassIndex = U2{0x0000}
	visitor.AccessFlags = U2{accessFlags}

	return &visitor
}

func (v *ClassVisitor) AddUtf8Data(text string) int {
	var bytes []int
	ints := StringToInts(text)
	bytes = append(bytes, 0x00, len(text))
	bytes = append(bytes, ints...)
	return v.AddConstantPoolData(Utf8, bytes...)
}

func (v *ClassVisitor) AsBytecode() []byte {
	var buf []int
	buf = append(buf, v.MagicNumber.Get()...)
	buf = append(buf, v.MinorVersion.Get()...)
	buf = append(buf, v.MajorVersion.Get()...)
	buf = append(buf, v.ConstantPoolSize.Get()...)

	for _, pool := range v.ConstantPool {
		buf = append(buf, pool.Tag)
		buf = append(buf, pool.Info...)
	}

	buf = append(buf, v.AccessFlags.Get()...)
	buf = append(buf, v.ClassIndex.Get()...)
	buf = append(buf, v.SuperClassIndex.Get()...)

	buf = append(buf, v.InterfaceLength.Get()...)
	buf = append(buf, v.FieldLength.Get()...)

	buf = append(buf, v.MethodLength.Get()...)
	for _, method := range v.Methods {
		buf = append(buf, method.AccessFlags.Get()...)
		buf = append(buf, method.MethodName.Get()...)
		buf = append(buf, method.MethodType.Get()...)
		buf = append(buf, method.AttributeLength.Get()...)

		for _, attr := range method.Attributes {
			buf = append(buf, attr.AttributeNameIndex.Get()...)
			buf = append(buf, attr.AttributeLength.Get()...)
			for _, info := range attr.Info {
				buf = append(buf, info)
			}
		}
	}

	buf = append(buf, v.AttributeLength.Get()...)

	bytes := make([]byte, len(buf))
	for i, b := range buf {
		bytes[i] = byte(b)
	}
	return bytes
}

func (v *ClassVisitor) NextConstantIndex() int {
	return v.ConstantPoolSize.V + 1
}

func (v *ClassVisitor) AddConstantPoolData(tag int, args ...int) int {
	v.ConstantPool = append(v.ConstantPool, ConstantPoolData{tag, args})
	v.ConstantPoolSize.V += 1
	return v.ConstantPoolSize.V - 1
}

func (v *ClassVisitor) AddConstantPoolDataNext(tag int) int {
	return v.AddConstantPoolData(tag, 0x0000, v.NextConstantIndex())
}

func (v *ClassVisitor) NewMethod(accessFlags int, name string, descriptor string) *MethodVisitor {
	nameIndex := v.AddUtf8Data(name)
	descriptorIndex := v.AddUtf8Data(descriptor)
	codeIndex := v.AddUtf8Data("Code")

	visitor := MethodVisitor{}
	visitor.ParentVisitor = v
	visitor.AccessFlags = U2{accessFlags}
	visitor.MethodName = U2{nameIndex}
	visitor.MethodType = U2{descriptorIndex}
	visitor.AttributeLength = U2{0x0001}

	codeAttr := AttributeInfo{}
	codeAttr.AttributeNameIndex = U2{codeIndex}

	visitor.Attributes = append(visitor.Attributes, &codeAttr)

	v.MethodLength.V += 1
	v.Methods = append(v.Methods, &visitor)
	return &visitor
}

type MethodVisitor struct {
	ParentVisitor *ClassVisitor

	AccessFlags     U2
	MethodName      U2
	MethodType      U2
	AttributeLength U2
	Attributes      []*AttributeInfo

	MaxStack  U2
	MaxLocals U2
	CodeSize  U4
}

func (m *MethodVisitor) MaxStackLocals(maxStack int, maxLocals int) {
	m.MaxStack = U2{maxStack}
	m.MaxLocals = U2{maxLocals}
}

func (m *MethodVisitor) AddLdc(javaType int, object interface{}) {
	var index int
	switch javaType {
	case TypeChar:
		fallthrough
	case TypeByte:
		fallthrough
	case TypeShort:
		fallthrough
	case TypeBoolean:
		fallthrough
	case TypeInt:
		index = m.ParentVisitor.AddConstantPoolData(Integer, Int32ToBinary(object.(int))...)
		break
	case TypeDouble:
		index = m.ParentVisitor.AddConstantPoolData(Double, Float64ToBinary(object.(float64))...)
		break
	case TypeFloat:
		index = m.ParentVisitor.AddConstantPoolData(Float, Float32ToBinary(object.(float32))...)
		break
	case TypeLong:
		index = m.ParentVisitor.AddConstantPoolData(Long, Int64ToBinary(object.(int64))...)
		break
	case TypeString:
		index = m.ParentVisitor.AddConstantPoolDataNext(String)
		m.ParentVisitor.AddUtf8Data(fmt.Sprintf("%v", object))
		break
	}
	m.AddInsn(Ldc, index)
}

func (m *MethodVisitor) AddMethodInsn(insn int, instance, name, descriptor string) {
	iIndex := m.ParentVisitor.AddConstantPoolDataNext(Class)
	m.ParentVisitor.AddUtf8Data(instance)

	nIndex := m.ParentVisitor.AddUtf8Data(name)
	dIndex := m.ParentVisitor.AddUtf8Data(descriptor)

	ntIndex := m.ParentVisitor.AddConstantPoolData(NameAndType,
		0x0000, nIndex, 0x0000, dIndex)

	mrIndex := m.ParentVisitor.AddConstantPoolData(Methodref,
		0x0000, iIndex, 0x0000, ntIndex)

	m.AddInsn(insn, 0x0000, mrIndex)
}

func (m *MethodVisitor) AddInsn(insn int, args ...int) {
	attr := m.Attributes[0]
	attr.Info = append(attr.Info, insn)
	attr.Info = append(attr.Info, args...)

	attr.ByteSize += 1 + len(args)
	m.CodeSize.V2.V += 1 + len(args)
}

func (m *MethodVisitor) End() {
	attr := m.Attributes[0]
	attr.AttributeLength = U4{U2{0x0000}, U2{attr.ByteSize + 12}}

	var attrInfo []int
	attrInfo = append(attrInfo, m.MaxStack.Get()...)
	attrInfo = append(attrInfo, m.MaxLocals.Get()...)

	attrInfo = append(attrInfo, m.CodeSize.Get()...)

	attrInfo = append(attrInfo, attr.Info...)
	attrInfo = append(attrInfo, 0, 0, 0, 0)

	attr.Info = attrInfo
}

type AttributeInfo struct {
	AttributeNameIndex U2
	AttributeLength    U4
	Info               []int

	ByteSize int
}

func StringToInts(text string) []int {
	bytes := []byte(text)

	return BytesToInts(bytes)
}

func Float32ToBinary(value float32) []int {
	bytes := make([]byte, 4)
	bits := math.Float32bits(value)
	binary.BigEndian.PutUint32(bytes, bits)

	return BytesToInts(bytes)
}

func Float64ToBinary(value float64) []int {
	bytes := make([]byte, 8)
	bits := math.Float64bits(value)
	binary.BigEndian.PutUint64(bytes, bits)

	return BytesToInts(bytes)
}

func Int32ToBinary(value int) []int {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))

	return BytesToInts(bytes)
}

func Int64ToBinary(value int64) []int {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint64(bytes, uint64(value))

	return BytesToInts(bytes)
}

func BytesToInts(bytes []byte) []int {
	ints := make([]int, len(bytes))
	for i, b := range bytes {
		ints[i] = int(b)
	}
	return ints
}
