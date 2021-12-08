module mpldr.codes/recter

go 1.17

require (
	git.sr.ht/~poldi1405/glog v1.0.0
	github.com/fsnotify/fsnotify v1.5.1
	github.com/spf13/viper v1.9.0
	github.com/valyala/fasthttp v1.31.0
)

require github.com/aelsabbahy/goss v0.3.16 // for testing purposes

require (
	git.sr.ht/~poldi1405/go-ansi v1.4.1 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/achanda/go-sysctl v0.0.0-20160222034550-6be7678c45d2 // indirect
	github.com/aelsabbahy/GOnetstat v0.0.0-20160428114218-edf89f784e08 // indirect
	github.com/aelsabbahy/go-ps v0.0.0-20201009164808-61c449472dcf // indirect
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/miekg/dns v1.1.35 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/oleiade/reflections v0.0.0-20160817071559-0e86b3c98b2f // indirect
	github.com/onsi/gomega v1.10.4 // indirect
	github.com/opencontainers/runc v0.0.0-20161107232042-8779fa57eb4a // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/urfave/cli v0.0.0-20161102131801-d86a009f5e13 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/ini.v1 v1.63.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// internal packages
require (
	internal/data v1.0.0
	internal/handler v1.0.0
)

replace (
	internal/data => ./internal/data
	internal/handler => ./internal/handler
)
