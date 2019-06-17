package main

import (
	"fmt"
	"strings"

	"gopkg.in/couchbase/gocb.v1"
)

// LsDefaultSettings of cname
type LsDefaultSettings struct {
	Type               int32   `json:"type"`
	EnabledModuleTypes []int32 `json:"enabledModuleTypes"`
}

// CompanySettings from couchbase
type CompanySettings struct {
	DisplayName   string `json:"displayName"`
	ParentCompany string `json:"parent_company"`
}

// IDResponse result from N1ql
type IDResponse struct {
	ID string `json:id`
}

func contains(arr []int32, moduleType int32) bool {
	for _, a := range arr {
		if a == moduleType {
			return true
		}
	}
	return false
}

func main() {
	cluster, err1 := gocb.Connect("couchbase://localhost")

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "",
		Password: "",
	})
	if err1 == nil {
		fmt.Println("Connected to COUCHBASE SERVER")
	} else {
		fmt.Println("ERROR CONNECT THE COUCHBASE SERVER:", err1)
	}
	bucket, _ := cluster.OpenBucket("ce", "")
	// Use query
	query := gocb.NewN1qlQuery("select meta(ce).id from ce where type=5")

	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"African Swallows"})
	var row IDResponse
	for rows.Next(&row) {
		cname := row.ID
		fmt.Println("Checking for: " + cname)
		var lsDefaultSettings LsDefaultSettings
		bucket.Get(cname, &lsDefaultSettings)
		//fmt.Printf("Default Settings: %v\n", lsDefaultSettings)

		if contains(lsDefaultSettings.EnabledModuleTypes, 9) || contains(lsDefaultSettings.EnabledModuleTypes, 10) || contains(lsDefaultSettings.EnabledModuleTypes, 11) || contains(lsDefaultSettings.EnabledModuleTypes, 12) {
			stringSlice := strings.Split(cname, ".")
			fmt.Println(stringSlice[0])
			var companySettings CompanySettings
			bucket.Get(stringSlice[0]+".settings", &companySettings)
			fmt.Println("Mission Enabled: " + stringSlice[0] + " Name: " + companySettings.ParentCompany)
		}
	}
}
