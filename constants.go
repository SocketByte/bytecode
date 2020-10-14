package bytecode

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

type Instruction struct {
    Opcode          byte
    StackIntake     uint16
    StackIntakeFlag int
    StackOutput     uint16
    StackOutputFlag int
}

const FlagStackEmpty = 0xaa
const FlagStackArgs = 0xbb
const FlagStackLabel = 0xcc
const FlagStackNone = 0xdd

/*
Instructions
*/
var Aaload = Instruction{byte(0x32), 2, FlagStackNone, 1, FlagStackNone}
var Aastore = Instruction{byte(0x53), 3, FlagStackNone, 0, FlagStackNone}
var Aconstnull = Instruction{byte(0x01), 0, FlagStackNone, 1, FlagStackNone}
var Aload = Instruction{byte(0x19), 0, FlagStackNone, 1, FlagStackNone}
var Aload0 = Instruction{byte(0x2a), 0, FlagStackNone, 1, FlagStackNone}
var Aload1 = Instruction{byte(0x2b), 0, FlagStackNone, 1, FlagStackNone}
var Aload2 = Instruction{byte(0x2c), 0, FlagStackNone, 1, FlagStackNone}
var Aload3 = Instruction{byte(0x2d), 0, FlagStackNone, 1, FlagStackNone}
var Anewarray = Instruction{byte(0xbd), 1, FlagStackNone, 1, FlagStackNone}
var Areturn = Instruction{byte(0xb0), 1, FlagStackNone, 0, FlagStackEmpty}
var Arraylength = Instruction{byte(0xbe), 1, FlagStackNone, 1, FlagStackNone}
var Astore = Instruction{byte(0x3a), 1, FlagStackNone, 0, FlagStackNone}
var Astore0 = Instruction{byte(0x4b), 1, FlagStackNone, 0, FlagStackNone}
var Astore1 = Instruction{byte(0x4c), 1, FlagStackNone, 0, FlagStackNone}
var Astore2 = Instruction{byte(0x4d), 1, FlagStackNone, 0, FlagStackNone}
var Astore3 = Instruction{byte(0x4e), 1, FlagStackNone, 0, FlagStackNone}
var Athrow = Instruction{byte(0xbf), 1, FlagStackNone, 1, FlagStackEmpty}
var Baload = Instruction{byte(0x33), 2, FlagStackNone, 1, FlagStackNone}
var Bastore = Instruction{byte(0x54), 3, FlagStackNone, 0, FlagStackNone}
var Bipush = Instruction{byte(0x10), 0, FlagStackNone, 1, FlagStackNone}
var Breakpoint = Instruction{byte(0xca), 0, FlagStackNone, 0, FlagStackNone}
var Caload = Instruction{byte(0x34), 2, FlagStackNone, 1, FlagStackNone}
var Castore = Instruction{byte(0x55), 3, FlagStackNone, 0, FlagStackNone}
var Checkcast = Instruction{byte(0xc0), 1, FlagStackNone, 1, FlagStackNone}
var D2f = Instruction{byte(0x90), 1, FlagStackNone, 1, FlagStackNone}
var D2i = Instruction{byte(0x8e), 1, FlagStackNone, 1, FlagStackNone}
var D2l = Instruction{byte(0x8f), 1, FlagStackNone, 1, FlagStackNone}
var Dadd = Instruction{byte(0x63), 2, FlagStackNone, 1, FlagStackNone}
var Daload = Instruction{byte(0x31), 2, FlagStackNone, 1, FlagStackNone}
var Dastore = Instruction{byte(0x52), 3, FlagStackNone, 0, FlagStackNone}
var Dcmpg = Instruction{byte(0x98), 2, FlagStackNone, 1, FlagStackNone}
var Dcmpl = Instruction{byte(0x97), 2, FlagStackNone, 1, FlagStackNone}
var Dconst0 = Instruction{byte(0x0e), 0, FlagStackNone, 1, FlagStackNone}
var Dconst1 = Instruction{byte(0x0f), 0, FlagStackNone, 1, FlagStackNone}
var Ddiv = Instruction{byte(0x6f), 2, FlagStackNone, 1, FlagStackNone}
var Dload = Instruction{byte(0x18), 0, FlagStackNone, 1, FlagStackNone}
var Dload0 = Instruction{byte(0x26), 0, FlagStackNone, 1, FlagStackNone}
var Dload1 = Instruction{byte(0x27), 0, FlagStackNone, 1, FlagStackNone}
var Dload2 = Instruction{byte(0x28), 0, FlagStackNone, 1, FlagStackNone}
var Dload3 = Instruction{byte(0x29), 0, FlagStackNone, 1, FlagStackNone}
var Dmul = Instruction{byte(0x6b), 2, FlagStackNone, 1, FlagStackNone}
var Dneg = Instruction{byte(0x77), 1, FlagStackNone, 1, FlagStackNone}
var Drem = Instruction{byte(0x73), 2, FlagStackNone, 1, FlagStackNone}
var Dreturn = Instruction{byte(0xaf), 1, FlagStackNone, 0, FlagStackEmpty}
var Dstore = Instruction{byte(0x39), 1, FlagStackNone, 0, FlagStackNone}
var Dstore0 = Instruction{byte(0x47), 1, FlagStackNone, 0, FlagStackNone}
var Dstore1 = Instruction{byte(0x48), 1, FlagStackNone, 0, FlagStackNone}
var Dstore2 = Instruction{byte(0x49), 1, FlagStackNone, 0, FlagStackNone}
var Dstore3 = Instruction{byte(0x4a), 1, FlagStackNone, 0, FlagStackNone}
var Dsub = Instruction{byte(0x67), 2, FlagStackNone, 1, FlagStackNone}
var Dup = Instruction{byte(0x59), 1, FlagStackNone, 2, FlagStackNone}
var Dupx1 = Instruction{byte(0x5a), 2, FlagStackNone, 3, FlagStackNone}
var Dupx2 = Instruction{byte(0x5b), 3, FlagStackNone, 4, FlagStackNone}
var Dup2 = Instruction{byte(0x5c), 1, FlagStackNone, 2, FlagStackNone}
var Dup2x1 = Instruction{byte(0x5d), 2, FlagStackNone, 3, FlagStackNone}
var Dup2x2 = Instruction{byte(0x5e), 2, FlagStackNone, 3, FlagStackNone}
var F2d = Instruction{byte(0x8d), 1, FlagStackNone, 1, FlagStackNone}
var F2i = Instruction{byte(0x8b), 1, FlagStackNone, 1, FlagStackNone}
var F2l = Instruction{byte(0x8c), 1, FlagStackNone, 1, FlagStackNone}
var Fadd = Instruction{byte(0x62), 2, FlagStackNone, 1, FlagStackNone}
var Faload = Instruction{byte(0x30), 2, FlagStackNone, 1, FlagStackNone}
var Fastore = Instruction{byte(0x51), 3, FlagStackNone, 0, FlagStackNone}
var Fcmpg = Instruction{byte(0x96), 2, FlagStackNone, 1, FlagStackNone}
var Fcmpl = Instruction{byte(0x95), 2, FlagStackNone, 1, FlagStackNone}
var Fconst0 = Instruction{byte(0x0b), 0, FlagStackNone, 1, FlagStackNone}
var Fconst1 = Instruction{byte(0x0c), 0, FlagStackNone, 1, FlagStackNone}
var Fconst2 = Instruction{byte(0x0d), 0, FlagStackNone, 1, FlagStackNone}
var Fdiv = Instruction{byte(0x6e), 2, FlagStackNone, 1, FlagStackNone}
var Fload = Instruction{byte(0x17), 0, FlagStackNone, 1, FlagStackNone}
var Fload0 = Instruction{byte(0x22), 0, FlagStackNone, 1, FlagStackNone}
var Fload1 = Instruction{byte(0x23), 0, FlagStackNone, 1, FlagStackNone}
var Fload2 = Instruction{byte(0x24), 0, FlagStackNone, 1, FlagStackNone}
var Fload3 = Instruction{byte(0x25), 0, FlagStackNone, 1, FlagStackNone}
var Fmul = Instruction{byte(0x6a), 2, FlagStackNone, 1, FlagStackNone}
var Fneg = Instruction{byte(0x76), 1, FlagStackNone, 1, FlagStackNone}
var Frem = Instruction{byte(0x72), 2, FlagStackNone, 1, FlagStackNone}
var Freturn = Instruction{byte(0xae), 1, FlagStackNone, 0, FlagStackEmpty}
var Fstore = Instruction{byte(0x38), 1, FlagStackNone, 0, FlagStackNone}
var Fstore0 = Instruction{byte(0x43), 1, FlagStackNone, 0, FlagStackNone}
var Fstore1 = Instruction{byte(0x44), 1, FlagStackNone, 0, FlagStackNone}
var Fstore2 = Instruction{byte(0x45), 1, FlagStackNone, 0, FlagStackNone}
var Fstore3 = Instruction{byte(0x46), 1, FlagStackNone, 0, FlagStackNone}
var Fsub = Instruction{byte(0x66), 2, FlagStackNone, 1, FlagStackNone}
var Getfield = Instruction{byte(0xb4), 1, FlagStackNone, 1, FlagStackNone}
var Getstatic = Instruction{byte(0xb2), 0, FlagStackNone, 1, FlagStackNone}
var Goto = Instruction{byte(0xa7), 0, FlagStackLabel, 0, FlagStackNone}
var Gotow = Instruction{byte(0xc8), 0, FlagStackLabel, 0, FlagStackNone}
var I2b = Instruction{byte(0x91), 1, FlagStackNone, 1, FlagStackNone}
var I2c = Instruction{byte(0x92), 1, FlagStackNone, 1, FlagStackNone}
var I2d = Instruction{byte(0x87), 1, FlagStackNone, 1, FlagStackNone}
var I2f = Instruction{byte(0x86), 1, FlagStackNone, 1, FlagStackNone}
var I2l = Instruction{byte(0x85), 1, FlagStackNone, 1, FlagStackNone}
var I2s = Instruction{byte(0x93), 1, FlagStackNone, 1, FlagStackNone}
var Iadd = Instruction{byte(0x60), 2, FlagStackNone, 1, FlagStackNone}
var Iaload = Instruction{byte(0x2e), 2, FlagStackNone, 1, FlagStackNone}
var Iand = Instruction{byte(0x7e), 2, FlagStackNone, 1, FlagStackNone}
var Iastore = Instruction{byte(0x4f), 3, FlagStackNone, 0, FlagStackNone}
var Iconstm1 = Instruction{byte(0x02), 0, FlagStackNone, 1, FlagStackNone}
var Iconst0 = Instruction{byte(0x03), 0, FlagStackNone, 1, FlagStackNone}
var Iconst1 = Instruction{byte(0x04), 0, FlagStackNone, 1, FlagStackNone}
var Iconst2 = Instruction{byte(0x05), 0, FlagStackNone, 1, FlagStackNone}
var Iconst3 = Instruction{byte(0x06), 0, FlagStackNone, 1, FlagStackNone}
var Iconst4 = Instruction{byte(0x07), 0, FlagStackNone, 1, FlagStackNone}
var Iconst5 = Instruction{byte(0x08), 0, FlagStackNone, 1, FlagStackNone}
var Idiv = Instruction{byte(0x6c), 2, FlagStackNone, 1, FlagStackNone}
var Ifacmpeq = Instruction{byte(0xa5), 2, FlagStackLabel, 0, FlagStackNone}
var Ifacmpne = Instruction{byte(0xa6), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmpeq = Instruction{byte(0x9f), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmpge = Instruction{byte(0xa2), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmpgt = Instruction{byte(0xa3), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmple = Instruction{byte(0xa4), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmplt = Instruction{byte(0xa1), 2, FlagStackLabel, 0, FlagStackNone}
var Ificmpne = Instruction{byte(0xa0), 2, FlagStackLabel, 0, FlagStackNone}
var Ifeq = Instruction{byte(0x99), 1, FlagStackLabel, 0, FlagStackNone}
var Ifge = Instruction{byte(0x9c), 1, FlagStackLabel, 0, FlagStackNone}
var Ifgt = Instruction{byte(0x9d), 1, FlagStackLabel, 0, FlagStackNone}
var Ifle = Instruction{byte(0x9e), 1, FlagStackLabel, 0, FlagStackNone}
var Iflt = Instruction{byte(0x9b), 1, FlagStackLabel, 0, FlagStackNone}
var Ifne = Instruction{byte(0x9a), 1, FlagStackLabel, 0, FlagStackNone}
var Ifnonnull = Instruction{byte(0xc7), 1, FlagStackLabel, 0, FlagStackNone}
var Ifnull = Instruction{byte(0xc6), 1, FlagStackLabel, 0, FlagStackNone}
var Iinc = Instruction{byte(0x84), 0, FlagStackNone, 0, FlagStackNone}
var Iload = Instruction{byte(0x15), 0, FlagStackNone, 1, FlagStackNone}
var Iload0 = Instruction{byte(0x1a), 0, FlagStackNone, 1, FlagStackNone}
var Iload1 = Instruction{byte(0x1b), 0, FlagStackNone, 1, FlagStackNone}
var Iload2 = Instruction{byte(0x1c), 0, FlagStackNone, 1, FlagStackNone}
var Iload3 = Instruction{byte(0x1d), 0, FlagStackNone, 1, FlagStackNone}
var Impdep1 = Instruction{byte(0xfe), 0, FlagStackNone, 0, FlagStackNone}
var Impdep2 = Instruction{byte(0xff), 0, FlagStackNone, 0, FlagStackNone}
var Imul = Instruction{byte(0x68), 2, FlagStackNone, 1, FlagStackNone}
var Ineg = Instruction{byte(0x74), 1, FlagStackNone, 1, FlagStackNone}
var Instanceof = Instruction{byte(0xc1), 1, FlagStackNone, 1, FlagStackNone}
var Invokedynamic = Instruction{byte(0xba), 1, FlagStackArgs, 1, FlagStackNone}
var Invokeinterface = Instruction{byte(0xb9), 1, FlagStackArgs, 1, FlagStackNone}
var Invokespecial = Instruction{byte(0xb7), 1, FlagStackArgs, 1, FlagStackNone}
var Invokestatic = Instruction{byte(0xb8), 0, FlagStackArgs, 1, FlagStackNone}
var Invokevirtual = Instruction{byte(0xb6), 1, FlagStackArgs, 1, FlagStackNone}
var Ior = Instruction{byte(0x80), 2, FlagStackNone, 1, FlagStackNone}
var Irem = Instruction{byte(0x70), 2, FlagStackNone, 1, FlagStackNone}
var Ireturn = Instruction{byte(0xac), 1, FlagStackNone, 0, FlagStackEmpty}
var Ishl = Instruction{byte(0x78), 2, FlagStackNone, 1, FlagStackNone}
var Ishr = Instruction{byte(0x7a), 2, FlagStackNone, 1, FlagStackNone}
var Istore = Instruction{byte(0x36), 1, FlagStackNone, 0, FlagStackNone}
var Istore0 = Instruction{byte(0x3b), 1, FlagStackNone, 0, FlagStackNone}
var Istore1 = Instruction{byte(0x3c), 1, FlagStackNone, 0, FlagStackNone}
var Istore2 = Instruction{byte(0x3d), 1, FlagStackNone, 0, FlagStackNone}
var Istore3 = Instruction{byte(0x3e), 1, FlagStackNone, 0, FlagStackNone}
var Isub = Instruction{byte(0x64), 2, FlagStackNone, 1, FlagStackNone}
var Iushr = Instruction{byte(0x7c), 2, FlagStackNone, 1, FlagStackNone}
var Ixor = Instruction{byte(0x82), 2, FlagStackNone, 1, FlagStackNone}
var Jsr = Instruction{byte(0xa8), 0, FlagStackLabel, 1, FlagStackNone}
var Jsrw = Instruction{byte(0xc9), 0, FlagStackLabel, 1, FlagStackNone}
var L2d = Instruction{byte(0x8a), 1, FlagStackNone, 1, FlagStackNone}
var L2f = Instruction{byte(0x89), 1, FlagStackNone, 1, FlagStackNone}
var L2i = Instruction{byte(0x88), 1, FlagStackNone, 1, FlagStackNone}
var Ladd = Instruction{byte(0x61), 2, FlagStackNone, 1, FlagStackNone}
var Laload = Instruction{byte(0x2f), 2, FlagStackNone, 1, FlagStackNone}
var Land = Instruction{byte(0x7f), 2, FlagStackNone, 1, FlagStackNone}
var Lastore = Instruction{byte(0x50), 3, FlagStackNone, 0, FlagStackNone}
var Lcmp = Instruction{byte(0x94), 2, FlagStackNone, 1, FlagStackNone}
var Lconst0 = Instruction{byte(0x09), 0, FlagStackNone, 1, FlagStackNone}
var Lconst1 = Instruction{byte(0x0a), 0, FlagStackNone, 1, FlagStackNone}
var Ldc = Instruction{byte(0x12), 0, FlagStackNone, 1, FlagStackNone}
var Ldcw = Instruction{byte(0x13), 0, FlagStackNone, 1, FlagStackNone}
var Ldc2w = Instruction{byte(0x14), 0, FlagStackNone, 1, FlagStackNone}
var Ldiv = Instruction{byte(0x6d), 2, FlagStackNone, 1, FlagStackNone}
var Lload = Instruction{byte(0x16), 0, FlagStackNone, 1, FlagStackNone}
var Lload0 = Instruction{byte(0x1e), 0, FlagStackNone, 1, FlagStackNone}
var Lload1 = Instruction{byte(0x1f), 0, FlagStackNone, 1, FlagStackNone}
var Lload2 = Instruction{byte(0x20), 0, FlagStackNone, 1, FlagStackNone}
var Lload3 = Instruction{byte(0x21), 0, FlagStackNone, 1, FlagStackNone}
var Lmul = Instruction{byte(0x69), 2, FlagStackNone, 1, FlagStackNone}
var Lneg = Instruction{byte(0x75), 1, FlagStackNone, 1, FlagStackNone}
var Lookupswitch = Instruction{byte(0xab), 1, FlagStackNone, 0, FlagStackNone}
var Lor = Instruction{byte(0x81), 2, FlagStackNone, 1, FlagStackNone}
var Lrem = Instruction{byte(0x71), 2, FlagStackNone, 1, FlagStackNone}
var Lreturn = Instruction{byte(0xad), 1, FlagStackNone, 0, FlagStackEmpty}
var Lshl = Instruction{byte(0x79), 2, FlagStackNone, 1, FlagStackNone}
var Lshr = Instruction{byte(0x7b), 2, FlagStackNone, 1, FlagStackNone}
var Lstore = Instruction{byte(0x37), 1, FlagStackNone, 0, FlagStackNone}
var Lstore0 = Instruction{byte(0x3f), 1, FlagStackNone, 0, FlagStackNone}
var Lstore1 = Instruction{byte(0x40), 1, FlagStackNone, 0, FlagStackNone}
var Lstore2 = Instruction{byte(0x41), 1, FlagStackNone, 0, FlagStackNone}
var Lstore3 = Instruction{byte(0x42), 1, FlagStackNone, 0, FlagStackNone}
var Lsub = Instruction{byte(0x65), 2, FlagStackNone, 1, FlagStackNone}
var Lushr = Instruction{byte(0x7d), 2, FlagStackNone, 1, FlagStackNone}
var Lxor = Instruction{byte(0x83), 2, FlagStackNone, 1, FlagStackNone}
var Monitorenter = Instruction{byte(0xc2), 1, FlagStackNone, 0, FlagStackNone}
var Monitorexit = Instruction{byte(0xc3), 1, FlagStackNone, 0, FlagStackNone}

//var Multianewarray = Instruction{byte(0xc5), 1, FlagStackArgs, 1, FlagStackNone}
var New = Instruction{byte(0xbb), 0, FlagStackNone, 1, FlagStackNone}
var Newarray = Instruction{byte(0xbc), 1, FlagStackNone, 1, FlagStackNone}
var Nop = Instruction{byte(0x00), 0, FlagStackNone, 0, FlagStackNone}
var Pop = Instruction{byte(0x57), 1, FlagStackNone, 0, FlagStackNone}
var Pop2 = Instruction{byte(0x58), 1, FlagStackNone, 0, FlagStackNone}
var Putfield = Instruction{byte(0xb5), 2, FlagStackNone, 0, FlagStackNone}
var Putstatic = Instruction{byte(0xb3), 1, FlagStackNone, 0, FlagStackNone}
var Ret = Instruction{byte(0xa9), 0, FlagStackNone, 0, FlagStackNone}
var Return = Instruction{byte(0xb1), 0, FlagStackNone, 0, FlagStackEmpty}
var Saload = Instruction{byte(0x35), 2, FlagStackNone, 1, FlagStackNone}
var Sastore = Instruction{byte(0x56), 3, FlagStackNone, 0, FlagStackNone}
var Sipush = Instruction{byte(0x11), 0, FlagStackNone, 1, FlagStackNone}
var Swap = Instruction{byte(0x5f), 2, FlagStackNone, 2, FlagStackNone}
var Tableswitch = Instruction{byte(0xaa), 1, FlagStackNone, 0, FlagStackNone}
