package views

import (
	"fmt"
	"strings"

	"github.com/sebafudi/bokplass-backend/internal/models"
)

func RenderBooks(books []models.Book) string {
	var sb strings.Builder
	for _, book := range books {
		sb.WriteString("<tr>")
		sb.WriteString(fmt.Sprintf("<td>%s</td>", book.Title))
		if book.Authors != nil {
			sb.WriteString("<td>")
			for i, author := range book.Authors {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(author.Name)
			}
			sb.WriteString("</td>")
		} else {
			sb.WriteString("<td></td>")
		}
		sb.WriteString("</tr>")
	}
	return sb.String()
}

// func RenderBooks(books []models.Book) string {
// 	var sb strings.Builder
// 	sb.WriteString("<h1>Books</h1>")
// 	sb.WriteString("<ul>")
// 	for _, book := range books {
// 		sb.WriteString(fmt.Sprintf("<li>%s</li>", book.Title))
// 		if book.Authors != nil {
// 			sb.WriteString("<ul>")
// 			for _, author := range book.Authors {
// 				sb.WriteString(fmt.Sprintf("<li>%s</li>", author.Name))
// 			}
// 			sb.WriteString("</ul>")
// 		}
// 	}
// 	sb.WriteString("</ul>")
// 	return sb.String()
// }
