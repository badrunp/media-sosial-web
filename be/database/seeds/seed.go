package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

type Seed interface {
	Run()
}

type Seeds interface {
	Add(seed Seed)
}

type seed struct {
	Seeds []Seed
}

func NewSeed() *seed {
	return &seed{}
}

func (s *seed) Add(sd Seed) {
	s.Seeds = append(s.Seeds, sd)
}

func RegisterSeeders(db *gorm.DB, args []string) string {
	if len(args[1:]) > 0 {
		isSeed := args[1:][0] == "seed:run"
		if isSeed {
			AddSeeds(db, args[0])

			if args[0] != "testing" {
				os.Exit(1)
			}
			return "seed success"
		}
	}else{
		if mode := os.Getenv("NODE_ENV"); mode == "production" {
			AddSeeds(db, "")
			return "seed success"
		}
	}

	return "seed failed"
}

func AddSeeds(db *gorm.DB, mode string){
	seed := NewSeed()
	userSeed := UserSeed(db)

	seed.Add(userSeed)

	if mode != "testing" {
		for _, mg := range seed.Seeds {
			mg.Run()
		}
	}

	fmt.Println("seed success")
}