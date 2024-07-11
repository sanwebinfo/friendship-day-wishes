package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

const port = 6054

func escapeText(text string) string {
	return html.EscapeString(text)
}

func cleanName(name string) string {
	return strings.ReplaceAll(name, "-", " ")
}

func asciiArt(name string) string {
	cleanedName := cleanName(name)
	art := `
 ★─▀██▀▀▀█
 ★──██▄█
 ★──██▀█
 ★─▄██ ANTASTIC Friend ★★★
	`

	quotes := []string{
		" Friendship is the compass\n that guides us\n through life's storm",
		" Friendship doubles your joy\n and divides your sorrow",
	}

	quote := quotes[len(name)%len(quotes)]

	return fmt.Sprintf("\n wishes@%s:~💚$\n%s\n\n%s", escapeText(cleanedName), art, quote)
}

func generateSlug(name string) string {
	var result []string
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, string(unicode.ToLower(r)))
		} else if r == ' ' {
			result = append(result, "-")
		}
	}
	return strings.Join(result, "")
}

func validateName(name string) (string, error) {
	if len(name) == 0 || len(name) > 36 {
		return "", fmt.Errorf("name length must be between 1 and 36 characters")
	}

	if valid := regexp.MustCompile(`^[\p{L}\p{N}\p{P}\p{Zs}\p{M}\p{Sm}\p{So}\p{Sk}]+$`).MatchString(name); !valid {
		return "", fmt.Errorf("name contains invalid characters")
	}

	return name, nil
}

// wishHTMLHandler handles requests for HTML responses for wishes.
func wishHTMLHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	validName, err := validateName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name = escapeText(validName)
	asciiText := asciiArt(name)
	slugText := generateSlug(name)
	baseURL := fmt.Sprintf("https://%s", r.Host)
	//wishURL := fmt.Sprintf("https://%s/wish/web", r.Host)
	TextURL := fmt.Sprintf("https://%s/wish/text", r.Host)
	shareURL := fmt.Sprintf("%s/wish/web?name=%s", baseURL, slugText)

	setHTMLHeaders(w)
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en" prefix="og: https://ogp.me/ns#">
<head>
    <meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="shortcut icon" type="image/x-icon" href="/favicon.ico" />
    <link rel="icon" type="image/png" sizes="196x196" href="/favicon-196.png" />

    <title>%s : Happy Friendship Wishes</title>
    <meta name="description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name."/>
    <meta name="canonical" href="%s">

    <meta property="og:site_name" content="%s : Happy Friendship Wishes">
    <meta property="og:type" content="website">
    <meta property="og:title" content="%s : Happy Friendship Wishes">
    <meta property="og:description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.">
    <meta property="og:url" content="%s">
    <meta property="og:image" content="https://img.sanweb.info/friend/friend?name=%s">
    <meta property="og:image:alt" content="%s : Happy Friendship Wishes" />
    <meta property="og:image:width" content="1080">
    <meta property="og:image:height" content="1080">

    <meta name="twitter:title" content="%s : Happy Friendship Wishes">
    <meta name="twitter:description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.">
    <meta name="twitter:url" content="%s">
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:image" content="https://img.sanweb.info/friend/friend?name=%s">

    <link rel="preconnect" href="https://cdnjs.cloudflare.com">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.9.3/css/bulma.min.css" integrity="sha512-IgmDkwzs96t4SrChW29No3NXBIBv8baW490zk5aXvhCD8vuZM3yUSkbyTBcXohkySecyzIrUwiF/qV0cuPcL3Q==" crossorigin="anonymous" referrerpolicy="no-referrer" />

	<style>
        html, body {
            min-height: 100vh;
        }
        body {
            font-family: monospace;
            background-color: #ffda79;
        }
        #quote-container {
            margin: 10px auto;
            padding: 20px;
            background-color: #4b4b4b;
        }
        #quote {
            font-size: 20px;
            margin-bottom: 20px;
            color: #333;
        }
        #quote-card {
            margin: 20px auto;
        }
        pre {
            font-family: monospace;
            font-size: 14px;
            background-color: #192a56;
            color: #fdcb6e;
            text-shadow: 0 0 3px #FFC312;
            padding: 20px;
            border-radius: 5px;
            word-wrap: break-word;
            overflow-x: auto;
            line-height: inherit;
        }
    </style>

<link rel="preconnect" href="https://img.sanweb.info">

</head>
<body>

<section class="section">
    <div class="container">
        <div id="quote-card" class="card">
            <div id="quote-container">
                <pre>
                %s
                </pre>
                <br>
            </div>
        </div>
        <br>
        <pre>$ curl -G --data-urlencode "name=%s" %s<br><br>$ http -b GET "%s" "name==%s"<br></pre>
    </div>
    <br>
    <div class="container">
        <div class="columns is-centered">
            <div class="column is-half">
                <div class="card">
                    <div class="card-image">
                        <figure class="image">
                            <img src="https://img.sanweb.info/friend/friend?name=%s" alt="Happy Friendship Day" loading="lazy">
                        </figure>
                    </div>
                </div>
            </div>
        </div>
    </div>
</section>

</body>
</html>
`, cleanName(name), shareURL, cleanName(name), cleanName(name), shareURL, slugText, cleanName(name), cleanName(name), shareURL, slugText, asciiText, cleanName(name), TextURL, TextURL, cleanName(name), cleanName(name))
}

// wishTextHandler handles requests for plain text responses for wishes.
func wishTextHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	validName, err := validateName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name = escapeText(validName)
	asciiText := asciiArt(name)
	slugText := generateSlug(name)
	baseURL := fmt.Sprintf("https://%s", r.Host)
	shareURL := fmt.Sprintf("%s/wish/web?name=%s", baseURL, slugText)

	setTextHeaders(w)
	fmt.Fprintf(w, "%s\n\n Web View URL: %s\n\n", asciiText, shareURL)
}

// setHTMLHeaders sets headers specific to HTML responses.
func setHTMLHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	setSecurityHeaders(w)
}

// setTextHeaders sets headers specific to plain text responses.
func setTextHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	setSecurityHeaders(w)
}

func setSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	setTextHeaders(w)
	fmt.Fprint(w, "404 Page Not Found")
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	setTextHeaders(w)
	fmt.Fprint(w, "500 Internal Server Error")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/wish/web", wishHTMLHandler)
	mux.HandleFunc("/wish/text", wishTextHandler)

	mux.HandleFunc("/404", notFoundHandler)
	mux.HandleFunc("/500", internalServerErrorHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFoundHandler().ServeHTTP(w, r)
	})

	log.Printf("Server starting on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
