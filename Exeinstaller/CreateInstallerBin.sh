#!/bin/bash

if [ "$#" -lt 2 ]; then
  echo "Usage: $0 <payload.tar.xz> <output-installer.bin>"
  exit 2
fi

PAYLOAD="$1"
OUT="$2"

if [ ! -f "$PAYLOAD" ]; then
  echo "Payload file not found"
  exit 3
fi

cat MyInstallScript.sh > "$OUT"
echo "" >> "$OUT"

printf "%s\n" "::hostzip::" >> "$OUT"
base64 -w 0 "$PAYLOAD" >> "$OUT" 
echo "" >> "$OUT"
printf "%s\n" "::hostzip_end::" >> "$OUT"

chmod +x "$OUT"
echo "Created Installer: $OUT"