FROM ruby:2.2.1
MAINTAINER Barrett Jones <jonesb@vmware.com>

RUN apt-get update && \
    apt-get install -y build-essential libxml2 libxml2-dev libxslt-dev cmake

ENV HOME /home/reservation-service
WORKDIR /home/reservation-service

# Bundle app source
COPY . /home/reservation-service

# Install gems
RUN bundle install

# Start server
EXPOSE 9292
CMD bundle exec rackup -i. -p 9292 -o 0.0.0.0
