# Plugin GO para Kong que adiciona um token JWT gerado no Keycloak para envio na requisição do backend.

1 - Obter token JWT no Keycloak
2 - Adicionar token no header do serviço

Uso: 

curl -s -X POST http://localhost:8001/services/3add879d-fb85-4315-9dac-82c8b3c8c01a/plugins \
	-d name=go-token-idp \
	-d config.URL_IDP=http://192.168.99.29:8180/auth/realms/realm/protocol/openid-connect/token \
	-d config.Client_id=backend-service \
	-d config.Client_secret=secret \
	-d config.Username=open-banking \
	-d config.Username=secret
