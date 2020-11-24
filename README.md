
O Cloud Build faz login no docker hub por isso é necessário criar uma Secret.

Exemplo do Google para criar uma Secret para armazenar a senha do Docker Hub:
https://cloud.google.com/cloud-build/docs/securing-builds/use-encrypted-secrets-credentials#store_credentials

O nome utilizado para armazenar a senha do docker é:
<b>docker-hub-password</b>

Docker Hub Image Name: https://hub.docker.com/r/robsantossilva/desafio-go-codeeducation