package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func DateMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func CreateAPIResponse(status int, body string) *events.APIGatewayProxyResponse {

	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func EscapeString(t string) string {
	desc := strings.ReplaceAll(t, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")
	return desc
}

func ParsePaginationAndSorting(query map[string]string) (page, limit int, sortBy, order string, err error) {
	page, limit = 1, 10
	sortBy, order = "id", "ASC"

	if val := strings.TrimSpace(query["page"]); val != "" {
		p, err := strconv.Atoi(val)
		if err != nil || p < 1 {
			return 0, 0, "", "", fmt.Errorf("invalid 'page' parameter")
		}
		page = p
	}

	if val := strings.TrimSpace(query["limit"]); val != "" {
		l, err := strconv.Atoi(val)
		if err != nil || l < 1 {
			return 0, 0, "", "", fmt.Errorf("invalid 'limit' parameter")
		}
		limit = l
	}

	if val := strings.TrimSpace(query["sort_by"]); val != "" {
		allowed := map[string]bool{
			"id": true, "title": true, "description": true, "price": true,
			"category_id": true, "stock": true, "created_at": true,
		}
		if !allowed[val] {
			return 0, 0, "", "", fmt.Errorf("invalid 'sort_by' parameter")
		}
		sortBy = val
	}

	if val := strings.ToUpper(strings.TrimSpace(query["order"])); val != "" {
		if val != "ASC" && val != "DESC" {
			return 0, 0, "", "", fmt.Errorf("invalid 'order' parameter")
		}
		order = val
	}

	return
}

func ParseOrdersPaginationAndSorting(query map[string]string) (page int, from_date, to_date string, err error) {
	page = 1
	from_date = "1970-01-01"
	to_date = time.Now().Format("2006-01-02")

	if val := strings.TrimSpace(query["page"]); val != "" {
		p, err := strconv.Atoi(val)
		if err != nil || p < 1 {
			return 0, from_date, to_date, fmt.Errorf("invalid 'page' parameter")
		}
		page = p
	}

	if val := strings.TrimSpace(query["from_date"]); val != "" {

		from_date = val
	}

	if val := strings.TrimSpace(query["to_date"]); val != "" {
		to_date = val
	}

	return
}
