FROM node:6

ENV workdir /usr/src/app

RUN npm install -g --quiet webpack@^3.8.x
RUN npm install -g --quiet grunt-cli@^1.0.x

RUN mkdir -p $app_dir

WORKDIR $workdir

ADD ./package.json $workdir/.

RUN npm install --quiet

CMD tail -f /dev/null