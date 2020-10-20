# banner

:smile: ultra lightweight ascii-art generator

![CI](https://github.com/moul/banner/workflows/CI/badge.svg)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/banner)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/banner/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/banner.svg)](https://github.com/moul/banner/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/banner)](https://goreportcard.com/report/moul.io/banner)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/banner/badge)](https://www.codefactor.io/repository/github/moul/banner)
[![codecov](https://codecov.io/gh/moul/banner/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/banner)
[![GolangCI](https://golangci.com/badges/github.com/moul/banner.svg)](https://golangci.com/r/github.com/moul/banner)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

Pure-go library to generate ASCII-art banners from text.

* doesn't need any external dependencies
* very lightweight footprint at the end in the generated binary
  * ~38kb with default build options
  * ~12kb with upx + binary strip
* only one font available ([small.flf](http://www.figlet.org/fontdb_example.cgi?font=small.flf))
  * incomplete alphabet (a-zA-Z.-_?)

## Usage

```go
import "moul.io/banner"

fmt.Println(banner.Inline("hey world."))
```

```
 _                                    _     _
| |_   ___  _  _   __ __ __ ___  _ _ | | __| |
| ' \ / -_)| || |  \ V  V // _ \| '_|| |/ _` | _
|_||_|\___| \_, |   \_/\_/ \___/|_|  |_|\__,_|(_)
            |__/
```

## Install

### Using go

```console
$ go get -u moul.io/banner
```

### Releases

See https://github.com/moul/banner/releases

## License

© 2020 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
