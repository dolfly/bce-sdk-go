/*
 * Copyright 2026 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing
 * permissions
 * and limitations under the License.
 */

package cloudmemoryexamples

import (
	"context"
	"fmt"
	"os"

	"github.com/baidubce/bce-sdk-go/services/cloudmemory/api"
)

func FilesRetain() {
	endpoint, apiKey := "Your endpoint", "Your apiKey"
	client := cloudmemory.New(endpoint, apiKey)
	f, err := os.Open("/path/to/your/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	requestJSON := `{"items":[{"content":"file upload"}]}`
	response, err := client.FilesRetain(context.Background(), "your-bank-id", []*os.File{f}, requestJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
