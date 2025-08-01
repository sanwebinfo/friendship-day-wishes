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
	name = strings.TrimSpace(name)
	return strings.ReplaceAll(name, "-", " ")
}

func asciiArt(name string) string {
	cleanedName := cleanName(name)
	art := `
   _
 |  _|
 | |_
 |  _|
 |_|ANTASTIC FRIEND ★★★
	`
	quotes := []string{
		" Friendship is the compass\n that guides us\n through life's storm",
	}

	quote := quotes[len(name)%len(quotes)]

	return fmt.Sprintf("\n wishes@%s:~💚$%s\n%s", escapeText(cleanedName), art, quote)
}

func generateSlug(name string) string {
	replacements := map[string]string{
		"+":   " ",
		"%20": " ",
		"%25": "",
	}

	for old, new := range replacements {
		name = strings.ReplaceAll(name, old, new)
	}

	var result []rune
	previousWasHyphen := false

	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, unicode.ToLower(r))
			previousWasHyphen = false
		} else if r == ' ' || r == '-' {
			if !previousWasHyphen {
				result = append(result, '-')
				previousWasHyphen = true
			}
		} else {
			previousWasHyphen = false
		}
	}

	slug := string(result)
	slug = strings.Trim(slug, "-")
	return slug
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
	TextURL := fmt.Sprintf("%s/wish/text", baseURL)
	shareURL := fmt.Sprintf("%s/wish/web?name=%s", baseURL, slugText)

	setHTMLHeaders(w)
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en" prefix="og: https://ogp.me/ns#">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
    <link rel="icon" type="image/png" sizes="196x196" href="/favicon-196.png">

    <title>%s : Happy Friendship Wishes</title>
    <meta name="description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.">
    <meta name="canonical" href="%s">

    <meta property="og:site_name" content="%s : Happy Friendship Wishes">
    <meta property="og:type" content="website">
    <meta property="og:title" content="%s : Happy Friendship Wishes">
    <meta property="og:description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.">
    <meta property="og:url" content="%s">
    <meta property="og:image" content="https://img.sanweb.info/friend/friend?name=%s">
    <meta property="og:image:alt" content="%s : Happy Friendship Wishes">
    <meta property="og:image:width" content="1080">
    <meta property="og:image:height" content="1080">

    <meta name="twitter:title" content="%s : Happy Friendship Wishes">
    <meta name="twitter:description" content="Happy Friendship Day ASCII Text Greeting Art - Friendship Day Greeting Generator With Name.">
    <meta name="twitter:url" content="%s">
    <meta name="twitter:card" content="summary_large_image">
    <meta name="twitter:image" content="https://img.sanweb.info/friend/friend?name=%s">

    <link rel="preconnect" href="https://cdnjs.cloudflare.com">
    <link rel="preconnect" href="https://img.sanweb.info">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/1.0.4/css/bulma.min.css" integrity="sha512-yh2RE0wZCVZeysGiqTwDTO/dKelCbS9bP2L94UvOFtl/FKXcNAje3Y2oBg/ZMZ3LS1sicYk4dYVGtDex75fvvA==" crossorigin="anonymous" referrerpolicy="no-referrer" />
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/7.0.0/css/all.min.css" integrity="sha512-DxV+EoADOkOygM4IR9yXP8Sb2qwgidEmeqAEmDKIOfPRQZOWbXCzLC6vjbZyy0vPisbH2SyW27+ddLVCN+OMzQ==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <style>
        html, body {
            min-height: 100vh;
            margin: 0;
            padding: 0;
        }
        body {
            font-family: "Roboto Condensed", sans-serif;
            background-color: #58B19F;
            min-height: 100vh;
        }
        #quote-container {
            margin: 10px auto;
            padding: 20px;
            background-color: #D6A2E8;
            position: relative;
        }
        #quote {
            font-size: 20px;
            margin-bottom: 20px;
            color: #333;
        }
        #quote-card {
            background-color: #D6A2E8;
            border-radius: 15px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
            padding: 20px;
            margin-top: 20px;
        }
        pre {
            font-family: monospace;
            font-size: 14px;
            background-color: #3d3d3d;
            color: #ecf0f1;
            text-shadow: 0 0 3px #ecf0f1;
            padding: 20px;
            border-radius: 10px; 
            word-wrap: break-word;
            overflow-x: auto;
            line-height: inherit;
            position: relative;
        }
        .copy-icon {
            position: absolute;
            top: 5px;
            right: 5px;
            cursor: pointer;
            color: #ecf0f1;
        }
        .notification {
            font-family: "Roboto Condensed", sans-serif;
            display: none;
            position: fixed;
            top: 10px;
            right: 10px;
            z-index: 1000;
        }
        .notification.is-primary {
            background-color: #204269ff;
            color: #fff;
        }
        .form-container {
            font-family: "Roboto Condensed", sans-serif;
            margin: 20px auto;
            padding: 20px;
            background-color: #4b4b4b;
            border-radius: 15px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
            max-width: 500px;
        }
        .form-container .field {
            font-family: "Roboto Condensed", sans-serif;
            margin-bottom: 15px;
        }
        .form-container .input,
        .form-container .button {
            font-family: "Roboto Condensed", sans-serif;
            border-radius: 10px;
            width: 100%%;
        }
        .form-container .button {
            font-family: "Roboto Condensed", sans-serif;
            background-color: #25d366; 
            border-color: transparent;
            color: #fff;
        }
        .form-container .button:hover {
            background-color: #1ebd74;
        }
    </style>
