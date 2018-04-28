# tempbot
A simple SlackBot that tells you the temperature and humidity using a DHT11 sensor attached to a Raspberry Pi.

## How to deploy
#### Assumptions
You are building the bot an a Raspberry Pi or similar ARM device. If not, you must make minor changes to the Dockerfile.
1. `$ docker build -t tempbot .`
2. `$ docker run --privileged -d -e SLACK_TOKEN=xoxb-your-slack-token tempbot`
* Add `--restart always` to the docker run command to make the container reboot after restarts
* The `--privileged` setting gives the container full access to the Pi's hardware, which may not be secure enough for your needs.
See https://stackoverflow.com/questions/30059784/docker-access-to-raspberry-pi-gpio-pins for more info.
