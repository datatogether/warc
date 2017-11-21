package warc

// Named fields within a WARC record provide information about the current
// record, and allow additional per-record information. WARC both reuses
// appropriate headers from other standards and defines new headers, all
// beginning "WARC-", for WARC-specific purposes.
//
// WARC named fields of the same type shall not be repeated in the same
// WARC record (for example, a WARC record shall not have several WARC-Date
// or several WARC-Target-URI), except as noted (e.g., WARC-Concurrent-To).
const (
	// An identifier assigned to the current record that is globally unique for
	// its period of intended use. No identifier scheme is mandated by this
	// specification, but each record-id shall be a legal URI and clearly
	// indicate a documented and registered scheme to which it conforms (e.g.,
	// via a URI scheme prefix such as "http:" or "urn:"). Care should be taken
	// to ensure that this value is written with no internal whitespace.
	FieldNameWARCRecordID = "WARC-Record-ID"
	// The number of octets in the block, similar to [RFC2616]. If no block is
	// present, a value of '0' (zero) shall be used.
	FieldNameContentLength = "Content-Length"
	// 	A 14-digit UTC timestamp formatted according to YYYY-MM-DDThh:mm:ssZ,
	// described in the W3C profile of ISO8601 [W3CDTF]. The timestamp shall
	// represent the instant that data capture for record creation began.
	// Multiple records written as part of a single capture event (see section
	// 5.7) shall use the same WARC-Date, even though the times of their
	// writing will not be exactly synchronized.
	FieldNameWARCDate = "WARC-Date"
	// 	The type of WARC record: one of 'warcinfo', 'response', 'resource',
	// 'request', 'metadata', 'revisit', 'conversion', or 'continuation'. Other
	// types of WARC records may be defined in extensions of the core format.
	// Types are further described in WARC Record Types.
	// A WARC file needs not contain any particular record types, though
	// starting all WARC files with a "warcinfo" record is recommended.
	FieldNameWARCType = "WARC-Type"
	// The MIME type [RFC2045] of the information contained in the record's
	// block. For example, in HTTP request and response records, this would be
	// 'application/http' as per section 19.1 of [RFC2616] (or
	// 'application/http; msgtype=request' and 'application/http;
	// msgtype=response' respectively). In particular, the content-type is not
	// the value of the HTTP Content-Type header in an HTTP response but a MIME
	// type to describe the full archived HTTP message (hence
	// 'application/http' if the block contains request or response headers).
	FieldNameContentType = "Content-Type"
	// 	The WARC-Record-IDs of any records created as part of the same capture
	// event as the current record. A capture event comprises the information
	// automatically gathered by a retrieval against a single target-URI; for
	// example, it might be represented by a 'response' or 'revisit' record
	// plus its associated 'request' record.
	// This field may be used to associate records of types 'request',
	// 'response', 'resource', 'metadata', and 'revisit' with one another when
	// they arise from a single capture event (When so used, any
	// WARC-Concurrent-To association shall be considered bidirectional even if
	// the  header only appears on one record.) The WARC Concurrent-to field
	// shall not be used in 'warcinfo', 'conversion', and 'continuation'
	// records.
	FieldNameWARCConcurrentTo = "WARC-Concurrent-To"
	// An optional parameter indicating the algorithm name and calculated value
	// of a digest applied to the full block of the record.
	// An example is a SHA-1 labelled Base32 ([RFC3548]) value:
	// WARC-Block-Digest: sha1:AB2CD3EF4GH5IJ6KL7MN8OPQ
	FieldNameWARCBlockDigest = "WARC-Block-Digest"
	// An optional parameter indicating the algorithm name and calculated value
	// of a digest applied to the payload referred to or contained by the
	// record - which is not necessarily equivalent to the record block.
	// The payload of an application/http block is its 'entity-body' (per
	// [RFC2616]). In contrast to WARC-Block-Digest, the WARC-Payload-Digest
	// field may also be used for data not actually present in the current
	// record block, for example when a block is left off in accordance with a
	// 'revisit' profile (see 'revisit'), or when a record is segmented (the
	// WARC-Payload-Digest recorded in the first segment of a segmented record
	// shall be the digest of the payload of the logical record).
	FieldNameWARCPayloadDigest = "WARC-Payload-Digest"
	// The numeric Internet address contacted to retrieve any included content.
	// An IPv4 address shall be written as a "dotted quad"; an IPv6 address
	// shall be written as per [RFC1884]. For an HTTP retrieval, this will be
	// the IP address used at retrieval time corresponding to the hostname in
	// the record's target-URI.
	FieldNameWARCIPAddress = "WARC-IP-Address"
	// The WARC-Refers-To field may be used to associate a 'metadata' record to
	// another record it describes. The WARC-Refers-To field may also be used
	// to associate a record of type 'revisit' or 'conversion' with the
	// preceding record which helped determine the present record content. The
	// WARC-Refers-To field shall not be used in 'warcinfo', 'response',
	// ‘resource’, 'request', and 'continuation' records.
	FieldNameWARCRefersTo = "WARC-Refers-To"
	// The original URI whose capture gave rise to the information content in
	// this record. In the context of web harvesting, this is the URI that was
	// the target of a crawler's retrieval request. For a 'revisit' record, it
	// is the URI that was the target of a retrieval request.  Indirectly, such
	// as for a 'metadata', or 'conversion' record, it is a copy of the
	// WARC-Target-URI appearing in the original record to which the newer
	// record pertains. The URI in this value shall be properly escaped
	// according to [RFC3986] and written with no internal whitespace.
	FieldNameWARCTargetURI = "WARC-Target-URI"
	// For practical reasons, writers of the WARC format may place limits on
	// the time or storage allocated to archiving a single resource. As a
	// result, only a truncated portion of the original resource may be
	// available for saving into a WARC record.
	//
	// Any record may indicate that truncation of its content block has
	// occurred and give the reason with a 'WARC-Truncated' field.
	FieldNameWARCTruncated = "WARC-Truncated"
	// When present, indicates the WARC-Record-ID of the associated 'warcinfo'
	// record for this record. Typically, the Warcinfo-ID parameter is used
	// when the context of the applicable 'warcinfo' record is unavailable,
	// such as after distributing single records into separate WARC files. WARC
	// writing applications (such web crawlers) may choose to always record
	// this parameter.
	FieldNameWARCWarcinfoID = "WARC-Warcinfo-ID"
	// The WARC-Filename field may be used in 'warcinfo' type records and shall
	// not be used for other record types.
	FieldNameWARCFilename = "WARC-Filename"
	// A URI signifying the kind of analysis and handling applied in a
	// 'revisit' record. (Like an XML namespace, the URI may, but need not,
	// return human-readable or machine-readable documentation.) If reading
	// software does not recognize the given URI as a supported kind of
	// handling, it shall not attempt to interpret the associated record block.
	FieldNameWARCProfile = "WARC-Profile"
	// The content-type of the record's payload as determined by an independent
	// check. This string shall not be arrived at by blindly promoting an HTTP
	// Content-Type value up from a record block into the WARC header without
	// direct analysis of the payload, as such values may often be unreliable.
	FieldNameWARCIdentifiedPayloadType = "WARC-Identified-Payload-Type"
	// Reports the current record's relative ordering in a sequence of
	// segmented records.
	// In the first segment of any record that is completed in one or more
	// later 'continuation' WARC records, this parameter is mandatory. Its
	// value there is "1". In a 'continuation' record, this parameter is also
	// mandatory. Its value is the sequence number of the current segment in
	// the logical whole record, increasing by 1 in each next segment.
	FieldNameWARCSegmentNumber = "WARC-Segment-Number"
	// Identifies the starting record in a series of segmented records whose
	// content blocks are reassembled to obtain a logically complete content
	// block.
	// This field is mandatory on all 'continuation' records, and shall not be
	// used in other records. See the section below, Record segmentation, for
	// full details on the use of WARC record segmentation.
	FieldNameWARCSegmentOriginID = "WARC-Segment-Origin-ID"
	// In the final record of a segmented series, reports the total length of
	// all segment content blocks when concatenated together.
	// This field is mandatory on the last 'continuation' record of a series,
	// and shall not be used elsewhere.
	FieldNameWARCSegmentTotalLength = "WARC-Segment-Total-Length"
)
