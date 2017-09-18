package ssh

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/gitlab-org/gitlab-runner/common"
	"gitlab.com/gitlab-org/gitlab-runner/executors"
	sshHelpers "gitlab.com/gitlab-org/gitlab-runner/helpers/ssh"
)

var (
	executorOptions = executors.ExecutorOptions{
		SharedBuildsDir: false,
		Shell: common.ShellScriptInfo{
			Shell:         "bash",
			Type:          common.NormalShell,
			RunnerCommand: "/usr/bin/gitlab-runner-helper",
		},
		ShowHostname: true,
	}
)

const SSH_SERVER_PRIVATE_KEY = "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIJCgNhsvCiKATDBXmRYHQfXIatKKOXGrmBLEVtGZtVv7oAoGCCqGSM49\nAwEHoUQDQgAE+36GvpVV34STGaV+YU4HHXCtJjburfo8IQDVTgLRwAkoLqLIl1cO\nduKDmdmeG/n66BNH1rJUkXFfEr4OYbZH5g==\n-----END EC PRIVATE KEY-----"

func TestPrepareSharedEnv(t *testing.T) {
	runnerConfig := &common.RunnerConfig{
		RunnerSettings: common.RunnerSettings{
			Executor: "ssh",
			SSH:      &sshHelpers.Config{User: "user", Password: "pass", Host: "127.0.0.1"},
		},
	}
	build := &common.Build{
		JobResponse: common.JobResponse{
			GitInfo: common.GitInfo{
				Sha: "1234567890",
			},
		},
		Runner: &common.RunnerConfig{},
	}

	sshConfig := runnerConfig.RunnerSettings.SSH
	server, err := NewMockServer(sshConfig.User, sshConfig.Password, []byte(SSH_SERVER_PRIVATE_KEY))
	assert.NoError(t, err)

	port, err := server.Start()
	assert.NoError(t, err)
	defer server.Stop()

	sshConfig.Port = strconv.Itoa(port)

	e := &executor{
		AbstractExecutor: executors.AbstractExecutor{
			ExecutorOptions: executorOptions,
		},
	}

	prepareOptions := common.ExecutorPrepareOptions{
		Config:  runnerConfig,
		Build:   build,
		Context: context.TODO(),
	}

	assert.False(t, build.SharedEnv)
	err = e.Prepare(prepareOptions)
	assert.NoError(t, err)
	assert.True(t, build.SharedEnv)
}
