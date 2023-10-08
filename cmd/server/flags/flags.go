package flags

import (
	"flag"
	"os"
)

var (
	Addr string
)

func ParseFlags() {
	flag.StringVar(&Addr, "a", "localhost:8080", "Адрес сервера формата host:port")
	flag.Parse()

	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		Addr = envRunAddr
	}
}
