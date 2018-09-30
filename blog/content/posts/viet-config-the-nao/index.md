---
title: "Viết config thế nào?"
date: 2018-09-30T17:17:03+07:00
draft: false
description: Tuần rồi mình được 2 người, ở 2 nơi khác nhau hỏi về chuyện viết config cho một ứng dụng, mình chợt nhận ra rằng việc tổ chức config cho một ứng dụng nhiều khi không đơn giản như chúng ta tưởng tượng, cái chính là không phải là cái config đó chạy được không, mà cơ bản là làm đúng hay không nữa.
tags: [config, go, viper, cobra]
images: ["/posts/viet-config-the-nao/images/config.jpeg"]
---

Tuần rồi mình được 2 người, ở 2 nơi khác nhau hỏi về chuyện viết config cho một ứng dụng, mình chợt nhận ra rằng việc tổ chức config cho một ứng dụng nhiều khi không đơn giản như chúng ta tưởng tượng, cái chính là không phải là cái config đó chạy được không, mà cơ bản là làm đúng hay không nữa.

## The twelve-factor app

Về cơ bản thì đây là một methodology để build software-as-a-service, lý do tại sao xài thì mọi người có thể đọc [**tại đây**](https://12factor.net/), mà bài này chỉ nói về config nên chúng ta chỉ quan tâm đến section 3 của nó:

> **III. Config**

> Store config in the environment

Ý tưởng cơ bản thì bạn sẽ lưu config tại biến môi trường (env), đại loại như `DATABASE_HOSTNAME`, một khi bạn thực hiện điều này nó sẽ giúp bạn có một số lợi ích sau đây:

* App của bạn có thể thay đổi config một cách dễ dàng, sửa env và restart, bạn không cần quan tâm nhiều điều config ở đâu, khai báo thế nào.
* App buộc phải restart lại mới load config mới, tất nhiên là vậy trừ phi ngay trong chính process của bạn tự override lại env
* Các deploy management hiện tại đều hỗ trợ config bằng env như một phương thức mặc định. AWS, Heroku, K8s...
* Config được stored từ các nơi khác dễ dàng có thể manipulate thành env chỉ với vài lệnh shell script (bạn có thể store config trong etcd, kms hoặc các secret vault và transform nó thành env)
* Env có thể bundle trong deployment thành các revision, rollback trong vòng 1 nốt nhạc.

## Implements

```
databaseHostname := os.Getenv("DATABASE_HOSTNAME")

if databaseHostname == "" {
    databaseHostname = defaultDatabaseHostname
}
```

Oh yeah, trông có vẻ đúng tinh thần của 12-factor đó, tuy nhiên cách này không có maintainable và extendsible. Thay vào đó đối với Go bạn có thể dùng 2 package là viper và cobra để bắt đầu, có người build sẵn rồi thì ta cứ xài thôi ahjhj.

#### Tạo một cái app

Giả sử tạo một application có cấu trúc như sau

```
/workspace
└── src
    └── duythinht
        └── 5kbps.io
            ├── cmd
            │   ├── config.go
            │   ├── config_test.go
            │   └── main.go
            ├── docs.go
            └── let_it.go
```

Chúng ta chỉ quan tâm đến nội dung trong `cmd` thôi, còn những cái còn lại tùy bạn. Đầu tiên chúng ta cần khai báo một struct cho config đã, lấy ví dụ config của chúng ta đơn giản chỉ có thông tin DB

```
// config.go
package main

var config struct {
    Database struct {
        Hostname string
        Username string
        Password string
    }
}
```

Viết một cái test đơn giản cho config

```
// config_test.go
package main

import "testing"

func TestDefaultConfigValue(t *testing.T) {
	initConfig()
	if config.Database.Hostname != "localhost" {
		t.Fatalf("Host must be 'localhost', got %s", config.Database.Hostname)
	}

	if config.Database.Username != "root" {
		t.Fatalf("Username must be 'root', got %s", config.Database.Username)
	}

	if config.Database.Password != "khongbiet" {
		t.Fatalf("Password must be 'khongbiet', got %s", config.Database.Password)
	}
}
```

#### Working with default config

Bước đầu tiên phải implements cho initConfig nó khởi tạo default config đã


```
// config.go

// default config in buffer
var configDefault = []byte(`
[database]
	hostname = "localhost"
	username = "root"
	password = "khongbiet"
`)

func initConfig() {
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(configDefault))
	viper.Unmarshal(&config)
}
```

Trong trường hợp của mình thì mình thích dùng `toml` format để định nghĩa config content, các bạn có thể dùng yaml, json hay props cũng được. Một số bạn sẽ hỏi là tại sao mình lại để config trong buffer mà không tách ra một file riêng, thật ra lý do đơn giản là vì khi distribute mình chỉ muốn ship 1 file binary duy nhất thôi, thay vì phải ship thêm một file config nữa, ngoài ra cái buffer sẽ có tác dụng về sau.

#### Load config từ environment var

Đoạn code trên chạy `go test` cơ bản là sẽ pass, nhưng mình muốn load config từ env cơ mà.

```
func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("DATABASE_HOSTNAME", "5kbps.io")
	initConfig()

	if config.Database.Hostname != "5kbps.io" {
		t.Fatalf("Host must be '5kbps.io', got %s", config.Database.Hostname)
	}
}
```

Và thêm `AutomaticEnv` vào initConfig

```
func initConfig() {
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(configDefault))

	// Load config from env
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.Unmarshal(&config)
}
```

Với các service cơ bản thì bạn chỉ cần dừng ở đây là chạy ngon lành cành đào rồi. Tuy nhiên nếu có yêu cầu phức tạp hơn có thể đọc tiếp.

#### Load config từ third party

3rd-party ở đây là chỉ đến dùng file để config hoặc dùng etcd, aws ssm, gcp params store... Ở đây mình chỉ ví dụ là dùng file, các trường hợp khác bạn có thể tự nghiên cứu implements

```
// main.go

package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

func main() {

	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Use:   "5kbps",
		Short: "5kbps is a very slow connection",
		Long:  `A very fast blog for slow connection built with love by duythinht and friends.
            Donate me for a food at https://5kbps.io`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Config is %+v\n", config)
		},
	}

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Path to config file")

	generateConfigCmd := &cobra.Command{
		Use:   "generate-config",
		Short: "Generate default config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s", bytes.TrimSpace(configDefault))
		},
	}

	rootCmd.AddCommand(generateConfigCmd)

	rootCmd.Execute()
}
```

Như mình đã nói, cái `defaultConfig` bây giờ mình có thể dùng để generate default config file(tránh mấy thanh niên tự tạo ra bị sai)
Và sửa lại hàm `initConfig` để load config từ file nếu flags config được spec

```
func initConfig() {
	viper.SetConfigType("toml")

	viper.ReadConfig(bytes.NewBuffer(configDefault))

	// Load config from file if posible
	if configFile != "" {
		viper.SetConfigFile(configFile)
		viper.ReadInConfig()
	}

	// Load config from env
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.Unmarshal(&config)
}
```

Compile toàn bộ rồi chạy thôi

```
$ go build -o 5kbps.io ./
$ ./5kbps.io help
A very fast blog for slow connection built with love by duythinht and friends.
Donate me for a food at https://5kbps.io

Usage:
  5kbps [flags]
  5kbps [command]

Available Commands:
  generate-config Generate default config
  help            Help about any command

Flags:
      --config string   Path to config file
  -h, --help            help for 5kbps

Use "5kbps [command] --help" for more information about a command.
$ ./5kbps.io
Config is {Database:{Hostname:localhost Username:root Password:khongbiet}}
$ ./5kbps.io generate-config > test-config.toml
$ vim test-config.toml # edit config file
$ cat test-config.toml
[database]
	hostname = "duythinht.io"
	username = "root"
	password = "khongbiet"
$ ./5kbps.io --config test-config.toml
Config is {Database:{Hostname:duythinht.io Username:root Password:khongbiet}}
```

#### Chốt

Bài viết này có giúp tôi hết thất nghiệp không? Thành thật mà nói thì không, trong những ngày tháng đói há mồm, bị hỏi nhiều quá nên tôi mới viết bài này, hi vọng có thanh niên nào thấy hay thì tuyển tôi về làm ku li cũng được.
