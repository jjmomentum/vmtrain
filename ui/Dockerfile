FROM node
MAINTAINER Barrett Jones, jonesb@vmware.com

# Install Bower & Grunt
RUN npm install -g bower
RUN npm install -g grunt-cli

# Install compass
RUN apt-get update -y \
  && apt-get install -y ruby ruby-dev build-essential \
  && gem install sass \
  && gem install compass

ENV HOME /home/app
WORKDIR /home/app

RUN useradd appuser -d /home/app -s /bin/bash -m

RUN chown -R appuser:appuser /home/app
USER appuser

# Bundle app source
COPY . /home/app

# Install app dependencies
RUN npm install
RUN bower install --config.interactive=false --allow-root

EXPOSE 9000
CMD grunt serve
