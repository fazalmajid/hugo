// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"bytes"

	"github.com/dchest/cssmin"
	"github.com/spf13/hugo/source"
	"github.com/spf13/hugo/tpl"
	"log"
)

func init() {
	RegisterHandler(new(cssHandler))
	RegisterAdvancedJPEGHandler()
	RegisterHandler(new(defaultHandler))
}

type basicFileHandler Handle

func (h basicFileHandler) Read(f *source.File, s *Site) HandledResult {
	return HandledResult{file: f}
}

func (h basicFileHandler) PageConvert(*Page, tpl.Template) HandledResult {
	return HandledResult{}
}

type defaultHandler struct{ basicFileHandler }

func (h defaultHandler) Extensions() []string { return []string{"*"} }
func (h defaultHandler) FileConvert(f *source.File, s *Site) HandledResult {
	s.writeDestFile(f.Path(), f.Contents)
	return HandledResult{file: f}
}

type cssHandler struct{ basicFileHandler }

func (h cssHandler) Extensions() []string { return []string{"css"} }
func (h cssHandler) FileConvert(f *source.File, s *Site) HandledResult {
	x := cssmin.Minify(f.Bytes())
	s.writeDestFile(f.Path(), bytes.NewReader(x))
	return HandledResult{file: f}
}

// special handler for JPEGs, not installed by default
type advancedJpegHandler struct{ basicFileHandler }

func (h advancedJpegHandler) Extensions() []string { return []string{"jpg", "jpeg"} }
func (h advancedJpegHandler) FileConvert(f *source.File, s *Site) HandledResult {
	s.writeDestFile(f.Path(), f.Contents)
	log.Println("JPEG doing", f.LogicalName())
	return HandledResult{file: f}
}

func RegisterAdvancedJPEGHandler() {
	RegisterHandler(new(advancedJpegHandler))
	log.Println("registered advanced JPEG handler")
}
