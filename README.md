# warc
[![GitHub](https://img.shields.io/badge/project-Data_Together-487b57.svg?style=flat-square)](http://github.com/datatogether)
[![Slack](https://img.shields.io/badge/slack-Archivers-b44e88.svg?style=flat-square)](https://archivers-slack.herokuapp.com/)
[![GoDoc](https://godoc.org/github.com/datatogether/warc?status.svg)](http://godoc.org/github.com/datatogether/warc)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](./LICENSE) 

warc is an implementation of ISO28500 1.0, the WebARCive specfication.
it provides readers, writers, and structs for working with warc records.

from the spec:
> The WARC (Web ARChive) file format offers a convention for concatenating
multiple resource records (data objects), each consisting of a set of
simple text headers and an arbitrary data block into one long file. The
WARC format is an extension of the ARC File Format [ARC] that has
traditionally been used to store "web crawls" as sequences of content
blocks harvested from the World Wide Web. Each capture in an ARC file is
preceded by a one-line header that very briefly describes the harvested
content and its length. This is directly followed by the retrieval
protocol response messages and content. The original ARC format file is
used by the Internet Archive (IA) since 1996 for managing billions of
objects, and by several national libraries.
package warc

## License & Copyright

[Affero General Public License v3](http://www.gnu.org/licenses/agpl.html) ]

## Getting Involved

We would love involvement from more people! If you notice any errors or would like to submit changes, please see our [Contributing Guidelines](./.github/CONTRIBUTING.md). 

We use GitHub issues for [tracking bugs and feature requests](https://github.com/datatogether/REPONAME/issues) and Pull Requests (PRs) for [submitting changes](https://github.com/datatogether/REPONAME/pulls)

## Usage
`import "gitnub.com/datatogether/warc"`
