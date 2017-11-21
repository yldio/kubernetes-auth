COMMA := ,
SPACE := $(noop) $(noop)
POSTGRES_PASS = 66964843358242dbaaa7778d8477c288

FILES = $(wildcard *.yaml)
NAMESPACE = "authentication"

.PHONY: all db kube

kube:
	@kubectl apply --namespace $(NAMESPACE) -f $(subst $(SPACE),$(COMMA),$(FILES))

all: db kube

db:
	@helm install --set persistence.enabled=false --set postgresDatabase=dex_db --set postgresUser=dex --set postgresPassword=$(POSTGRES_PASS) --namespace $(NAMESPACE) --name authentication-db stable/postgresql
