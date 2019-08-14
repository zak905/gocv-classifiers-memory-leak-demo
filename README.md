## How to test gocv memory leak:

- `docker build . -t gocv-memory-leak-example`
- `docker run -p8080:8080 gocv-memory-leak-example`
- run `docker stats` in a different terminal or tab
- run ./requests.sh


Notice how the memory usage increases even if memory is freed after each request. Run another round of ./requests.sh and monitor again. about 100 MB is leaked with each round which is a lot, knowing that the applications start with only 7MB initially. 