# mongo-crud-go

proyecto golang creacion de pruebas unitarias para testeo del crud en mongodb

## Configuracion

En la carpeta docker se encuentra los archivos

.env => contiene la configuracion para autenticar en la base de datos de la aplicacion cargada en actions.go y de la imagen docker que contiene la base de datos mongo y mongo express

mongo-docker-compose.yml => definicion de las imagenes de docker para la base de datos de mongo y mongo express

## Docker

El proyecto incorpora una imagen de mongodb y mongo express para correr simulando el ambiente local y dev:  
#### Using Docker to simplify development (optional)
```
docker-compose -f docker/mongo-docker-compose.yml up -d
```
Para detenerlo y quitar del contenedor, ejecute:
```
docker-compose -f docker/mongo-docker-compose.yml down
```