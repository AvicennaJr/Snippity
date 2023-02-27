package main

import "github.com/AvicennaJr/Snippity/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
