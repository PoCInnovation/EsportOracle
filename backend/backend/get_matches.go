package backend

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	PandaScoreAPIToken = os.Getenv("PANDASCORE_API_TOKEN")
}

var BaseURL = "https://api.pandascore.co"
var PandaScoreAPIToken string

func SendResponseClient(w http.ResponseWriter, req *http.Request) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("http.DefaultClient.Do(req): %v", err), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("unexpected status code: %d", res.StatusCode), res.StatusCode)
		return
	}

	// Copy the response body directly to the client
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, res.Body); err != nil {
		http.Error(w, fmt.Sprintf("io.Copy(w, res.Body): %v", err), http.StatusInternalServerError)
		return
	}
}

func GetTeamFromID(w http.ResponseWriter, r *http.Request) {
	teamID := mux.Vars(r)["teamID"]
	fmt.Println("Fetching team with ID:", teamID)
	url := fmt.Sprintf("%s/csgo/teams?filter[acronym]=%s", BaseURL, teamID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("http.NewRequest(\"GET\", url, nil): %v", err), http.StatusInternalServerError)
		return
	}

	req.Header.Add("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Add("Accept", "application/json")

	SendResponseClient(w, req)

	fmt.Println("Successfully fetched team from ID from PandaScore API")

}

func GetMatchByID(w http.ResponseWriter, r *http.Request) {
	matchID := mux.Vars(r)["matchID"]
	fmt.Println("Fetching match with ID:", matchID)
	url := fmt.Sprintf("%s/csgo/matches/running?filter[id]=%s", BaseURL, matchID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("http.NewRequest(\"GET\", url, nil): %v", err), http.StatusInternalServerError)
		return
	}

	req.Header.Add("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Add("Accept", "application/json")

	SendResponseClient(w, req)
}

func GetCurrentMatches(w http.ResponseWriter, r *http.Request) {
	teamIDs := r.URL.Query().Get("teamId")

	var url string;

	if (teamIDs != "") {
		url = fmt.Sprintf("%s/matches/running?filter[opponent_id]=%s", BaseURL, teamIDs)
	} else {
		url = fmt.Sprintf("%s/matches/running", BaseURL)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("http.NewRequest(\"GET\", url, nil): %v", err), http.StatusInternalServerError)
		return
	}

	req.Header.Add("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Add("Accept", "application/json")

	SendResponseClient(w, req)

	fmt.Println("Successfully fetched current matches from PandaScore API")
}

func GetPastMatches(w http.ResponseWriter, r *http.Request) {
	teamIDs := r.URL.Query().Get("teamId")

	var url string;

	if (teamIDs != "") {
		url = fmt.Sprintf("%s/matches/?filter[status]=finished&filter[opponent_id]=%s&sort=-begin_at&per_page=50&page=1", BaseURL, teamIDs)
	} else {
		url = fmt.Sprintf("%s/matches?filter[status]=finished&sort=-begin_at&per_page=50&page=1", BaseURL)
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		http.Error(w, fmt.Sprintf("http.NewRequest(\"GET\", url, nil): %v", err), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Add("Accept", "application/json")

	SendResponseClient(w, req)

	fmt.Println("Successfully fetched past matches from PandaScore API")
}

func GetUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	teamIDs := r.URL.Query().Get("teamId")

	var url string;

	if (teamIDs != "") {
		url = fmt.Sprintf("%s/matches/?filter[status]=not_started&filter[opponent_id]=%s&sort=begin_at&per_page=50&page=1", BaseURL, teamIDs)
	} else {
		url = fmt.Sprintf("%s/matches?filter[status]=not_started&sort=begin_at&per_page=50&page=1", BaseURL)
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		http.Error(w, fmt.Sprintf("http.DefaultClient.Do(req): %v", err), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Authorization", "Bearer "+PandaScoreAPIToken)
	req.Header.Add("Accept", "application/json")

	SendResponseClient(w, req)

	fmt.Println("Successfully fetched upcoming matches from PandaScore API")
}

// SetupRoutes registers the REST API endpoints and returns a mux.Router.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/MatchByID/{matchID}", GetMatchByID).Methods("GET")
	//router.HandleFunc("/TeamFromID/{teamID}", GetTeamFromID).Methods("GET")
	router.HandleFunc("/matches/current", GetCurrentMatches).Methods("GET")
	router.HandleFunc("/matches/past", GetPastMatches).Methods("GET")
	router.HandleFunc("/matches/upcoming", GetUpcomingMatches).Methods("GET")
	return router
}
