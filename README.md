# Automobile Web Scraper

<img height="20px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"> &nbsp; web scraper is designed to extract automobile listings from [mobile.bg](https://www.mobile.bg), a popular Bulgarian automobile website. The scraper collects data such as the title, price, and parameters of the vehicles listed on multiple pages and consolidates the information into a single string for further processing or storage.

## Features

- Scrapes automobile listings from [mobile.bg](https://www.mobile.bg).
- Extracts relevant details such as title, price, and additional parameters.
- Supports scraping multiple pages concurrently for improved efficiency.
- Automatically detects and handles different character encodings of the web pages.


## Usage

### Scrape a Single Page

You can scrape a single page from mobile.bg by calling the `ScrapePage` function and passing the URL of the page you want to scrape.

### Scrape Multiple Pages

To scrape multiple pages from mobile.bg, use the `ScrapeAllPages` function. You can specify the range of pages you want to scrape by passing the start and end page numbers.
