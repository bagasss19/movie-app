package repository

import "gorm.io/gorm"

func filterMovieByID(query *gorm.DB, ID uint64) *gorm.DB {
	return query.Where("id = ?", ID)
}

func filterMovieSearchByTitle(query *gorm.DB, title string) *gorm.DB {
	return query.Where("title ILIKE ?", "%"+title+"%")
}

func filterMovieSearchByDescription(query *gorm.DB, description string) *gorm.DB {
	return query.Where("description ILIKE ?", "%"+description+"%")
}

func filterMovieSearchByArtist(query *gorm.DB, artist string) *gorm.DB {
	return query.Where("artist ILIKE ?", "%"+artist+"%")
}

func filterMovieSearchByGenre(query *gorm.DB, genre string) *gorm.DB {
	return query.Where("genre ILIKE ?", "%"+genre+"%")
}
