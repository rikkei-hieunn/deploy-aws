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
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei1Object:           "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "",
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI1_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionKehaiKei1Object: "",
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI1_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei2Object:  "",
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_QUOTE_CODES_DEFINITION_TICK_KEI2_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionKehaiKei2Object: "",
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_QUOTE_CODES_DEFINITION_KEHAI_KEI2_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei1Object:           "",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_GROUP_DEFINITION_KEI1_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei1Object:           "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionKei2Object:           "",
				InstancePathKey:                     "tk_jushin_process_path",
			},
			expect: "TK_SYSTEM_GROUP_DEFINITION_KEI2_OBJECT is required",
		},
		{
			name: "TK_SYSTEM_INSTANCE_PATH_ENV_KEY is empty",
			args: configs.TickSystem{
				QuoteCodesDefinitionTickKei1Object:  "environment_variables/quote_code_definition_tick_kei1.json",
				QuoteCodesDefinitionKehaiKei1Object: "environment_variables/quote_code_definition_kehai_kei1.json",
				QuoteCodesDefinitionTickKei2Object:  "environment_variables/quote_code_definition_tick_kei2.json",
				QuoteCodesDefinitionKehaiKei2Object: "environment_variables/quote_code_definition_kehai_kei2.json",
				S3Region:                            "ap-northeast-1",
				S3Bucket:                            "test-config-bp08",
				GroupDefinitionKei1Object:           "environment_variables/groups_definition_first_kei.json",
				GroupDefinitionKei2Object:           "environment_variables/groups_definition_second_kei.json",
				InstancePathKey:                     "",
			},
			expect: "TK_SYSTEM_INSTANCE_PATH_ENV_KEY is required",
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
