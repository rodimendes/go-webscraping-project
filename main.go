package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

type MyNews struct {
	ID int64 
	MainSite string `json:"MainSite"`
	ArticleTitle string `json:"ArticleTitle"`
	ArticleURL string `json:"ArticleURL"`
	Date string `json:"date"`
}

func main() {

	// Capture connections properties
	cfg := mysql.Config {
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "news",
	}

	// Get a database handle
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	fileName := "news.json"

	collector := colly.NewCollector(
		// Visit only indicated domain
		colly.AllowedDomains("www.globo.com", "globo.com"),
	)

	searchTerm := "São Paulo"

	// Set User-Agent to mimic a browser
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	})

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, searchTerm) {
			currentTime := time.Now().Local().Format("2006/01/02 15:04:05")
			domain := e.Request.URL.Host
			
			var newsSP []MyNews

			content, err := os.ReadFile(fileName)
			if err == nil {
				err = json.Unmarshal(content, &newsSP)
				if err != nil {
					fmt.Println("Error parsing JSON:", err)
					return
				}
			}

			id := len(newsSP) + 1

			news := MyNews {
				ID: int64(id),
				MainSite: domain,
				ArticleTitle: e.Text,
				ArticleURL: e.Attr("href"),
				Date: currentTime,
			}

			repetido := false
			for _, existing := range newsSP {
				if news.ArticleURL == existing.ArticleURL {
					repetido = true
					break
				}
			}
			
			if !repetido {
				newsSP = append(newsSP, news)
				newsID, err := newNews(db, MyNews{
					MainSite: domain,
					ArticleTitle: e.Text,
					ArticleURL: e.Attr("href"),
					Date: currentTime,
				})
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("ID of added news: %v\n", newsID)
			}

			new_data, err := json.MarshalIndent(newsSP, "", " ")
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				return
			}

			err = os.WriteFile(fileName, new_data, 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
		}
	})
		
	// Before making a request print "Visiting ..."
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://www.globo.com
	collector.Visit("https://www.globo.com")

	sites, err := newsByMainSite(db, "www.globo.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("News found: %v\n", sites)


}

func newsByMainSite(db *sql.DB, site string) ([]MyNews, error) {
	var noticias []MyNews
	//>
	if db == nil {
		log.Fatalf("Database connection is nil")
	}
	//<
	rows, err := db.Query("SELECT id, mainSite, articleTitle, articleURL, `date` FROM mynews WHERE mainSite = ?", site)
	if err != nil {
		return nil, fmt.Errorf("newsByMainSite %q: %v", site, err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var notic MyNews
		if err := rows.Scan(&notic.ID, &notic.MainSite, &notic.ArticleTitle, &notic.ArticleURL, &notic.Date); err != nil {
			return nil, fmt.Errorf("newsByMainSite %q: %v", site, err)
		}
		noticias = append(noticias, notic)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("newsByMainSite %q: %v", site, err)
	}
	return noticias, nil
}

func newNews(db *sql.DB, notic MyNews) (int64, error) {
	result, err := db.Exec("INSERT INTO mynews (mainSite, articleTitle, articleURL, `date`) VALUES (?, ?, ?, ?)", notic.MainSite, notic.ArticleTitle, notic.ArticleURL, notic.Date)
	if err != nil {
		return 0, fmt.Errorf("addNews: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addNews: %v", err)
	}
	return id, nil
}