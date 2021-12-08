# recter

![current version](https://img.shields.io/badge/dynamic/json?color=green&label=Version&query=%24.latest_version&url=https%3A%2F%2Fmpldr.codes%2Frecter%2Fapi%2Fversions%2Flatest&style=flat-square&logo=git&color=F05032)
[![docker-pulls](https://img.shields.io/docker/pulls/mpldr/recter?logo=docker&logoColor=white&style=flat-square)](https://hub.docker.com/r/mpldr/recter)
[![gimme-money](https://img.shields.io/badge/dynamic/json?color=yellow&label=donations&query=%24.total&suffix=%E2%82%AC&url=https%3A%2F%2Fmoritz.sh%2Fdonate%2Fstats.json&style=flat-square)](https://moritz.sh/donate/)
![License](https://img.shields.io/static/v1?label=license&message=GPL-3&color=blue&style=flat-square)

[![Demo](https://img.shields.io/website?down_color=red&down_message=offline%20%3A%28&label=demo&style=for-the-badge&up_color=green&up_message=click%20me&url=https%3A%2F%2Fmpldr.codes)](https://mpldr.codes)

Moving go modules between hosting providers is a hazzle, you have to change the
import path, notify your users and do all those tedious things. recter takes
this from you.

## So what does it do?

It redirects go to wherever your code actually lives and is able to be easily
changed to direct somewhere else if the need arises.

`github.com/your-username/some-obscure-wordplay-with-go` becomes
`your.domain.tld/some-project`

For an example look no further than to this very project, which is available at
[mpldr.codes/recter](https://mpldr.codes/recter)

But wait, there's more! recter also features the following:

- a simple JSON API to query basic repo information
- the possibility for a nice portal
- a sensible docker setup
- batteries included

# License

recter is [Libre Software](https://moritz.sh/blog/libre-not-free/) as defined
by the FSF and licensed under the GPL
