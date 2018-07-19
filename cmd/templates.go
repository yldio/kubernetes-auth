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
	Group         string
}

var tokenTmpl = template.Must(template.New("token.html").Parse(`<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="description" content="Kubernetes auth paired with dex, to provide a streamlined kubernetes cluster authentication workflow">
		<title>Kubernetes Auth | Login to your Kubernetes clusters.</title>
		<link rel="icon" href="https://kubernetes.io/images/favicon.png">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.0/normalize.min.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/milligram/1.3.0/milligram.min.css">
		<link rel="stylesheet" href="https://milligram.io/styles/main.css">
	</head>
	<body>
		<main class="wrapper">
			<section class="container" id="dependencies">
				<h4>Congratulations!</h4>
				<p>You have successfully authenticated with your authentication provider to enable access to your Kubernetes cluster.</p>
				<p><strong>Run the following command to set your credentials for this environment.</strong></p>
				<pre><code>kubectl config set-cluster {{ .Cluster }} \
    --server={{ .ApiServer }} \
    --insecure-skip-tls-verify=true
kubectl config set-context {{ .Cluster }} \
    --cluster={{ .Cluster }} \
    --user={{ .Email }}-{{ .Cluster }}
kubectl config set-credentials {{ .Email }}-{{ .Cluster }} --auth-provider=oidc \
    --auth-provider-arg=client-id={{ .KclientID }} \
    --auth-provider-arg=client-secret={{ .KclientSecret }} \
    --auth-provider-arg=id-token={{ .IDToken }} \
    --auth-provider-arg=idp-issuer-url={{ .Iss }} \
    --auth-provider-arg=refresh-token={{ .RefreshToken }}
kubectl config use-context {{ .Cluster }}</code></pre>
				<p>If this is your first time connecting to this environment, use the following to setup your default namespace.</p>
				<pre><code>kubectl config set-context $(kubectl config current-context) --namespace=&lt;a namespace&gt;</code></pre> 
				<p>To confirm everything is working as expected, and that you can access this cluster, please test by running the command <code>kubectl get pods</code>.</p>
				<details>
					<summary>Debug Information</summary>
					<p></p>
					<p>Refresh Token: <code>{{ .RefreshToken }}</code></p>
					<form action="{{ .RedirectURL }}" method="post">
						<input type="hidden" name="refresh_token" value="{{ .RefreshToken }}">
						<input type="submit" value="Redeem refresh token">
					</form>
					<p>For any issues send the following code snippet to your Kubernetes administrators.</p>
					<pre><code>{{ .Claims }}</code></pre>
				</details>
			</section>
		</main>
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
		Group:         "default",
	})
}

var indexTmpl = template.Must(template.New("index.html").Parse(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="description" content="Kubernetes auth paired with dex, to provide a streamlined kubernetes cluster authentication workflow">
		<title>Kubernetes Auth | Login to your Kubernetes clusters.</title>
		<link rel="icon" href="https://kubernetes.io/images/favicon.png">
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.0/normalize.min.css">
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/milligram/1.3.0/milligram.min.css">
		<link rel="stylesheet" href="https://milligram.io/styles/main.css">
	</head>
	<body>
		<main class="wrapper">
			<header class="header" id="home">
				<section class="container">
					<h1 class="title">Kubernetes Auth</h1>
					<p class="description">Welcome to our Kubernetes cluster <br><i><small>Currently v1.10.3</small></i></p>
					<form action="/login" method="post">
						<fieldset>
							<input class="button button-black" type="submit" title="Sign In" value="Sign In" style="background-color: #326ce5; border-color: #326ce5">
						</fieldset>
					</form>
				</section>
			</header>
			<section class="container" id="dependencies">
				<h3>Dependencies</h3>
				<p>You will need to install <code>kubectl</code> before continuing, and to ensure that you keep your own version upto date. Failure to do so could result in interrupted Kubernetes operational usage.</p>
				<p><strong>macOS</strong></p>
				<pre class=""><code class="">brew install kubernetes-cli</code></pre>
				<p><strong>linux</strong></p>
				<pre class=""><code class="">curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl</code></pre>
				<p><strong>windows</strong></p>
				<pre class=""><code class="">choco install kubernetes-cli</code></pre>
			</section>
			<section class="container" id="how">
				<h3>How does this work</h3>
				<p>We will authenticate you with the configured provider that has been setup, and then provide a set of commands to run using <code>kubectl</code>.</p>
				<p>You might be prompted by your authentication provider to provide limited access to your email and name, this application does not store or transfer this information, instead it is a front-end to a Oauth like workflow, which integrates with <a href="https://github.com/coreos/dex">Dex</a>.</p>
			</section>
		</main>
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
