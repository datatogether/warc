package warc

// Records is a slice of records
// A WARC format file is the simple concatenation of one or more WARC
// records. The first record usually describes the records to follow. In
// general, record content is either the direct result of a retrieval
// attempt — web pages, inline images, URL redirection information, DNS
// hostname lookup results, standalone files, etc. — or is synthesized
// material (e.g., metadata, transformed content) that provides additional
// information about archived content.
type Records []*Record

// FilterTypes return all record types that match a provide
// list of RecordTypes
func (rs Records) FilterTypes(types ...RecordType) Records {
	res := Records{}
	for _, rec := range rs {
		for _, t := range types {
			if rec.Type == t {
				res = append(res, rec)
			}
		}
	}
	return res
}

// TargetURIRecord returns a record matching uri optionally filtered by
// a list of record types. There are a number of "gotchas" if multiple
// record types of the same url are in the list.
// TODO - eliminate "gotchas"
func (rs Records) TargetURIRecord(uri string, types ...RecordType) *Record {
	for _, rec := range rs {
		if rec.TargetURI() == uri {
			if len(types) == 0 {
				return rec
			}
			for _, t := range types {
				if rec.Type == t {
					return rec
				}
			}
		}
	}
	return nil
}

// RemoveTargetURIRecords returns a Records slice with all records
// that refer to uri removed
func (rs Records) RemoveTargetURIRecords(uri string) (recs Records) {
	recs = rs
	for i, rec := range rs {
		if rec.TargetURI() == uri {
			recs = append(recs[:i], recs[i+1:]...)
		}
	}
	return
}
