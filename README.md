# Al-Quran Indonesia

This repositories is for my collague task

You can direct see on web https://quran.reclaimyt.site

## Dependencys

- golang v1.22
- node v20.18 (lts)

## How to run?

### Build

First, you need to build backend and frontend to binary file.

```
make build
```

It will run:

- npm install
- npm run build
- go build -o main .

### Then run

```
./main [num_workers]
```