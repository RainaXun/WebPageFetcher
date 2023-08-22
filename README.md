# WebPageFetcher

**WebPageFetcher** is a command line tool that fetches web pages and stores them locally.

## Installation
Run Dockerfile to build the docker image

```bash
docker build -t webpagefetcher .
```

## Start the program in the docker container

```bash
webpagefetcher
```

## After running the program, two modes will be supported:

### Mode 1: fetch content from multiple urls and save them into a local storage:

Commands:

```bash
./fetch https://www.google.com https://autify.com
```

./fetch should be the first argument, pointing to the directory of the local storage

the following arguments are urls, which will be fetched and saved into the local storage

### Mode 2: Output metadata including site url, number of links, number of images and last fetch time for given urls to the console:

```bash
./fetch --metadata https://www.google.com
```

./fetch should be the first argument, pointing to the directory of the local storage

--metadata should be the second argument, indicating the mode

the following arguments are urls, which will be fetched and analyzed, then output the metadata to the console

## Note
### if no http:// or https:// is provided, https:// will be added automatically
