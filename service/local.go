package service

import (
	"github.com/suppayami/pettanko/models"
	"github.com/suppayami/pettanko/repositories"
)

var (
	animeList []*models.Anime
	animeRepo repositories.AnimeRepository
)

func init() {
	animeList = []*models.Anime{
		&models.Anime{
			ID:    0,
			Title: "Clannad",
			Tags:  []string{"drama", "comedy", "slice of life"},
		},

		&models.Anime{
			ID:    1,
			Title: "Puella Magi Madoka Magica",
			Tags:  []string{"mahou shoujo", "wtf?!"},
		},

		&models.Anime{
			ID:    2,
			Title: "Steins;Gate",
			Tags:  []string{"scifi", "time travel"},
		},
	}

	animeRepo = repositories.NewAnimeStaticRepo(animeList)
}

// AnimeRepository returns anime repo
func AnimeRepository() repositories.AnimeRepository {
	return animeRepo
}
