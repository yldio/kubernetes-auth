package main

import (
	"html/template"
	"log"
	"net/http"
)

type tokenTmplData struct {
	IDToken       string
	RefreshToken  string
	RedirectURL   string
	Claims        string
	Iss           string
	Aud           string
	Email         string
	Cluster       string
	ApiServer     string
	KclientID     string
	KclientSecret string
}

var tokenTmpl = template.Must(template.New("token.html").Parse(`<html>
  <head>
    <style>
/* make pre wrap */
pre {
 white-space: pre-wrap;       /* css-3 */
 white-space: -moz-pre-wrap;  /* Mozilla, since 1999 */
 white-space: -pre-wrap;      /* Opera 4-6 */
 white-space: -o-pre-wrap;    /* Opera 7 */
 word-wrap: break-word;       /* Internet Explorer 5.5+ */
}
    </style>
  </head>
  <body>
		<h1>Run the following command to set your credentials for this environment</h1>
		<p> <pre><code>
		$ kubectl config set-cluster {{ .Cluster }} \
				--server={{ .ApiServer }} \
				--insecure-skip-tls-verify=true

		$ kubectl config set-context {{ .Cluster }} \
				--cluster={{ .Cluster }} \
				--namespace={{ .Group }} \
				--user={{ .Email }}-{{ .Cluster }}

		$ kubectl config set-credentials {{ .Email }}-{{ .Cluster }} --auth-provider=oidc \
				--auth-provider-arg=client-id={{ .KclientID }} \
				--auth-provider-arg=client-secretr={{ .KclientSecret }} \
				--auth-provider-arg=id-token={{ .IDToken }} \
				--auth-provider-arg=idp-issuer-url={{ .Iss }} \
				--auth-provider-arg=refresh-token={{ .RefreshToken }}
		</code></pre> </p>
		<p>Test the config is working by running the following command:</p>
		<p> <pre><code>kubectl get nodes</code></pre> </p>
    <p> Refresh Token: <pre><code>{{ .RefreshToken }}</code></pre></p>
        <form action="{{ .RedirectURL }}" method="post">
          <input type="hidden" name="refresh_token" value="{{ .RefreshToken }}">
          <input type="submit" value="Redeem refresh token">
    </form>
		<h2>For any issues send the following to devops</h1>
    <p> <pre><code>{{ .Claims }}</code></pre></p>
  </body>
</html>
`))

func renderToken(w http.ResponseWriter, redirectURL, idToken, refreshToken string,
	claims []byte, iss string, aud string, email string, cluster string, apiServer string, kclientId string, kclientSecret string) {

	renderTemplate(w, tokenTmpl, tokenTmplData{
		IDToken:       idToken,
		RefreshToken:  refreshToken,
		RedirectURL:   redirectURL,
		Claims:        string(claims),
		Iss:           iss,
		Aud:           aud,
		Email:         email,
		Cluster:       cluster,
		ApiServer:     apiServer,
		KclientID:     kclientId,
		KclientSecret: kclientSecret,
	})
}

var indexTmpl = template.Must(template.New("index.html").Parse(`<html>
  <body>
    <form action="/login" method="post">
       <input type="submit" value="Login">
    </form>
  </body>
</html>`))

func renderIndex(w http.ResponseWriter) {
	renderTemplate(w, indexTmpl, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err == nil {
		return
	}

	switch err := err.(type) {
	case *template.Error:
		// An ExecError guarantees that Execute has not written to the underlying reader.
		log.Printf("Error rendering template %s: %s", tmpl.Name(), err)

		// TODO(ericchiang): replace with better internal server error.
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	default:
		// An error with the underlying write, such as the connection being
		// dropped. Ignore for now.
	}
}
