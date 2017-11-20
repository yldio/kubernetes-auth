package main

import (
	"html/template"
	"log"
	"net/http"
)

type tokenTmplData struct {
	IDToken      string
	RefreshToken string
	RedirectURL  string
	Claims       string
	Iss          string
	Aud          string
	Email        string
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
		<h1>Copy the following file to ~/.kube/config to login to the cluster</h1>
		<p> <pre><code>
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURoRENDQW15Z0F3SUJBZ0lRYzA1VkdhRVZMUXNreEtQZ2loQWNkakFOQmdrcWhraUc5dzBCQVFzRkFEQmMKTVFrd0J3WURWUVFHRXdBeENUQUhCZ05WQkFnVEFERUpNQWNHQTFVRUJ4TUFNUWt3QndZRFZRUVJFd0F4RVRBUApCZ05WQkFvVENHSnZiM1JyZFdKbE1Ra3dCd1lEVlFRTEV3QXhFREFPQmdOVkJBTVRCMnQxWW1VdFkyRXdIaGNOCk1UY3hNVEE1TVRRMU1qSXpXaGNOTWpBeE1UQTRNVFExTWpJeldqQmNNUWt3QndZRFZRUUdFd0F4Q1RBSEJnTlYKQkFnVEFERUpNQWNHQTFVRUJ4TUFNUWt3QndZRFZRUVJFd0F4RVRBUEJnTlZCQW9UQ0dKdmIzUnJkV0psTVFrdwpCd1lEVlFRTEV3QXhFREFPQmdOVkJBTVRCMnQxWW1VdFkyRXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCCkR3QXdnZ0VLQW9JQkFRQy9iU1dJTUluTitJU2tjRStDR0psMk5mN1lVaGNiS3QvMzc1S0QvdkMyY05QM3dsU2UKd1hFRDR0M2RGTzkwU25SSlZPRXd6RTBmUGVPcXpzKzVlOWY4QWxlb2g4Z3l1NlovZlZkZFVFY1dSY2Uyci85UAo5dzRLQTVFdGpEUVJ0bVV2Wk5tazZSVXVjQjJqMmt2NTgzcy8ydHVmbmJLTG00MG04WjhMcUJGaDBzdTcrYjVSCnd2S2pJamhRWHJtY1ZIU3FNeThmZnVQZ1NReVlJZDhXRlRTSTIzbG1RYXErKytLMVJJY1NRSmxxUWJRN1JJODMKNkxPbk42NHBaa0JXZ3dOd0d6NXZYbXdOakZyeXFIa3BMaUNndXYyei8raytNTk1TeVVJbngrYlJYaHo2UmFZYQpJNEwvT2g4cGVscWY5MEpNcDYvUTBXVGxnR3B5RFFHZlJibzFBZ01CQUFHalFqQkFNQTRHQTFVZER3RUIvd1FFCkF3SUNwREFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlE0ZmZqMkQxZzltRVpJM2FURTBCc1EKdHNIUE9qQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFYNmsvOU8wd3ZVdm53akNXTmgxYUhPQUx0bmx2bHVXRgpLSDJSckoxbkdtTkF2TmJ2aDBSMFd6VlVwa2NYMW5BaTYrMnJ3OXY4WkIzRFZwb1hTN0dENFJwKzlYai9udXFFClpKRUt1bGFJVHFVK0NTekhNejc0cXJxU3lvd3drL0cvOTRPNU1nMVE0cFBtcGIxdHZYdVM3VG9EOUpJSTh5NXoKUFFuRDhMKzVVejJoT1hDWk0wVFNIaXhRTW53TTM4YlRuM1YzMFFtRXpPVFhoNWxpbVArWVhlU0U4RXp6N1JvaApuREpjWVdYc1ZRQlY3UytKNUpvNWtseDNzck9jVUM1emx0Mm1ZcTNLRjBlK1hDYmpFSDJiZVA3eUM0bFV1d3E3ClQveHExNHNFUWNrcFRNSk5SRUZnakRKYklXTTlyWTNHQVdPYVBnU2MvdEp5VDV3NFNDOXFVQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://k8s-eu-api.sandbox.yld.io:443
  name: k8s-eu
contexts:
- context:
    cluster: k8s-eu
    user: {{ .Email }}
  name: ""
current-context: ""
kind: Config
preferences: {}
users:
- name: {{ .Email }}
  user:
    as-user-extra: {}
    auth-provider:
      config:
        client-id: k8s-auth
        client-secret: ZXhhbXBsZS1hcHAtc2VjcmV0
        id-token: {{ .IDToken }}
        idp-issuer-url: {{ .Iss }}
        refresh-token: {{ .RefreshToken }}
      name: oidc
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

func renderToken(w http.ResponseWriter, redirectURL, idToken, refreshToken string, claims []byte, iss string, aud string, email string) {
	renderTemplate(w, tokenTmpl, tokenTmplData{
		IDToken:      idToken,
		RefreshToken: refreshToken,
		RedirectURL:  redirectURL,
		Claims:       string(claims),
		Iss:          iss,
		Aud:          aud,
		Email:        email,
	})
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
