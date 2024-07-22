module github.com/ShouheiNishi/openapi5g

go 1.18

require (
	github.com/deepmap/oapi-codegen/v2 v2.0.1-0.20240123090344-d326c01d279a
	github.com/free5gc/openapi v1.0.8
	github.com/getkin/kin-openapi v0.122.0
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.6.0
	github.com/invopop/yaml v0.3.1
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/oapi-codegen/runtime v1.1.1
	github.com/samber/lo v1.46.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/bytedance/sonic v1.10.0-rc3 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/deepmap/oapi-codegen/v2 v2.0.1-0.20240123090344-d326c01d279a => github.com/ShouheiNishi/oapi-codegen/v2 v2.0.0-20240130024248-ea21b85ac0b4
	github.com/getkin/kin-openapi v0.122.0 => github.com/ShouheiNishi/kin-openapi v0.120.1-0.20240130043656-caf78162e0eb
)
