package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	return param
}

func LengthMultiParam[A any, B any](param1 A, param2 B) (A, B) {
	return param1, param2
}

func TestSample(t *testing.T) {
	assert.True(t, true)
}

func TestWithParam(t *testing.T) {
	fmt.Println(Length(3 * 2))
	fmt.Println(Length("Makan Siang"))
	fmt.Println(Length(true))
	fmt.Println(Length(0.4553434))

	// cara seperti ini berarti paramnya harus string, dan returnnya juga string
	var result1 string = Length[string]("Fauzan")
	fmt.Println(result1)
	var result2 int = Length[int](123)
	fmt.Println(result2)
}

func TestWithMultiParam(t *testing.T) {
	result1, result2 := LengthMultiParam("Makan", "Siang")
	fmt.Println(result1, result2)
}

func IsSame[C comparable](value1, value2 C) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[int](2, 2))
	assert.Equal(t, false, IsSame[string]("3", "2"))
}

// mencoba konsep inheritance pada generic
// employee

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// Manager

type Manager interface {
	// GetName() ini adalah mengambil dari sebuah generic GetName dengan struct-nya adalah Employee
	GetName() string
	GetManagerName() string
}

type ManagerImpl struct {
	Name string
}

func (m *ManagerImpl) GetName() string {
	return "Manager " + m.Name
}

func (m *ManagerImpl) GetManagerName() string {
	return m.Name
}

// Vice President

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type VicePresidentImpl struct {
	Name string
}

func (v *VicePresidentImpl) GetName() string {
	return v.Name
}

func (v *VicePresidentImpl) GetVicePresidentName() string {
	return v.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Manager Fauzan", GetName[Manager](&ManagerImpl{Name: "Fauzan"}))
	assert.Equal(t, "Fauzan", GetName[VicePresident](&VicePresidentImpl{Name: "Fauzan"}))
}

// type set
type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
	// dengan menggunakan tanda ~ pada tipe datanya, maka berarti tipe alias yang terkait dengan tipe data tersebut akan bisa diproses oleh generic
}

// kedua value tersebut datanya harus bisa dilakukan operasi menggunakan operator misalnya
func IsBig[T Number](value1 T, value2 T) T {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}

func TestIsBig(t *testing.T) {
	assert.Equal(t, 50, IsBig(4, 50))
	assert.Equal(t, float32(50), IsBig[float32](4, 50))
	assert.Equal(t, Age(50), IsBig[Age](4, 50))
}

// selain dibuat berupa function, generic bisa dibuat secara type declaration
type Bag[B any] []B

func GetInsideBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestGetInsideBag(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	GetInsideBag[int](numbers)

	names := Bag[any]{"fauzan", "susi", "rudi", 10, true}
	fmt.Println(names)
	// GetInsideBag[string](names)
}

// menggunakan generic di struct
type Data[T any] struct {
	Firstname T
	Lastname  T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirstname(firstname T) T {
	d.Firstname = firstname
	return firstname
}

func TestData(t *testing.T) {
	data := Data[string]{
		Firstname: "Fauzan",
		Lastname:  "Nurhidayat",
	}
	fmt.Println(data.Firstname, data.Lastname)
	// fmt.Println(data.SayHello("Fauzan"))
	assert.Equal(t, "Hello Fauzan", data.SayHello("Fauzan"))
	data.ChangeFirstname("Rudi")
	assert.Equal(t, "Rudi Nurhidayat", (data.Firstname + " " + data.Lastname))
}

// membuat generic di interface
// jika interface berupa generic, maka struct yang mengimplementasikannya juga harus berupa generic
type Cars[T any] interface {
	SetModel(value T)
	GetModel() T
}

func ChangeModel[T any](param Cars[T], value T) T {
	param.SetModel(value)
	return param.GetModel()
}

// implementasikan method yang ada di interface

type MyCar[T any] struct {
	Model T
}

func (g *MyCar[T]) SetModel(value T) {
	g.Model = value	
}

func (g *MyCar[T]) GetModel() T {
	return g.Model
}

func TestGenericInterface(t *testing.T) {
	myCar := MyCar[string]{}
	result := ChangeModel[string](&myCar, "Ferrari")
	fmt.Println(result)
	fmt.Println(myCar.Model)
}

// in line type constraint pada generic
// yang dimana kita tidak perlu membuat type set terlebih dahulu untuk membuat list tipe data apa saja yang bisa dimasukkan

func FindSmallest[T interface{int|int64|int32}](value1 T, value2 T) T {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}

func TestFindSmallest(t *testing.T) {
	assert.Equal(t, 50, FindSmallest(4, 50))
	assert.Equal(t, int64(50), FindSmallest[int64](4, 50))
}

// generic di type parameter
func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Fauzan","Susi","Rahmat","hah",
	}

	numbers := []int{
		1,2,3,4,5,
	}

	first := GetFirst[[]string, string](names)
	number := GetFirst[[]int, int](numbers)
	fmt.Println(first)
	fmt.Println(number)
}