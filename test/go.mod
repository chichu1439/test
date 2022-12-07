module git.sirius.io/banking/test

go 1.18

replace (
	git.multiverse.io/eventkit/common v0.0.20 => git.sirius.io/git.multiverse.io/eventkit/common.git v0.0.20
	git.multiverse.io/eventkit/kit v0.0.24 => git.sirius.io/git.multiverse.io/eventkit/kit.git v0.0.24
	git.multiverse.io/eventkit/sed v0.0.20 => git.sirius.io/git.multiverse.io/eventkit/sed.git v0.0.20
	git.multiverse.io/framework/common v0.0.19 => git.sirius.io/git.multiverse.io/framework/common.git v0.0.19
	git.multiverse.io/framework/log v0.0.11 => git.sirius.io/git.multiverse.io/framework/log.git v0.0.11
	git.sirius.io/banking/nfps/common v0.0.1 => ../../git.sirius.io/banking/nfps/common

)

require (
	git.multiverse.io/framework/common v0.0.19
	git.multiverse.io/framework/log v0.0.11
	github.com/Jeffail/tunny v0.1.4
	github.com/dongri/emv-qrcode v0.1.1
	github.com/ebfe/bcrypt_pbkdf v0.0.0-20140212075826-3c8d2dcb253a
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/ideazxy/iso8583 v0.0.0-20160317060925-d06dcb8f1fc4
	github.com/wumansgy/goEncrypt v0.0.0-20210730092718-e359121aa81e
	github.com/xormplus/xorm v0.0.0-20210512135344-8123d584d5f5
	github.com/yudaprama/iso20022 v0.0.0-20200629044733-52689e16a9cf
	github.com/zh-five/golimit v0.0.0-20200506070750-1320ef5af62a
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
)

require (
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
)

require (
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/Unknwon/goconfig v0.0.0-20200908083735-df7de6a44db8 // indirect
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/beego/i18n v0.0.0-20161101132742-e9308947f407 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.5 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	golang.org/x/sys v0.0.0-20210112080510-489259a85091 // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
