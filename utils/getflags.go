package utils

import "flag"

type Flags struct {
	Part     int
	Filepath string
}

func GetFlags() Flags {
	part := flag.Int("p", 1, "part 1 or 2")
	filepath := flag.String("i", "input", "input filepath")
	flag.Parse()

	return Flags{
		Part:     *part,
		Filepath: *filepath,
	}
}
