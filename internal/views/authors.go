package views

import (
	"fmt"
	"strings"

	"github.com/sebafudi/bokplass-backend/internal/models"
)

func RenderAuthors(authors []models.Author) string {
	var sb strings.Builder
	sb.WriteString("<h1>Authors</h1>")
	sb.WriteString("<ul>")
	for _, author := range authors {
		sb.WriteString(fmt.Sprintf("<li>%s</li>", author.Name))
		if author.Books != nil {
			sb.WriteString("<ul>")
			for _, book := range author.Books {
				sb.WriteString(fmt.Sprintf("<li>%s</li>", book.Title))
			}
			sb.WriteString("</ul>")
		}
	}
	sb.WriteString("</ul>")
	return sb.String()
}
