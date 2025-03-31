#!/bin/bash

# Define an array of build targets: (GOOS GOARCH OUTPUT)
targets=(
  "windows amd64 status_win64.exe"
  "windows 386 status_win32.exe"
  "darwin amd64 status_intel"
  "darwin arm64 status_arm64"
)

SOURCE_FILE="./cmd/status.go"
DIST_DIR="./dist"
CHECKSUMS_FILE="$DIST_DIR/checksums.txt"

rm -rf $DIST_DIR
mkdir $DIST_DIR

touch $CHECKSUMS_FILE

# Iterate over each target and build
for target in "${targets[@]}"; do
  set -- $target  # Split the string into separate variables
  GOOS=$1
  GOARCH=$2
  OUTPUT=$DIST_DIR/$3

  echo "Building for $GOOS/$GOARCH..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT $SOURCE_FILE
  shasum -a 256 $OUTPUT >> $CHECKSUMS_FILE

  if [ $? -eq 0 ]; then
    echo "✅ Successfully built: $OUTPUT"
  else
    echo "❌ Failed to build: $OUTPUT"
    exit 1
  fi
done

echo "🎉 All builds completed!"
