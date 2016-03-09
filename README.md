# Image Converter

Converter images on fly.


## Clone

```bash
git clone git@github.com:c0va23/go-imageconverter.git
```


## Build

```bash
docker build -t imageconverter go-imageconverter
```


## Run

```bash
docker run --rm -it -p 7070:5050 -v $(pwd)/data:/go/data imageconverter --converter graphicsmagick
```
