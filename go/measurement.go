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

type Measurement struct {

	// Name of the variable for which we are creating the measurement records
	VariableName string `json:"variableName,omitempty"`

	// Application or device used to record the measurement values
	SourceName string `json:"sourceName,omitempty"`

	// Start Time for the measurement event in UTC ISO 8601 \"YYYY-MM-DDThh:mm:ss\"
	StartTimeString string `json:"startTimeString,omitempty"`

	// Seconds between the start of the event measured and 1970 (Unix timestamp)
	StartTimeEpoch int32 `json:"startTimeEpoch,omitempty"`

	HumanTime HumanTime `json:"humanTime,omitempty"`

	// Converted measurement value in requested unit
	Value float64 `json:"value,omitempty"`

	// Original value as originally submitted
	OriginalValue int32 `json:"originalValue,omitempty"`

	// Original Unit of measurement as originally submitted
	OriginalunitAbbreviatedName string `json:"originalunitAbbreviatedName,omitempty"`

	// Abbreviated name for the unit of measurement
	UnitAbbreviatedName string `json:"unitAbbreviatedName,omitempty"`

	// Note of measurement
	Note string `json:"note,omitempty"`
}
