FROM golang:1.6

RUN apt-get update && apt-get install -y \
  libmagickwand-dev \
  libgraphicsmagick-dev \
  libgif-dev \
  imagemagick \
  graphicsmagick

ENV OMP_NUM_THREADS=1

COPY . ./src/imageconverter

RUN go get ./src/imageconverter && go install imageconverter

EXPOSE 5050
VOLUME /go/data

ENTRYPOINT ["imageconverter"]
