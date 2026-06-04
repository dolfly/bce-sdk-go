// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License
// See https://github.com/vectorize-io/hindsight
//
// The original MIT LICENSE text from the Hindsight project is preserved
// in the LICENSE file in this directory.

// Package cloudmemory wraps the upstream Hindsight Go SDK
// (github.com/vectorize-io/hindsight/hindsight-clients/go) and exposes a
// simplified, idiomatic facade aligned with the Baidu Cloud Memory HTTP
// reference table at https://cloud.baidu.com/doc/VDB/s/cmpl4uayz.
//
// Endpoints are 1:1 with Hindsight; each method delegates to the underlying
// hindsight.APIClient. Callers may also access the raw client via Underlying().
package cloudmemory
