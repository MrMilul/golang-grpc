package sample

import "example.com/laptop_store/proto"

func NewCPU() *pb.CPU{
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	minGhz := randFlout64(2.5, 5.5)
	cpu := &pb.CPU{
		Name: name,
		Brand: brand, 
		Cores: uint32(randomInt(2, 8)),
		Threads:  uint32(randomInt(2, 8)),
		MaxGhz: randFlout64(minGhz, 10),
		MinGhz: minGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU{
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	minGhz := randFlout64(2.5, 5.5)
	memory := &pb.Memory{
		Value: int32(randomInt(3, 10)),
		Unit: pb.Memory_MEGABYTE,
	}
	gpu := &pb.GPU{
		Name: name,
		Brand: brand, 
		Cores: uint32(randomInt(2, 8)),
		Threads:  uint32(randomInt(2, 8)),
		Memory: memory,
		MaxGhz: randFlout64(minGhz, 10),
		MinGhz: minGhz,
	}
	return gpu
}

func NewLaptop() *pb.Laptop{
	id := randomId()
	laptop := &pb.Laptop{
		Id: id,
		Brand: "Apple",
		Name: "M1 Pro",
		Cpu: NewCPU(),
		Gpus: [] *pb.GPU{NewGPU(), NewGPU()},
		Weight: &pb.Laptop_WightKg{
			WightKg: int32(2),
		},

	}
	return laptop
}