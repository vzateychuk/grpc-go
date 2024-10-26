package sample

import (
	"math/rand"

	"vez/grpc/pb"
)

func NewKeyboard() *pb.Keyboard {
	kb := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(), // Ensure randomBool is defined
	}

	return kb
}

func NewCPU() *pb.Cpu {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	cores := uint32(rand.Intn(8) + 1)

	cpu := &pb.Cpu{
		Brand: brand,
		Name:  name,
		Cores: cores,
	}

	return cpu
}

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(rand.Intn(64) + 1),
		Unit:  pb.Memory_GB,
	}

	return ram
}

func NewSDD() *pb.Storage {
	sdd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(rand.Intn(1024) + 1),
			Unit:  pb.Memory_GB,
		},
	}
	return sdd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(rand.Intn(1024) + 1),
			Unit:  pb.Memory_GB,
		},
	}
	return hdd
}

func NewScreen() *pb.Screen {
	// 	Resolution *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// Panel      Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=Screen_Panel" json:"panel,omitempty"`

	screen := &pb.Screen{
		Size: float32(rand.Intn(24) + 1),
		Resolution: &pb.Screen_Resolution{
			Width:  uint32(rand.Intn(3840) + 1),
			Height: uint32(rand.Intn(2160) + 1),
		},
		Panel: randomScreenPanel(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:    randomID(),
		Brand: brand,
		Name:  name,
		Cpu:   NewCPU(),
		Ram:   NewRam(),
		Storages: []*pb.Storage{
			NewSDD(),
			NewHDD(),
		},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
	}

	return laptop
}
