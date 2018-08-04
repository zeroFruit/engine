/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

type MockProcess interface {
	Init(id string, ipAddress string)

	GetId() string

	Publish(exchange string, topic string, data interface{}) error

	HandleEvent()

	HandleCommand()

	TriggerCommand(data interface{}, process MockProcess)

	TriggerEvent(data interface{}, process MockProcess)
}