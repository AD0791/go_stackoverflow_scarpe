package services

import (
	"log"
	"regexp"
	"strconv"

	"github.com/AD0791/SO/scraper/models"
	"github.com/gocolly/colly"
)

type ScraperService struct{}

func NewScraperService() *ScraperService {
	return &ScraperService{}
}

func (s *ScraperService) ScrapeStackOverflowQuestions() (models.QuestionsResponse, error) {
	// URL of StackOverflow questions page
	url := "https://stackoverflow.com/questions"

	// Create a new Colly collector
	c := colly.NewCollector()

	var questions []models.Question

	// Find each question
	c.OnHTML(".s-post-summary.js-post-summary", func(e *colly.HTMLElement) {
		title := e.ChildText(".s-link")
		user := e.ChildText(".s-user-card--link a")

		votesText := e.ChildText(".s-post-summary--stats-item__emphasized")
		answersText := e.ChildText(".s-post-summary--stats-item:nth-child(2)")
		viewsText := e.ChildText(".s-post-summary--stats-item:nth-child(3)")

		votes := s.extractNumber(votesText)
		answers := s.extractNumber(answersText)
		views := s.extractNumber(viewsText)

		question := models.Question{
			Title:   title,
			User:    user,
			Votes:   votes,
			Answers: answers,
			Views:   views,
		}

		questions = append(questions, question)
	})

	// Visit the StackOverflow questions page
	err := c.Visit(url)
	if err != nil {
		log.Printf("Error fetching StackOverflow page: %v", err)
		return models.QuestionsResponse{}, err
	}

	return models.QuestionsResponse{
		Questions: questions,
		Total:     len(questions),
	}, nil
}

// extractNumber extracts the first numeric value from a string
func (s *ScraperService) extractNumber(text string) int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindStringSubmatch(text)

	if len(matches) > 0 {
		num, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}
