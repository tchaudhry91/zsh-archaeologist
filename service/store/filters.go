package store

import "go.mongodb.org/mongo-driver/bson"

// SelectAllFilter returns all documents
func SelectAllFilter() bson.D {
	return bson.D{}
}

// SelectForUserFilter returns all documents for a single user
func SelectForUserFilter(user string) bson.D {
	return bson.D{
		{Key: "user", Value: user},
	}
}

// SelectSinceTimestampFilter returns all documents since a particular timestamp
func SelectSinceTimestampFilter(ts uint64) bson.D {
	return bson.D{
		{
			Key: "entry.timestamp",
			Value: map[string]uint64{
				"$gt": ts,
			},
		},
	}
}

// SelectTillTimestampFilter returns all documents till a particular timestamp
func SelectTillTimestampFilter(ts uint64) bson.D {
	return bson.D{
		{
			Key: "entry.timestamp",
			Value: map[string]uint64{
				"$lt": ts,
			},
		},
	}
}

// SelectTimerangeFilter returns all documents within two timestamps
func SelectTimerangeFilter(start, end uint64) bson.D {
	return AndMergeFilters(SelectSinceTimestampFilter(start), SelectTillTimestampFilter(end))
}

// AndMergeFilters combines multiple filters with an and operation
func AndMergeFilters(filters ...bson.D) bson.D {
	return bson.D{
		{
			Key:   "$and",
			Value: filters,
		},
	}
}
