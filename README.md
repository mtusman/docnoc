# docnoc 
Docnoc is a Go CLI that scans docker containers and alerts users automatically when resources fail to meet predefined conditions via Slack and performs actions to ensure containers are in a specified state.

<img width="1106" alt="Screenshot 2019-09-29 at 13 03 06" src="https://user-images.githubusercontent.com/25107174/65832327-edebee00-e2b9-11e9-80ba-15c1223ecb77.png">

## Getting Started
Start by installing the CLI:
```go
# Clone outside of GOPATH
git clone https://github.com/mtusman/docnoc
cd docnoc
# Build and install
go install
# Run
docnoc
````

Define a docnoc config yaml file. `default` will run your predefined conditions on every docker container running on your machine, but you can prevent checks on certain containers by including them in the `exlude` list. To put specific checks on specific docker containers, define them in `containers` as a list, using the name of the container. Here is an example of a docnoc config yaml file:
```yaml
docnoc:
    default:
        cpu:
            min: 0
            max: 180
        memory:
            max: 90
        block_write:
            max: 100
        action: stop
    containers:
        web:
            cpu: 
                min: 30
                max: 200
    exclude:
        - nginx
    slack_webhook: https://hooks.slack.com/services/URL
```

Once you done that, you should be able to run docnoc via the command `docnoc -f docnoc_config.yaml`

Checks can be placed on `cpu`, `memory`, `block_write`, `block_read`, `network_rx` and `network_tx` and are always of type min and max.

### Use the mtusman/docnoc image
Alternatively, you could use the `mtusman/docnoc` image. Here you'll define a docnoc config yaml file and use the image to launch a container that mounts your docnoc config file to `/tmp/docnoc_config.yaml`. Below is an example of the how you can define the container as a `docker_compose.yaml` service.
```yaml
daemon:
  image: mtusman/docnoc
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock
    - ./examples/docnoc_config.yaml:/tmp/docnoc_config.yaml
```
