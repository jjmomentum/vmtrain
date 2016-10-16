# app

## Build & development

Run `grunt` for building and `grunt serve` for preview.

## Testing

Running `grunt test` will run the unit tests with karma.

## Build with Docker

Run `docker build -t reservation-ui --rm=true .` to build docker image

Run `docker run -d -p 9000:9000 reservation-ui` to start UI container

You can then access the UI at http://localhost:9000

This project is generated with [yo angular generator](https://github.com/yeoman/generator-angular)
version 0.15.1.
