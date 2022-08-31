package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"start-ecs/configs"
	"testing"
)

func Test_Init(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			path string
			name string
		}
		expect *configs.Server
		err    error
	}{
		{
			name: "parse environment variables file success",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables.json",
			},
			expect: &configs.Server{
				TickSystem: configs.TickSystem{
					BP03FirstRunningTypeEnvKeys:  []string{"TK_OPERATION_TYPE", "TK_TARGET_SYSTEM", "TK_KUBUN", "TK_HASSIN", "TK_CREATE_DATE", "TK_CREATE_TIME", "TK_START_INDEX"},
					BP03SecondRunningTypeEnvKeys: []string{"TK_OPERATION_TYPE", "TK_TARGET_SYSTEM", "TK_FOLDER_PATH"},
					BP05FirstRunningTypeEnvKeys:  []string{"TK_OPERATION_TYPE", "TK_DATE_REQUEST", "TK_TARGET_KEI"},
					BP05SecondRunningTypeEnvKeys: []string{"TK_OPERATION_TYPE", "TK_TYPE_NUMBER_FILE", "TK_DATE_REQUEST", "TK_TIME_REQUEST"},
					BP05ThirdRunningTypeEnvKeys:  []string{"TK_OPERATION_TYPE", "TK_DATE_REQUEST", "TK_FILE_LIST", "TK_TARGET_KEI"},
					BP05FourthRunningTypeEnvKeys: []string{"TK_OPERATION_TYPE", "TK_TYPE_NUMBER_FILE", "TK_DATE_REQUEST", "TK_TIME_REQUEST", "TK_FILE_LIST"},
					BP06EnvKeyParams:             []string{"id", "createdAt", "updatedAt", "version", "filesize", "name", "expiredAt", "categoryId", "userId", "path"},
					BP07FirstRunningTypeEnvKeys:  []string{"TK_BACKUP_DATE"},
					BP07SecondRunningTypeEnvKeys: []string{"TK_BACKUP_SOURCE_CLASS"},
					BP07ThirdRunningTypeEnvKeys:  []string{"TK_BACKUP_TABLES"},
					MaxWaitTime: 10000,
					RetryWaitTime: 5000,
					MaxCountTime: 3,
				},
			},
			err: nil,
		},
		{
			name: "parse environment variables file error file not found",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables1.json",
			},
			expect: nil,
			err:    viper.ConfigFileNotFoundError{},
		},
		{
			name: "parse environment variables file error wrong data format",
			args: struct {
				path string
				name string
			}{
				path: ".",
				name: "environment_variables_wrong_format.json",
			},
			expect: nil,
			err:    viper.ConfigParseError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := configs.Init(tt.args.path, tt.args.name)
			if err != nil {
				assert.Equal(t, fmt.Sprintf("%T", tt.err), fmt.Sprintf("%T", err))
			} else {
				assert.Equal(t, tt.expect.BP03FirstRunningTypeEnvKeys, result.BP03FirstRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP03SecondRunningTypeEnvKeys, result.BP03SecondRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP05FirstRunningTypeEnvKeys, result.BP05FirstRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP05SecondRunningTypeEnvKeys, result.BP05SecondRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP05ThirdRunningTypeEnvKeys, result.BP05ThirdRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP05FourthRunningTypeEnvKeys, result.BP05FourthRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP06EnvKeyParams, result.BP06EnvKeyParams)
				assert.Equal(t, tt.expect.BP07FirstRunningTypeEnvKeys, result.BP07FirstRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP07SecondRunningTypeEnvKeys, result.BP07SecondRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.BP07ThirdRunningTypeEnvKeys, result.BP07ThirdRunningTypeEnvKeys)
				assert.Equal(t, tt.expect.MaxWaitTime, result.MaxWaitTime)
				assert.Equal(t, tt.expect.RetryWaitTime, result.RetryWaitTime)
				assert.Equal(t, tt.expect.MaxCountTime, result.MaxCountTime)
			}
		})
	}
}
