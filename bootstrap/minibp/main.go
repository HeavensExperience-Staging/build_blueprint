// Copyright 2014 Google Inc. All rights reserved.
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

package main

import (
	"flag"
	"path/filepath"

	"github.com/google/blueprint"
	"github.com/google/blueprint/bootstrap"
)

var runAsPrimaryBuilder bool

func init() {
	flag.BoolVar(&runAsPrimaryBuilder, "p", false, "run as a primary builder")
}

type Config struct {
}

func (c Config) SrcDir() string {
	return bootstrap.CmdlineArgs.BuildDir
}

func (c Config) RemoveAbandonedFilesUnder(buildDir string) (under, exempt []string) {
	if !runAsPrimaryBuilder {
		under = []string{filepath.Join(buildDir, ".bootstrap")}
		exempt = []string{filepath.Join(buildDir, ".bootstrap", "build.ninja")}
	}
	return
}

func main() {
	flag.Parse()

	ctx := blueprint.NewContext()
	if !runAsPrimaryBuilder {
		ctx.SetIgnoreUnknownModuleTypes(true)
	}

	config := Config{}
	bootstrap.Main(ctx, config, !runAsPrimaryBuilder)
}