</head>
<body>

<section class="section">
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
        <div class="buttons is-centered">
            <a class="button is-warning is-rounded" href="https://img.sanweb.info/dl/file?url=https://img.sanweb.info/friend/friend?name=%s" target="_blank" rel="nofollow noopener">
                <i class="fa fa-download" aria-hidden="true"></i>&nbsp;Download Image
            </a>
        </div>
        <pre id="ascii-art">
%s
<span class="icon copy-icon" onclick="copyToClipboard()">
    <i class="fas fa-copy"></i>
</span>
        </pre>
        <br>
        <pre>$ curl -G --data-urlencode "name=%s" %s<br><br>$ http -b GET "%s" "name==%s"</pre>
        <br>
        <div class="form-container">
            <h2 class="title is-4 has-text-centered has-text-light">Create Your Greeting</h2>
            <form action="/wish/web" method="get" onsubmit="sanitizeInput(event)">
                <div class="field">
                    <label class="label has-text-warning has-text-centered" for="name">Your Name</label>
                    <div class="control">
                        <input class="input" type="text" id="name" name="name" placeholder="Enter your name" minlength="2" maxlength="36" required>
                    </div>
                </div>
                <div class="field">
                    <div class="control">
                        <button class="button is-primary" type="submit">Generate Greeting</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</section>

<div class="notification is-primary" id="copy-notification">
    ✅ Copied to clipboard
</div>

<script>
    function copyToClipboard() {
        const asciiArt = document.getElementById('ascii-art').innerText;
        navigator.clipboard.writeText(asciiArt).then(() => {
            const notification = document.getElementById('copy-notification');
            notification.style.display = 'block';
            setTimeout(() => {
                hideNotification();
            }, 2000);
        }).catch(err => {
            console.error('Failed to copy text: ', err);
        });
    }

    function hideNotification() {
        const notification = document.getElementById('copy-notification');
        notification.style.display = 'none';
    }
    function sanitizeInput(event) {
        event.preventDefault();
        const form = event.target;
        const nameInput = form.querySelector('#name');
        const sanitizedValue = slugify(nameInput.value.trim());
        nameInput.value = sanitizedValue;
        form.submit();
    }
    function slugify(text) {
        return text.toString().toLowerCase()
            .replace(/\s+/g, '-')
            .replace(/[^\w\-]+/g, '')
            .replace(/\-\-+/g, '-')
            .replace(/^-+/, '')
            .replace(/-+$/, '');
    }
</script>

