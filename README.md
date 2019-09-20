# Cloud Platform Custom Error Pages

This repo creates a docker image for sample application with the HTTP handler implemented. It accepts a URL param and returns the HTTP status.

## Build and Push Image

1. Edit the `makefile` and change the version number of the `IMAGE` line near the top of the file.
2. Run `make all` to create a new version of the docker image, with your custom error pages.
3. Run `make push` to tag and push the image to docker hub.

## Usage

This is to use for Integration tests for Nginx Ingress Controller custom error pages. Change the error code for the host you want to get HTTP status. From the example below you can get a response for error code "503"

https://#{host}/err?code=503
