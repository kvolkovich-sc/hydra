// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jwk

import (
	"crypto/rsa"

	"github.com/pkg/errors"
	"github.com/square/go-jose"
)

func MustRSAPublic(key *jose.JSONWebKey) *rsa.PublicKey {
	res, err := ToRSAPublic(key)
	if err != nil {
		panic(err.Error())
	}
	return res

}

func ToRSAPublic(key *jose.JSONWebKey) (*rsa.PublicKey, error) {
	res, ok := key.Key.(*rsa.PublicKey)
	if !ok {
		return res, errors.New("Could not convert key to RSA Private Key.")
	}
	return res, nil
}

func MustRSAPrivate(key *jose.JSONWebKey) *rsa.PrivateKey {
	res, err := ToRSAPrivate(key)
	if err != nil {
		panic(err.Error())
	}
	return res
}

func ToRSAPrivate(key *jose.JSONWebKey) (*rsa.PrivateKey, error) {
	res, ok := key.Key.(*rsa.PrivateKey)
	if !ok {
		return res, errors.New("Could not convert key to RSA Private Key.")
	}
	return res, nil
}