</body>
</html>
`, cleanName(name), shareURL, cleanName(name), cleanName(name), shareURL, slugText, cleanName(name), cleanName(name), shareURL, slugText, slugText, slugText, asciiText, cleanName(name), TextURL, TextURL, cleanName(name))
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

func homeHandler(w http.ResponseWriter, r *http.Request) {

	setHTMLHeaders(w)
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en" prefix="og: https://ogp.me/ns#">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
    <link rel="icon" type="image/png" sizes="196x196" href="/favicon-196.png">

    <title>Friendship Day Greeting Generator</title>
    <meta name="description" content="Create beautiful ASCII art greetings for your friends.">

    <meta property="og:site_name" content="Friendship Day Greeting Generator">
    <meta property="og:type" content="website">
    <meta property="og:title" content="Friendship Day Greeting Generator">
    <meta property="og:description" content="Create beautiful ASCII art greetings for your friends.">
    <meta property="og:image" content="https://img.sanweb.info/friend/friend?name=Your-Name">
    <meta property="og:image:alt" content="Happy Friendship Wishes">
    <meta property="og:image:width" content="1080">
    <meta property="og:image:height" content="1080">
    
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600;700&family=Dancing+Script:wght@700&display=swap" rel="stylesheet">

   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/7.0.0/css/all.min.css" integrity="sha512-DxV+EoADOkOygM4IR9yXP8Sb2qwgidEmeqAEmDKIOfPRQZOWbXCzLC6vjbZyy0vPisbH2SyW27+ddLVCN+OMzQ==" crossorigin="anonymous" referrerpolicy="no-referrer" />
    
    <style>
        :root {
            --primary-color: #58B19F;
            --secondary-color: #25d366;
            --dark-color: #2C3A47;
            --light-color: #f5f6fa;
            --accent-color: #FD7272;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Poppins', sans-serif;
            background-color: var(--primary-color);
            color: var(--dark-color);
            line-height: 1.6;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            padding: 1rem;
        }
        
        .container {
            max-width: 800px;
            margin: auto;
            background-color: white;
            padding: 2.5rem;
            border-radius: 16px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
            text-align: center;
            width: 100%%;
            animation: fadeIn 0.5s ease-out;
        }
        
        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }
        
        .logo {
            font-family: 'Dancing Script', cursive;
            font-size: 2.5rem;
            color: var(--accent-color);
            margin-bottom: 1rem;
        }
        
        h1 {
            font-family: 'Poppins', sans-serif;
            font-size: 2.2rem;
            font-weight: 700;
            color: var(--dark-color);
            margin-bottom: 1rem;
            line-height: 1.2;
        }
        
        .subtitle {
            font-family: 'Poppins', sans-serif;
            font-size: 1.1rem;
            color: #666;
            margin-bottom: 2rem;
            max-width: 600px;
            margin-left: auto;
            margin-right: auto;
        }
        
        .form-container {
            margin: 2rem auto;
            max-width: 500px;
        }
        
        .input-group {
            position: relative;
            margin-bottom: 1.5rem;
        }
        
        .input-icon {
            position: absolute;
            left: 1rem;
            top: 50%%;
            transform: translateY(-50%%);
            color: var(--primary-color);
        }
        
        input {
            width: 100%%;
            padding: 1rem 1rem 1rem 3rem;
            font-family: 'Poppins', sans-serif;
            font-size: 1rem;
            border: 2px solid #e0e0e0;
            border-radius: 8px;
            transition: all 0.3s ease;
            font-family: inherit;
        }
        
        input:focus {
            outline: none;
            border-color: var(--primary-color);
            box-shadow: 0 0 0 3px rgba(88, 177, 159, 0.2);
        }
        
        input::placeholder {
            font-family: 'Poppins', sans-serif;
            color: #aaa;
        }
        
        .btn {
            display: inline-block;
            background-color: var(--secondary-color);
            color: white;
            border: none;
            padding: 1rem 2.5rem;
            font-family: 'Poppins', sans-serif;
            font-size: 1rem;
            font-weight: 600;
            border-radius: 8px;
            cursor: pointer;
            transition: all 0.3s ease;
            text-transform: uppercase;
            letter-spacing: 1px;
            width: 100%%;
            max-width: 300px;
            margin-top: 1rem;
        }
        
        .btn:hover {
            background-color: #1ebd74;
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(37, 211, 102, 0.3);
        }
        
        .btn:active {
            transform: translateY(0);
        }
        
        .features {
            display: flex;
            flex-wrap: wrap;
            justify-content: center;
            gap: 1.5rem;
            margin-top: 3rem;
        }
        
        .feature {
            flex: 1 1 200px;
            padding: 1.5rem;
            background: rgba(255, 255, 255, 0.8);
            border-radius: 10px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
        }
        
        .feature i {
            font-size: 2rem;
            color: var(--primary-color);
            margin-bottom: 1rem;
        }
        
        .feature h3 {
            font-size: 1.1rem;
            margin-bottom: 0.5rem;
        }
        
        .feature p {
            font-family: 'Poppins', sans-serif;
            font-size: 0.9rem;
            color: #000;
        }
        
        footer {
            margin-top: 3rem;
            font-size: 0.9rem;
            color: black;
        }
        
        @media (max-width: 768px) {
            .container {
                padding: 1.5rem;
            }
            
            h1 {
                font-size: 1.8rem;
            }
            
            .logo {
                font-size: 2rem;
            }
            
            .features {
                flex-direction: column;
            }
        }
        button {
          font-family: 'Poppins', sans-serif;
          font-weight: 700;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">Friendship Day</div>
        <h1>Create Your Personalized Greeting</h1>
        <p class="subtitle">Generate beautiful ASCII art greetings to share with your friends and loved ones</p>
        
        <div class="form-container">
            <form action="/wish/web" method="get">
                <div class="input-group">
                    <i class="fas fa-user input-icon"></i>
                    <input type="text" name="name" placeholder="Enter your name" required 
                           minlength="2" maxlength="36" 
                           pattern="[A-Za-z ]+" 
                           title="Please enter only letters and spaces">
                </div>
                <button type="submit" class="btn">
                    <i class="fas fa-magic"></i> Create
                </button>
            </form>
        </div>
        
        <div class="features">
            <div class="feature">
                <i class="fas fa-paint-brush"></i>
                <h3>Beautiful Art</h3>
                <p>Stunning ASCII designs and wishing image with your name that impress your friends</p>
            </div>
            <div class="feature">
                <i class="fas fa-share-alt"></i>
                <h3>Easy Sharing</h3>
                <p>Share your creations via social media or messaging</p>
            </div>
            <div class="feature">
                <i class="fas fa-mobile-alt"></i>
                <h3>Mobile Friendly</h3>
                <p>Works perfectly on all devices</p>
            </div>
        </div>
        
        <footer>
            <p>Made with <i class="fas fa-heart" style="color: var(--accent-color);"></i> for Friendship Day</p>
        </footer>
    </div>
</body>
</html>
`)
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
	setHTMLHeaders(w)
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en" prefix="og: https://ogp.me/ns#">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
    <link rel="icon" type="image/png" sizes="196x196" href="/favicon-196.png">

    <title>Friendship Day Greeting Generator</title>
    <meta name="description" content="Create beautiful ASCII art greetings for your friends.">

    <meta property="og:site_name" content="Friendship Day Greeting Generator">
    <meta property="og:type" content="website">
    <meta property="og:title" content="Friendship Day Greeting Generator">
    <meta property="og:description" content="Create beautiful ASCII art greetings for your friends.">
    <meta property="og:image" content="https://img.sanweb.info/friend/friend?name=Your-Name">
    <meta property="og:image:alt" content="Happy Friendship Wishes">
    <meta property="og:image:width" content="1080">
    <meta property="og:image:height" content="1080">
    
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600;700&family=Dancing+Script:wght@700&display=swap" rel="stylesheet">
    
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/7.0.0/css/all.min.css" integrity="sha512-DxV+EoADOkOygM4IR9yXP8Sb2qwgidEmeqAEmDKIOfPRQZOWbXCzLC6vjbZyy0vPisbH2SyW27+ddLVCN+OMzQ==" crossorigin="anonymous" referrerpolicy="no-referrer" />
    
    <style>
        :root {
            --primary-color: #58B19F;
            --secondary-color: #25d366;
            --dark-color: #2C3A47;
            --light-color: #f5f6fa;
            --accent-color: #FD7272;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: 'Poppins', sans-serif;
            background-color: var(--primary-color);
            color: var(--dark-color);
            line-height: 1.6;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            padding: 1rem;
        }
        
        .container {
            max-width: 800px;
            margin: auto;
            background-color: white;
            padding: 2.5rem;
            border-radius: 16px;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
            text-align: center;
            width: 100%%;
            animation: fadeIn 0.5s ease-out;
        }
        
        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }
        
        .logo {
            font-family: 'Dancing Script', cursive;
            font-size: 2.5rem;
            color: var(--accent-color);
            margin-bottom: 1rem;
        }
        
        h1 {
            font-size: 2.2rem;
            font-weight: 700;
            color: var(--dark-color);
            margin-bottom: 1rem;
            line-height: 1.2;
        }
        
        .subtitle {
            font-size: 1.1rem;
            color: #666;
            margin-bottom: 2rem;
            max-width: 600px;
            margin-left: auto;
            margin-right: auto;
        }
        
        .form-container {
            margin: 2rem auto;
            max-width: 500px;
        }
        
        .input-group {
            position: relative;
            margin-bottom: 1.5rem;
        }
        
        .input-icon {
            position: absolute;
            left: 1rem;
            top: 50%%;
            transform: translateY(-50%%);
            color: var(--primary-color);
        }
        
        input {
            width: 100%%;
            padding: 1rem 1rem 1rem 3rem;
            font-size: 1rem;
            border: 2px solid #e0e0e0;
            border-radius: 8px;
            transition: all 0.3s ease;
            font-family: inherit;
        }
        
        input:focus {
            outline: none;
            border-color: var(--primary-color);
            box-shadow: 0 0 0 3px rgba(88, 177, 159, 0.2);
        }
        
        input::placeholder {
            color: #aaa;
        }
        
        .btn {
            display: inline-block;
            background-color: var(--secondary-color);
            color: white;
            border: none;
            padding: 1rem 2.5rem;
            font-size: 1rem;
            font-weight: 600;
            border-radius: 8px;
            cursor: pointer;
            transition: all 0.3s ease;
            text-transform: uppercase;
            letter-spacing: 1px;
            width: 100%%;
            max-width: 300px;
            margin-top: 1rem;
        }
        
        .btn:hover {
            background-color: #1ebd74;
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(37, 211, 102, 0.3);
        }
        
        .btn:active {
            transform: translateY(0);
        }
        
        .features {
            display: flex;
            flex-wrap: wrap;
            justify-content: center;
            gap: 1.5rem;
            margin-top: 3rem;
        }
        
        .feature {
            flex: 1 1 200px;
            padding: 1.5rem;
            background: rgba(255, 255, 255, 0.8);
            border-radius: 10px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
        }
        
        .feature i {
            font-size: 2rem;
            color: var(--primary-color);
            margin-bottom: 1rem;
        }
        
        .feature h3 {
            font-size: 1.1rem;
            margin-bottom: 0.5rem;
        }
        
        .feature p {
            font-size: 0.9rem;
            color: #666;
        }
        
        footer {
            margin-top: 3rem;
            font-size: 0.9rem;
            color: white;
        }
        
        @media (max-width: 768px) {
            .container {
                padding: 1.5rem;
            }
            
            h1 {
                font-size: 1.8rem;
            }
            
            .logo {
                font-size: 2rem;
            }
            
            .features {
                flex-direction: column;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">Friendship Day</div>
        <h1>Create Your Personalized Greeting</h1>
        <p class="subtitle">Generate beautiful ASCII art greetings to share with your friends and loved ones</p>
        
        <div class="form-container">
            <form action="/wish/web" method="get">
                <div class="input-group">
                    <i class="fas fa-user input-icon"></i>
                    <input type="text" name="name" placeholder="Enter your name" required 
                           minlength="2" maxlength="36" 
                           pattern="[A-Za-z ]+" 
                           title="Please enter only letters and spaces">
                </div>
                <button type="submit" class="btn">
                    <i class="fas fa-magic"></i> Create
                </button>
            </form>
        </div>
        
        <div class="features">
            <div class="feature">
                <i class="fas fa-paint-brush"></i>
                <h3>Beautiful Art</h3>
                <p>Stunning ASCII designs that impress your friends</p>
            </div>
            <div class="feature">
                <i class="fas fa-share-alt"></i>
                <h3>Easy Sharing</h3>
                <p>Share your creations via social media or messaging</p>
            </div>
            <div class="feature">
                <i class="fas fa-mobile-alt"></i>
                <h3>Mobile Friendly</h3>
                <p>Works perfectly on all devices</p>
            </div>
        </div>
        
        <footer>
            <p>Made with <i class="fas fa-heart" style="color: var(--accent-color);"></i> for Friendship Day</p>
        </footer>
    </div>
</body>
</html>
`)
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	setHTMLHeaders(w)
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>500 Internal Server Error</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            padding: 50px;
            background-color: #58B19F;
        }
        h1 {
            font-size: 50px;
            color: #fff;
        }
        p {
            font-size: 20px;
            color: #fff;
        }
        a {
            color: #fff;
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <h1>500</h1>
    <p>Internal Server Error</p>
    <p><a href="/">Go to Home Page</a></p>
</body>
</html>
`)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/wish/web", wishHTMLHandler)
	mux.HandleFunc("/wish/text", wishTextHandler)
	mux.HandleFunc("/404", notFoundHandler)
	mux.HandleFunc("/500", internalServerErrorHandler)
	mux.HandleFunc("/", homeHandler)

	log.Printf("Server starting on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
