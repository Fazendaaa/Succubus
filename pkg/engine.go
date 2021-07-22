package succubus

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
	// ctx := context.Background()
	// response, fail := docker.client.ContainerExecCreate(ctx,
	// 	docker.containerID,
	// 	types.ExecConfig{
	// 		AttachStderr: true,
	// 		AttachStdout: true,
	// 		Cmd:          command[0].commands,
	// 	},
	// )

	if nil != fail {
		return ok, fmt.Errorf("%w;\nerror while trying to execute given '%s' command", fail, command[0].commands)
	}

	return ok, fail
}

func (engine *Engine) execute(command []Commands) (ok bool, fail error) {
	return engine.docker.execute(command)
}
