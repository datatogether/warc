// Package warc is an implementation of ISO28500 1.0, the WebARCive specfication.
// it provides readers, writers, and structs for working with warc records.
// from the spec:
// The WARC (Web ARChive) file format offers a convention for concatenating
// multiple resource records (data objects), each consisting of a set of
// simple text headers and an arbitrary data block into one long file. The
// WARC format is an extension of the ARC File Format [ARC] that has
// traditionally been used to store "web crawls" as sequences of content
// blocks harvested from the World Wide Web. Each capture in an ARC file is
// preceded by a one-line header that very briefly describes the harvested
// content and its length. This is directly followed by the retrieval
// protocol response messages and content. The original ARC format file is
// used by the Internet Archive (IA) since 1996 for managing billions of
// objects, and by several national libraries.
package warc
