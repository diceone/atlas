/*
   Copyright 2019 Stellar Project

   Permission is hereby granted, free of charge, to any person obtaining a copy of
   this software and associated documentation files (the "Software"), to deal in the
   Software without restriction, including without limitation the rights to use, copy,
   modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
   and to permit persons to whom the Software is furnished to do so, subject to the
   following conditions:

   The above copyright notice and this permission notice shall be included in all copies
   or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
   INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
   PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
   FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
   TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
   USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package client

import (
	"context"
	"fmt"
	"strings"

	api "github.com/stellarproject/atlas/api/v1"
)

// Create is used to create new records
func (c *Client) Create(name string, records []*api.Record) error {
	ctx := context.Background()
	if _, err := c.nameserverService.Create(ctx, &api.CreateRequest{
		Name:    name,
		Records: records,
	}); err != nil {
		return err
	}
	return nil
}

// List returns all records in the datastore
func (c *Client) List() ([]*api.Record, error) {
	ctx := context.Background()
	resp, err := c.nameserverService.List(ctx, &api.ListRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Records, nil
}

// Delete removes records from the datastore
func (c *Client) Delete(name string) error {
	ctx := context.Background()
	if _, err := c.nameserverService.Delete(ctx, &api.DeleteRequest{
		Name: name,
	}); err != nil {
		return err
	}
	return nil
}

// RecordType is a helper function to resolve the record type from a human friendly string
func (c *Client) RecordType(rtype string) (api.RecordType, error) {
	switch strings.ToUpper(rtype) {
	case "A":
		return api.RecordType_A, nil
	case "CNAME":
		return api.RecordType_CNAME, nil
	case "SRV":
		return api.RecordType_SRV, nil
	case "TXT":
		return api.RecordType_TXT, nil
	case "MX":
		return api.RecordType_MX, nil
	default:
		return api.RecordType_UNKNOWN, fmt.Errorf("unsupported record type %q", rtype)
	}
}
