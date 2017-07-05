/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"Status":1,"Message":"","Data":[{"logonAddr":"p103.tdstoy.com","versionAndroid":3,"versionIOS":5,"logonPort":9185,"msgAddr":"http://www.tdzcc.com","messageAddr":"p103.tdstoy.com","msgPort":9183,"urlAddr":"http://www.tdzcc.com","downloadUrlAndroid":"https://app.td3g.com/GW.html","downloadUrlIOS":"https://itunes.apple.com/cn/app/tdqp2/id1222825360?l=zh&ls=1&mt=8"}]}`)
}
