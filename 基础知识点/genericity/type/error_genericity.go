package main

// 错误，类型形参不能单独使用
//type CommonType[T int | string | float32] T

//✗ 错误。T *int会被编译器误认为是表达式 T乘以int，而不是int指针
//type NewType [T * int][]T

// 上面代码再编译器眼中：它认为你要定义一个存放切片的数组，数组长度由 T 乘以 int 计算得到
//type NewType [T * int][]T

//✗ 错误。和上面一样，这里不光*被会认为是乘号，| 还会被认为是按位或操作
//type NewType2 [T*int | *float64][]T

//✗ 错误
//type NewType2 [T(int)][]T
type NewType[T interface{ *int }] []T
type NewType2[T interface{ *int | *float64 }] []T

// 如果类型约束中只有一个类型，可以添加个逗号消除歧义
type NewType3[T *int,] []T

//✗ 错误。如果类型约束不止一个类型，加逗号是不行的
type NewType4[T *int | *float32,] []T

//--------------------------

// 先定义个泛型类型 Slice[T]
type Slice[T int | string | float32 | float64] []T

// ✗ 错误。泛型类型Slice[T]的类型约束中不包含uint, uint8
//type UintSlice[T uint | uint8] Slice[T]

// ✓ 正确。基于泛型类型Slice[T]定义了新的泛型类型 FloatSlice[T] 。FloatSlice[T]只接受float32和float64两种类型
type FloatSlice[T float32 | float64] Slice[T]

// 在map中套一个泛型类型Slice[T]
type WowMap[T int | string] map[string]Slice[T]
