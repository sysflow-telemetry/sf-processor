//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package exporter

// Event defines an interface for exported event objects.
type Event interface {
	ToJSON() []byte
	ToJSONStr() string
}
