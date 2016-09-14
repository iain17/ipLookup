# IP Lookup API
A quick example showcasing how a simple and neat GoLang api could look like. A little deployment directory containing some supporting files:
- GeoIP2-City.mmdb: [MaxMind IP Datbase](https://www.maxmind.com).
- config.toml: A config file containing the port the api will bind to.

## How to run
- [Install Go](https://golang.org/doc/install) (Make sure you've correctly set-up the GOPATH etc.)
- Clone this repo and do run this command in your terminal: ``sh run.sh``

## Examples
- [IP lookup](http://localhost:8080/lookup/ip/81_2_69_142) (The underscores (_) should be set instead of the dots. Couldn't quickly figure out why the dots were not working)