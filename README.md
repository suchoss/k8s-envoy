# k8s-envoy
Simple go api to test Kubernetes settings.

It runs on port 8000.

## Api functions:

Everything is http GET request.

- **/** - returns simple hello world text
- **/envs** - returns all environmental variables in container
- **/readdir** - lists directory content. Example: `localhost:8000/readdir?q=/tmp`
- **/readfile** - reads file content. Example: `localhost:8000/readfile?q=/tmp/filename.txt`
- **/healthcheckok** - always returns status 200
- **/healthcheckfail** - for first 3 minutes returns status 200, after that it returns status 400


Maybe i might add also ping function to try ping other containers inside kubernetes.