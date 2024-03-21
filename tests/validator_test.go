package tests

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationField(t *testing.T) {
	validate := validator.New()
	var user string = ""

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	var validate *validator.Validate = validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTagValidation(t *testing.T) {
	validate := validator.New()

	username := "fauzan123"

	err := validate.Var(username, "required, number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidatorWithParam(t *testing.T) {
	validate := validator.New()
	username := "fauzan123"

	err := validate.Var(username, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// validator with struct

type User struct {
	Username string `validate:"required,min=5,max=10,email"`
	Password string `validate:"required,min=5"`
}

func TestValidateStruct(t *testing.T) {
	validate := validator.New()
	loginRequest := &User{
		Username: "Fauzan12345678910",
		Password: "adadasdadsad",
	}

	err := validate.Struct(loginRequest)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// jika menggunakan validation error
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type Register struct {
	Username        string `validate:"required,email"`
	Password        string `validate:"required,min=5"`
	ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
}

func TestEqualFieldInStruct(t *testing.T) {
	validate := validator.New()
	registerRequest := &Register{
		Username:        "fauzannurhidayat8@gmail.com",
		Password:        "Makan Siang Gratis Cuyyy",
		ConfirmPassword: "Makan Siang Gratis Cuyyy",
	}

	err := validate.Struct(registerRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

// validasi nested Struct

type Address struct {
	City    string `validate:"required,min=5"`
	Country string `validate:"required,min=1,max=5"`
}

type Customer struct {
	Name    string  `validate:"required,min=5"`
	Address Address `validate:"required"`
}

func TestNestedStruct(t *testing.T) {
	validation := validator.New()
	customer := &Customer{
		Name: "Fauzan Nur Hidayat",
		Address: Address{
			City:    "Kebumen",
			Country: "Indonesia",
		},
	}
	err := validation.Struct(customer)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

// menggunakan validasi dive agar bisa memvalidasi tiap slice-nya
type CustomerNew struct {
	Name    string    `validate:"required,min=5"`
	Address []Address `validate:"required,dive"`
}

func TestNestedStructValidationDive(t *testing.T) {
	validation := validator.New()
	customer := &CustomerNew{
		Name: "Fauzan Nur Hidayat",
		Address: []Address{
			{
				City:    "Kebumen",
				Country: "Indonesia",
			},
			{
				City:    "Kebumen",
				Country: "Indonesia",
			},
		},
	}
	err := validation.Struct(customer)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

type NewUser struct {
	Name    string   `validate:"required,min=5,max=10"`
	Hobbies []string `validate:"dive,required,min=1,max=5"`
}

func TestValidationInCollection(t *testing.T) {
	validation := validator.New()
	newUser := &NewUser{
		Name: "Fauzan",
		Hobbies: []string{
			"Makan",
			"Minum",
			"Nonton Film",
			"Ngaji",
			"Sholat",
			"Liburan",
		},
	}

	err := validation.Struct(newUser)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
	/*
		hasilnya :
		error Hobbies[2] on tag max with error Key: 'NewUser.Hobbies[2]' Error:Field validation for 'Hobbies[2]' failed on the 'max' tag
		error Hobbies[4] on tag max with error Key: 'NewUser.Hobbies[4]' Error:Field validation for 'Hobbies[4]' failed on the 'max' tag
		error Hobbies[5] on tag max with error Key: 'NewUser.Hobbies[5]' Error:Field validation for 'Hobbies[5]' failed on the 'max' tag
	*/
}

type School struct {
	Name string `validate:"required,min=1,max=5"`
}

type Student struct {
	NIM     string            `validate:"required"`
	Name    string            `validate:"required"`
	Address []Address         `validate:"required,dive"`
	Hobbies []string          `validate:"dive,required,min=1"`
	Schools map[string]School `validate:"dive,keys,required,min=1,endkeys,required"`
}

func TestValidationWithMap(t *testing.T) {
	validation := validator.New()
	student := Student{
		NIM:  "12345",
		Name: "Fauzan Nur",
		Address: []Address{
			{
				City:    "Kebumen",
				Country: "Indo",
			},
			{
				City:    "Cakung",
				Country: "Indo",
			},
		},
		Hobbies: []string{
			"makan",
			"minum",
		},
		Schools: map[string]School{
			"SMP": {
				Name: "SMP N 2 Pejagoan",
			},
			"SMK": {
				Name: "SMK N 2 Kebumen",
			},
		},
	}

	err := validation.Struct(student)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestRegisterAlias(t *testing.T) {
	validation := validator.New()
	validation.RegisterAlias("BatasMaksimumKarakter", "max=120")

	nama := "Fauzan Nur Hidayat Ga Ngerokok"
	err := validation.Var(nama, "BatasMaksimumKarakter")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// membuat custom validation pada field struct atau variabel
func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return false
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestMustValidPin(t *testing.T) {
	validation := validator.New()
	validation.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "081asd",
		Pin:   "123456asda",
	}

	err := validation.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrValidation(t *testing.T) {
	validation := validator.New()

	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required,numeric"`
	}

	request := Login{
		Username: "0812132",
		Password: "123456",
	}
	err := validation.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// megecek apakah field yang divalidasi itu memiliki nilai pada field yang dirujuk
func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		// username harus sama dengan Email atau Phone
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}
	request := User{
		Username: "0813345321",
		Email:    "fauzannurhidayat8@gmail.com",
		Phone:    "0813345321",
		Name:     "Fauzan Nur Hidayat",
	}

	err := validate.Struct(request) 
	if err != nil {
		panic(err)
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email string `validate:"required,email"`
	Phone string `validate:"required,numeric"`
	Password string `validate:"required"`
}

// custom validation struct
func MustValidRegisterSuccess(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// success
	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "validate_username", "")
	}
}

func TestCustomValidationStruct(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	registerRequest := RegisterRequest{
		Username: "081335457601",
		Email: "fauzannurhidayat8@gmail.com",
		Phone: "081335457601",
		Password: "haha123",
	}
	
	err := validate.Struct(registerRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}