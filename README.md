#Linkedin Scraper
This is a linkedin job scraper written in golang.

## How to run
Clone this repo:

SSH:
```$ git clone git@github.com:kauakirchner/go-scraper.git```

HTTPS: ```$ git clone https://github.com/kauakirchner/go-scraper.git```

Go into the project's root folder.

```$ cd path/to/go-scraper```

Run the project:.

```$ go run pkg/main.go```

Use the API:

GET endpoint `scrape`:

Available parameters:
- keywords
- excluded
- location
- remote

Example:
```http://localhost:8081/scrape?keywords=golang%20developer&excluded=junior&location=UK&remote=true```

Results:
```
  "status": 200,
  "message": "120 jobs found!",
  "jobs": [
    {
			"title": "Go Developer",
			"src": "https://uk.linkedin.com/jobs/view/go-developer-at-yo-hr-consultancy-4011858374?position=1&pageNum=0&refId=nHz1zYAVs8va7zH3Fxr3hg%3D%3D&trackingId=ll%2BzmnrtOP2wtjThxXbuCg%3D%3D&trk=public_jobs_jserp-result_search-card",
			"company": "YO HR Consultancy",
			"companySrc": "https://in.linkedin.com/company/yo-hr-consultancy-pvt-ltd?trk=public_jobs_jserp-result_job-search-card-subtitle",
			"location": "United Kingdom",
			"postedAt": "2024-08-28"
		},
		{
			"title": "Junior Golang Engineer at AIPRM",
			"src": "https://uk.linkedin.com/jobs/view/junior-golang-engineer-at-aiprm-at-aiprm-3891678486?position=2&pageNum=0&refId=nHz1zYAVs8va7zH3Fxr3hg%3D%3D&trackingId=FbsUrNR9FG2MU2ivpzc4SQ%3D%3D&trk=public_jobs_jserp-result_search-card",
			"company": "AIPRM",
			"companySrc": "https://www.linkedin.com/company/aiprm?trk=public_jobs_jserp-result_job-search-card-subtitle",
			"location": "Birmingham, England, United Kingdom",
			"postedAt": "2024-03-15"
		},
      ...
  ]
```