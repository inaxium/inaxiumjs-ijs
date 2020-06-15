package public

import (
	"log"
)

type Meta struct {
	Version          string
	Destination          string
	Type          string
}

func (m *Meta) Log() {
	log.Printf("Version => %s", m.Version)
	log.Printf("Target => %s", m.Destination)
	log.Printf("Type => %s", m.Type)
}
