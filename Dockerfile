FROM heroku/heroku:22-build as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN ls -l
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN ls /tmp/buildpack/heroku/go
RUN pwd
#RUN wget https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go
#RUN curl http://192.168.62.156:8000/go.tgz | tar xz -C /tmp/buildpack/heroku/go
RUN cp go.tgz /tmp/buildpack/heroku/go
RUN ls /tmp/buildpack/heroku/go
RUN cd /tmp/buildpack/heroku/go
RUN tar xzf go.tgz -C /tmp/buildpack/heroku/go
RUN ls -l /tmp/buildpack/heroku/go/*
RUN cd /app
#RUN tar -xfz go.tgz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-22 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare final, minimal image
FROM heroku/heroku:22

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
USER heroku
CMD /app/bin/go-getting-started