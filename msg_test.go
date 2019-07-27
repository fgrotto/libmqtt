/*
 * Copyright Go-IIoT (https://github.com/goiiot)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package libmqtt

import (
	"fmt"
	"testing"
)

func TestNotifyMsg(t *testing.T) {
	msgCh := make(chan *message)
	testErr := fmt.Errorf("test error")
	go func() {
		notifyNetMsg(msgCh, "test srv", testErr)
		notifyPersistMsg(msgCh, nil, testErr)
		notifyPubMsg(msgCh, "test topic", testErr)
		notifySubMsg(msgCh, []*Topic{}, testErr)
		notifyUnsubMsg(msgCh, []string{}, testErr)

		close(msgCh)
	}()

	for msg := range msgCh {
		if msg.err == nil {
			t.Error("message error nil")
		}
	}
}
