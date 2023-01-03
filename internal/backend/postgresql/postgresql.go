package postgresql

import (
	"context"
	"fmt"
	"sencha-twirp-rpc/internal/backend"
	"sencha-twirp-rpc/internal/models"

	"gorm.io/gorm"
)

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) backend.Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}

func (b *PostgreSQLBackend) AddThemes(ctx context.Context) error {
	var themes map[string]*models.Theme = map[string]*models.Theme{
		"default":      {Background: "black", Foreground: "white"},
		"cfi-blue":     {Background: "#1974D2", Foreground: "white"},
		"aqua":         {Background: "black", Foreground: "#33ffd0"},
		"white-orange": {Background: "#ff8000", Foreground: "white"},
		"light-blue":   {Background: "black", Foreground: "#33bbff"},
		"yellow":       {Background: "black", Foreground: "yellow"},
		"pinkish":      {Background: "black", Foreground: "#DE3163"},
		"dark":         {Background: "black", Foreground: "#f2f2f2"},
		"light":        {Background: "#f2f2f2", Foreground: "black"},
		"orange":       {Background: "black", Foreground: "#EA5B0C"},
		"cyan":         {Background: "black", Foreground: "#4CBEC5"},
		"green":        {Background: "black", Foreground: "#00CC11"},
		"pink":         {Background: "black", Foreground: "#FF6666"},
		"faint-orange": {Background: "black", Foreground: "#996633"},
		"neon-blue":    {Background: "black", Foreground: "#0033FF"},
		"ultra-green":  {Background: "black", Foreground: "#0AFF84"},
		"ultra-purple": {Background: "black", Foreground: "#8709F4"},
		"iron-gray":    {Background: "black", Foreground: "#52595D"},
		"bright-gray":  {Background: "black", Foreground: "#dcdcdc"},
		"bright-blue":  {Background: "black", Foreground: "#006EF0"},
	}

	for k := range themes {
		fmt.Println("themes", k, themes[k].Background, themes[k].Foreground)

		theme := models.Theme{
			Name:       k,
			Background: themes[k].Background,
			Foreground: themes[k].Foreground,
		}

		result := b.DB.Create(&theme)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (b *PostgreSQLBackend) GetThemes(ctx context.Context) ([]*models.Theme, error) {
	var themes []*models.Theme
	result := b.DB.Find(&themes)
	if result == nil {
		return nil, result.Error
	}
	return themes, nil
}

func (b *PostgreSQLBackend) CreateTheme(ctx context.Context, name string, background string, foreground string) (*models.Theme, error) {
	theme := models.Theme{
		Name:       name,
		Background: background,
		Foreground: foreground,
	}

	result := b.DB.Create(&theme)
	if result.Error != nil {
		return nil, result.Error
	}

	return &theme, nil
}

func (b *PostgreSQLBackend) GetTheme(ctx context.Context, name string) (*models.Theme, error) {
	theme := models.Theme{}

	result := b.DB.First(&theme, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}

	return &theme, nil
}
