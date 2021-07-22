package succubus

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	client      *client.Client
	containerID string
}

type LXC struct {
}

type HyperV struct {
}

type WindowsContainers struct {
}

type RKT struct {
}

type Engine struct {
	kind   string
	docker Docker
}

func (docker *Docker) start() (ok bool, fail error) {
	docker.client, fail = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if nil != fail {
		return false, fmt.Errorf("%w;\ncould not start Docker connection", fail)
	}

	return ok, fail
}

func (engine *Engine) start(kind string) (ok bool, fail error) {
	if "docker" != kind {
		return ok, fmt.Errorf("engine '%s' is not supported yet", kind)
	}

	engine.kind = kind
	engine.docker = Docker{}

	return engine.docker.start()
}

func (docker *Docker) registryImage(image Image) (ok bool, fail error) {
	ctx := context.Background()
	reader, fail := docker.client.ImagePull(ctx, "docker.io/library/go", types.ImagePullOptions{})

	if nil != fail {
		return false, fmt.Errorf("%w;\nerror while fetching docker '%s' image", fail, image.name)
	}

	io.Copy(os.Stdout, reader)

	return ok, fail
}

func (engine *Engine) registryImage(image Image) (ok bool, fail error) {
	return engine.docker.registryImage(image)
}

func (docker *Docker) execute(command []Commands) (ok bool, fail error) {
	ctx := context.Background()
	response, fail := docker.client.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}, nil, nil, nil, "")

	if nil != fail {
		return ok, fmt.Errorf("%w;\nerror while trying to create a container to run given '%s' command", fail, command[0].commands)
	}

	if fail = docker.client.ContainerStart(ctx, response.ID, types.ContainerStartOptions{}); fail != nil {
		return ok, fmt.Errorf("%w;\nerror while trying to execute given '%s' command", fail, command[0].commands)
	}

	statusCh, errCh := docker.client.ContainerWait(ctx, response.ID, container.WaitConditionNotRunning)
	select {
	case fail = <-errCh:
		if fail != nil {
			return ok, fmt.Errorf("%w;\nerror while trying to create a container to run given '%s' command", fail, command[0].commands)
		}
	case <-statusCh:
	}

	out, fail := docker.client.ContainerLogs(ctx, response.ID, types.ContainerLogsOptions{ShowStdout: true})
	if fail != nil {
		return ok, fmt.Errorf("%w;\nerror while trying to create a container to run given '%s' command", fail, command[0].commands)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return ok, fail
}

func (engine *Engine) execute(command []Commands) (ok bool, fail error) {
	return engine.docker.execute(command)
}
