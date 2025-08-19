/*
Copyright 2021 Loggie Authors

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

package file

import (
	"crypto/md5"
	"fmt"
	"os"
	"testing"
)

func TestGenerateIdentifier(t *testing.T) {
	readSize := 150
	fileName := "/tmp/loggie/test.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file fail: %s\n", err)
		return
	}
	readBuffer := make([]byte, readSize)
	l, err := file.Read(readBuffer)
	if err != nil {
		fmt.Printf("read file fail: %s\n", err)
		return
	}
	if l < readSize {
		fmt.Printf("read bytes(%d) too short for readSize(%d)", l, readSize)
		return
	}
	identifier := fmt.Sprintf("%x", md5.Sum(readBuffer))
	fmt.Printf("identifier: %s\n", identifier)
}
