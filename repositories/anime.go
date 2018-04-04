package repositories

import (
	"github.com/suppayami/pettanko/models"
)

// AnimeRepository gives access to anime list
type AnimeRepository interface {
	AllAnime() []*models.Anime
	Anime(id int) *models.Anime
}

type animeStaticRepo struct {
	animeList map[int]*models.Anime
}

func (repo *animeStaticRepo) AllAnime() []*models.Anime {
	result := make([]*models.Anime, 0, len(repo.animeList))
	for _, anime := range repo.animeList {
		result = append(result, anime)
	}
	return result
}

func (repo *animeStaticRepo) Anime(id int) *models.Anime {
	return repo.animeList[id]
}

// NewAnimeStaticRepo - animeStaticRepo constructor
func NewAnimeStaticRepo(animeList []*models.Anime) AnimeRepository {
	list := make(map[int]*models.Anime)
	for _, anime := range animeList {
		list[anime.ID] = anime
	}
	return &animeStaticRepo{animeList: list}
}
