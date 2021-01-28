#Thomas' URL Shortener Microservice

## Quick Reference
1. [Building the URL shortener](#Building)
	1. [Locally](#BuildLocally)
	1. [Containerized](#BuildContainerized)
1. [Running the URL shortener](#Running)
	1. [Locally](#RunLocally)
	1. [Containerized](#RunContainerized)
1. [Using the Service](#Usage)

## Building the URL shortener <a name="Building"></a>

### Locally <a name="BuildLocally"></a>
The URL shortener can be built locally with ```make build```

Unit tests are run with ```make test```

### Containerized <a name="BuildContainerized"></a>
The service can be built into a container using [Docker](https://www.docker.com/):
```
docker build -t url-shortener .
```

## Running the URL shortener <a name="Running"></a>

### Locally <a name="RunLocally"></a>
The URL shortener is created in the ```dist``` directory: ```./dist/url-shortener```

### Containerized <a name="RunContainerized"></a>
A volume is recommended for persisting the database
```
docker volume create shortener-db
docker run -dp 8080:8080 -v shortener-db:/app/url-shortener/etc url-shortener
```

Alternative to starting the volume and container seperately, the two can be run together via docker compose:
```
docker-compose up
```

## Using the Service <a name="Usage"></a>
Once running, the service will listen on port ```8080```. URL's can be shortened by making calls to ```\shorten``` of the form ```\shorten?url=<url to shorten>```. E.g.
```
localhost:8080/shorten?url=http://www.example.com
```

The service will return the shortened URL. E.g.:
```
http://www.example.com can now be accessed on /0
```

Calls to the shortened URL will redirect to the full URL. E.g. ```localhost:8080/0``` will redirect to ```http://www.example.com```
