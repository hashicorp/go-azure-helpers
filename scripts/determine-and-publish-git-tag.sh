#!/bin/bash
# Copyright IBM Corp. 2018, 2025
# SPDX-License-Identifier: MPL-2.0


set -e

function determineGitTag {
  # Whilst we could do something fancy here to determine the version, tbh
  # it doesn't matter providing it's:
  # a) 0-based (so we don't bump the Go module version/trigger import path changes)
  # b) ordered
  # c) clear when it was released
  #
  # whilst an odd choice for an SDK, makes sense insofar as the API Versions
  # themselves are versioned too - so we can use a date/time stamp for now:
  # e.g.v0.YYYYMMDD.HHMMSS (v0.20220504.115712)
  #
  # turns out that Go doesn't like the minor/patch to start with a 0, else
  # the Go Module Proxy returns:
  # > bad request: version "v0.20220622.050833" is not canonical (wanted "")
  # as such we'll prefix the Hours-Minutes-Seconds segment with a `1` for now
  date '+v0.%Y%m%d.1%H%M%S'
}

function publish {
  local version=$1

  echo "Tagging as '$version'.."
  git tag "$version"

  echo "Pushing Tags.."
  git push --tags
}

function main {
  local gitTag
  gitTag=$(determineGitTag)
  publish "$gitTag"
}

main
