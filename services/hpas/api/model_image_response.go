/*
 * Copyright 2025 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

type ImageResponse struct {
	ImageId          string   `json:"imageId"`
	Name             string   `json:"name"`
	ImageType        string   `json:"imageType"`
	ImageStatus      string   `json:"imageStatus"`
	CreateTime       string   `json:"createTime"`
	SupportedAppType []string `json:"supportedAppType"`
	ImageSizeInGB    int 	  `json:"imageSizeInGB"`
	MinDiskInGB  	 int 	  `json:"minDiskInGB"`
}
