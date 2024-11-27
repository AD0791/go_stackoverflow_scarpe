package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO Fix test error
func TestScraperService_ScrapeStackOverflowQuestions(t *testing.T) {
	// Create a new server to serve the mock StackOverflow page
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<div class="s-post-summary js-post-summary">
				<a href="#" class="s-link">Question 1</a>
				<a href="#" class="s-user-card--link">User 1</a>
				<div class="s-post-summary--stats-item__emphasized">10</div>
				<div class="s-post-summary--stats-item">5</div>
				<div class="s-post-summary--stats-item">100</div>
			</div>
			<div class="s-post-summary js-post-summary">
				<a href="#" class="s-link">Question 2</a>
				<a href="#" class="s-user-card--link">User 2</a>
				<div class="s-post-summary--stats-item__emphasized">20</div>
				<div class="s-post-summary--stats-item">3</div>
				<div class="s-post-summary--stats-item">150</div>
			</div>
		`))
	}))
	defer server.Close()

	// Create a new ScraperService and scrape the mock StackOverflow page
	scraperService := NewScraperService()
	questionsResponse, err := scraperService.ScrapeStackOverflowQuestions()

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, 2, len(questionsResponse.Questions))
	assert.Equal(t, 2, questionsResponse.Total)

	assert.Equal(t, "Question 1", questionsResponse.Questions[0].Title)
	assert.Equal(t, "User 1", questionsResponse.Questions[0].User)
	assert.Equal(t, 10, questionsResponse.Questions[0].Votes)
	assert.Equal(t, 5, questionsResponse.Questions[0].Answers)
	assert.Equal(t, 100, questionsResponse.Questions[0].Views)

	assert.Equal(t, "Question 2", questionsResponse.Questions[1].Title)
	assert.Equal(t, "User 2", questionsResponse.Questions[1].User)
	assert.Equal(t, 20, questionsResponse.Questions[1].Votes)
	assert.Equal(t, 3, questionsResponse.Questions[1].Answers)
	assert.Equal(t, 150, questionsResponse.Questions[1].Views)
}
