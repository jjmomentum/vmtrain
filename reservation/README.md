Reservation Service
========

Reservation Service is a REST service to communicate between the Reservation UI and Lab Data service.

### Getting Started ###

## Build with Docker

Run `docker build -t reservation-service --rm=true .` to build docker image

Run `docker run -d -p 9292:9292 reservation-service` to start container

## Setup Locally

1.) Make sure you have ruby 2.2.1+ installed.

2.) Clone the repo to the desired directory

3.) Install gem dependencies

  * RHEL: `yum install -y libxml2 libxml2-devel libxslt-devel cmake`
  * Debian: `apt-get install -y libxml2 libxml2-dev libxslt-dev cmake`

4.) Install gems

  * `bundle install`

5.) You are now ready to host Reservation Service. This can be started from the 'Reservation Service' base directory by running:

  * `sh scripts/start_service.sh`
