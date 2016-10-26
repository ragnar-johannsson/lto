## Limited Time Offer

LTO is a static file server with expiring temporary links. New links are generated on-demand with a configurable time-to-live. Additionally and optionally, if a link with a valid time token points to a file that doesn't exist, a signed temporary redirect to Amazon S3 is provided instead.


### Configuration

See [cloud-config.yml](./cloud-config.yml) for an example deployment. Adjust the environment variables in `/etc/default/lto` to reflect your setup. Alternatively, run `ragnarb/lto` directly with the environment variables passed in.

Supported environment variables:

Name | Usage | Default value | Required
--- | --- | --- | ---
REDIS_ADDR | Redis server address | redis:6379 |
REDIS_PASSWD | Redis password |  |
REDIS_DB | Redis database | 0 |
S3_ACCESS_KEY | S3 access key |  |
S3_SECRET_KEY | S3 secret key |  |
S3_REGION | S3 region |  |
S3_BUCKET | S3 bucket |  |
TOKEN_SIZE | Character length of the time token | 12 |
LISTEN | Interface and port to listen to | :3000 |
URL_TTL | TTL of generated URLs | 3600s | 
URL_SECRET | Shared secret for registering new URLs |  | YES
BASE_URL | Base URL for returned register URL response |  | YES
FILES_PATH | Path to files to serve |  | YES

Configure the S3 variables to enable redirecting to S3 for files not found locally.

### Usage

To get a new temporary link to an example_file with expiry as specified in with `URL_TTL` above:

```
$ curl "lto_host/example_file?register=$URL_SECRET"
{"url":"<BASE_URL>/example_file?token=<GENERATED_TOKEN>"}
$
```

### License

BSD 2-Clause. See the LICENSE file for details.
