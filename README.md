# News App made with Go

Build and deploy a complete web service in Go's standard library.

Using [NewsAPI.org](https://newsapi.org/)'s JSON API integration to colect data,
fetch news articles matching a user search query
and present the results on the page. News API is a simple HTTP REST
API for searching and retrieving live articles from all over the web.

Input a search term and News API will scan thousands of news websites and
organsie the information of the latest articles.

You can search for articles with any combination of the following criteria:

- **Keyword or phrase.** Eg: find all articles containing the word 'Microsoft'.
- **Date published.** Eg: find all articles published yesterday.
- **Source name.** Eg: find all articles by 'TechCrunch'.
- **Source domain name.** Eg: find all articles published on thenextweb.com.
- **Language.** Eg: find all articles written in English.

---

When building the app locally you will requires a `.env` file to be created in
the root with the following [info](info.md):

```txt
PORT=

NEWS_API_KEY=
```

---

Technology Stack:

- HTML
- CSS
- Javascript
- Golang

