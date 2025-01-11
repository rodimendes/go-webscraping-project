<h2>Code description</h2>

<h3>This code is a Go program designed to scrape news articles from a specific website (www.globo.com) and store the collected information in both a JSON file and a MySQL database. Here's a brief explanation of its main components and functionality:</h3>
<h3>This code serves as an **educational** example to demonstrate various important concepts in web scraping, data storage, and database operations using Go.</h3>

<h4>1. Imports and Struct Definition:</h4>

   <p>- The program imports necessary packages for web scraping, database operations, and file handling.
   - It defines a `MyNews` struct to represent news articles with fields like ID, MainSite, ArticleTitle, ArticleURL, and Date.</p>

<h4>2. Database Connection:</h4>

   <p>- The program establishes a connection to a MySQL database using environment variables for credentials.</p>

<h4>3. Web Scraping:</h4>

   <p>- It uses the `colly` library to scrape the Globo website.</p>
   <p>- The scraper is configured to only visit allowed domains and set a User-Agent to mimic a browser.</p>

<h4>4. Article Collection:</h4>

   <p>- The program searches for links containing the term "São Paulo" on the webpage.</p>
   <p>- When a matching link is found, it extracts relevant information and creates a `MyNews` object.</p>

<h4>5. Data Storage:</h4>

   <p>- The collected news articles are stored in two ways:</p>
     <ul>
     <li>a. Appended to a JSON file named "news.json".<li>
     <li>b. Inserted into a MySQL database table named "mynews".</li>
    </ul>

<h4>6. Duplicate Prevention:</h4>

   <p>- The code checks for duplicate articles based on the URL before adding new entries.</p>

<h4>7. Database Operations:</h4>

   <p>- Two functions, `newsByMainSite` and `newNews`, handle database queries for retrieving and inserting news articles, respectively.</p>

<h4>8. Main Execution:</h4>
   <p>- The program initiates the scraping process by visiting "https://www.globo.com".</p>
   <p>- After scraping, it retrieves and prints all news articles from the database for the "www.globo.com" site.</p>

<h3>This script automates the process of collecting, storing, and retrieving news articles from a specific source, making it easier to track and analyze news content over time.</h3>
