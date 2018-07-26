# kubernetes-auth

An authentication front-end to Kubernetes clusters, enabling users to log into
a Kubernetes cluster through the configuration and use of [Dex](https://github.com/coreos/dex),
[OIDC](https://github.com/coreos/dex/blob/5e34f0d1a6e22725b39f521178baac2cddd0a306/Documentation/openid-connect.md) 
and [Kubernetes OIDC](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens)

This has been developed to enable ease of use of initial developer setup and to
provide an easy way to switch between environments / regions in a non-federated
deployments.

It also provide an easy method to switch out Dex Connectors, so when your team
ends up moving from Github Teams to Okta, you have a minimal set of changes to
implement.

## What does it look like?

Login Page            | Command Page
:-------------------------:|:-------------------------:
![Login page](./images/login.png) | ![Command page](./images/commands.png)

## I have a cluster, lets go!

```bash
kubectl apply -f infrastructure/
```

## Preamble

First we need to have a primer on how this all fits togther, this repo is an application that sits between [Dex]() and Kubernetes using [OIDC tokens](https://kubernetes.io/docs/admin/authentication/#openid-connect-tokens).

The outcome after configuring will be asociated RBAC groups inside of the Kubernetes cluster based on the Github organisation and the teams which the member belongs to.

For a Github Organisation as such:

### Yldio
#### Teams
- devops
- software-engineering

It becomes possible to map the team `devops` with the Kuerbetes RBAC ClusterRole `cluster-admin` to give anyone in the team `devops` cluster-wide access to the kubernetes cluster. As such if we gave the team `software-engineering` a Role to a specific namespace, any new members will have access to the kubernetes cluster in that specific namespace.  
  

## Dex? Aka Kubernetes Authentication

![Workflow Image...](https://d33wubrfki0l68.cloudfront.net/d65bee40cabcf886c89d1015334555540d38f12e/c6a46/images/docs/admin/k8s_oidc_login.svg)

Dex acts as an intermediary between Github authentication and Kubernetes acting
as an identity provider. This gives us the flexibility to move to another backed
(LDAP, SAML, etc.) at some point in the future.

At the moment user logins are federated by github teams. Each team then belongs
to a namespace with view on everything in that namespace. As time progresses we
might want to restrict / expand on this.

To Login a user will use the following flow, with sandbox being replaced by their
environment of choice (levels of access will be handled):

- Navigate to http://k8s-auth.sandbox.yld.io
- Login to Github and authorise the YLD github application
- Follow the instructions and Copy the kubeconfig to your local ~/.kube/config
- check access with `kubectl get pods`

## Configuring Kubernetes

Once the application has been deployed an running, the next step is to point 
Kubernetes' OIDC options in the Kube API server.

```
--oidc-issuer-url=https://dex.sandbox.yld.io
--oidc-client-id=kubernetes-auth
--oidc-ca-file=/etc/kubernetes/ssl/openid-ca.pem
--oidc-username-claim=email
--oidc-groups-claim=groups
```

---

# Development

To enable development, it is required to run `dex` locally, so that `kubernetes-auth`
can resolve and connect to it.

```bash
echo $(minikube ip) cluster-auth.minikube.local | sudo tee -a /etc/hosts
minikube ssh 'echo 127.0.2.1 cluster-auth.minikube.local | sudo tee -a /etc/hosts'
helm upgrade --install dex ./infrastructure/dex --set secrets.github.client.id=abcdef --set secrets.github.client.secret=abcedf
kubectl apply -f infrastructure/dex/minikube.yaml
```
