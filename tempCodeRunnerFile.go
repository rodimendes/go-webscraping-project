func newsByMainSite(db *sql.DB, site string) ([]MyNews, error) {
	var noticias []MyNews
	//>
	if db == nil {
		log.Fatalf("Database connection is nil")
	}