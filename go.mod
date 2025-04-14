module github.com/philwelz/aksgpt

go 1.24.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.9.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice v1.0.0
	github.com/fatih/color v1.18.0
	github.com/openai/openai-go v0.1.0-beta.9
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.18.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.1 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.4.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

replace github.com/philwelz/aksgpt/pkg/azure => ./aksgpt/pkg/azure

replace github.com/philwelz/aksgpt/pkg/ai => ./aksgpt/pkg/ai

replace github.com/philwelz/aksgpt/pkg/cluster => ./aksgpt/pkg/cluster

replace github.com/philwelz/aksgpt/pkg/instructions => ./aksgpt/pkg/instructions
