/* 
 * QuantiModo
 *
 * QuantiModo makes it easy to retrieve normalized user data from a wide array of devices and applications. [Learn about QuantiModo](https://quantimo.do), check out our [docs](https://github.com/QuantiModo/docs) or contact us at [help.quantimo.do](https://help.quantimo.do). 
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"time"
)

type Update struct {

	// id
	Id int32 `json:"id,omitempty"`

	// userId
	UserId int32 `json:"userId,omitempty"`

	// connectorId
	ConnectorId int32 `json:"connectorId,omitempty"`

	// numberOfMeasurements
	NumberOfMeasurements int32 `json:"numberOfMeasurements,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// When the record was first created. Use UTC ISO 8601 \"YYYY-MM-DDThh:mm:ss\"  datetime format
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// When the record in the database was last updated. Use UTC ISO 8601 \"YYYY-MM-DDThh:mm:ss\"  datetime format
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
