package configs

import (
	"github.com/magiconair/properties/assert"
	"start-jushin/configs"
	"testing"
)

func Test_TickSystem_Validate(t *testing.T) {
	tests := []struct {
		name   string
		args   configs.TickSystem
		expect string
	}{
		{
			name: "Validate success",
			args: configs.TickSystem{
				S3Region:                    "ap-northeast-1",
				S3Bucket:                    "test-config-bp08",
				GroupDefinitionForFirstKei:  "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionForSecondKei: "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:             "tk_jushin_process_path",
			},
			expect: "",
		},
		{
			name: "TK_SYSTEM_FIRST_KEI_S3_PATH",
			args: configs.TickSystem{
				S3Region:                    "ap-northeast-1",
				S3Bucket:                    "test-config-bp08",
				GroupDefinitionForFirstKei:  "",
				GroupDefinitionForSecondKei: "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:             "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_FIRST_KEI_S3_PATH is required ",
		},
		{
			name: "TK_SYSTEM_SECOND_KEI_S3_PATH",
			args: configs.TickSystem{
				S3Region:                    "ap-northeast-1",
				S3Bucket:                    "test-config-bp08",
				GroupDefinitionForFirstKei:  "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionForSecondKei: "",
				InstancePathKey:             "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_SECOND_KEI_S3_PATH is required ",
		},
		{
			name: "TK_SYSTEM_INSTANCE_PATH_ENV_KEY",
			args: configs.TickSystem{
				S3Region:                    "ap-northeast-1",
				S3Bucket:                    "test-config-bp08",
				GroupDefinitionForFirstKei:  "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionForSecondKei: "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:             "",
			},
			expect: "TK_SYSTEM_INSTANCE_PATH_ENV_KEY is required ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.Validate()
			if tt.name != "Validate success" {
				assert.Equal(t, result.Error(), tt.expect)
			} else {
				assert.Equal(t, result, nil)
			}
		})
	}
}
