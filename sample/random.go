package sample

import (
	"math/rand"
	"vez/grpc/pb"

	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(2) {
	case 0:
		return pb.Keyboard_QWERTY
	case 1:
		return pb.Keyboard_DVORAK
	}
	return pb.Keyboard_QWERTY
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

// randomCPUBrand returns a random CPU brand.
func randomCPUBrand() string {
	brands := []string{"Intel", "AMD"}
	return randomStringFromSet(brands...)
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet("Core i7", "Xeon", "Atom")
	case "AMD":
		return randomStringFromSet("Ryzen 7", "EPYC", "Athlon")
	}
	return ""
}

func randomStringFromSet(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	return s[rand.Intn(len(s))]
}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(2) {
	case 0:
		return pb.Screen_IPS
	case 1:
		return pb.Screen_OLED
	}
	return pb.Screen_IPS
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	brands := []string{"Apple", "Dell", "Lenovo"}
	return randomStringFromSet(brands...)
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("MacBook Pro", "MacBook Air")
	case "Dell":
		return randomStringFromSet("Latitude", "Inspiron", "XPS")
	case "Lenovo":
		return randomStringFromSet("ThinkPad X1", "IdeaPad Yoga", "Legion")
	}
	return ""
}
