# htmltopdf

Convert HTML to PDF. Styling CSS is never fun. The only thing worse is styling a PDF.

## Dependencies

### Darwin

- wkhtmltopdf

### Linux

- wkhtmltopdf
- xvfb-run

## Usage

### Go Package

```` go

import "github.com/robertlestak/htmltopdf"

...

pdfpath, err := htmltopdf.Convert("/path/to/file.html")
if err != nil {
	// handle error
}

// pdfpath is the string path to the generated PDF file

````

### REST Server

````
cp .env-sample .env
````

````
vgo run cmd/htmltopdf/htmltopdf.go
````

````
curl http://localhost:6547/convert -f "html=@/path/to/file.html" -o /path/to/file.pdf
````

#### Docker Deployment

````
docker build . -f build/Dockerfile -t htmltopdf:latest
docker run -d -p 6547:6547 htmltopdf:latest
````
