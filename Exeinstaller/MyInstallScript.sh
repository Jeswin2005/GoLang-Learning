#!/bin/bash

SELF="$0"
INSTALL_DIR="/home/jeswin-pt8024"
PRODUCT_DIR_NAME="myagent"
TARGET_DIR="$INSTALL_DIR/$PRODUCT_DIR_NAME"
HOSTZIP_SUBDIR="hostzip"
LOG_FILE="$INSTALL_DIR/logs/commonlogs.log"

echo "Installer started..."

mkdir -p "$TARGET_DIR"

if grep --text -q "^::hostzip::" "$SELF" && grep --text -q "^::hostzip_end::" "$SELF"; then
    echo "Payload markers found — extracting payload directly to $TARGET_DIR ..."
    if ! awk '/^::hostzip::/{flag=1;next} /^::hostzip_end::/{flag=0} flag{print}' "$SELF" \
           | base64 -d \
           | tar -xJf - -C "$TARGET_DIR"; then
        echo "Extraction failed — cleaning up partial extraction..."
        if [ -d "$TARGET_DIR" ]; then
            rm -rf -- "$TARGET_DIR"
        fi
        echo "Exiting due to extraction failure."
        exit 1
    fi
else
    echo "No Payload found. exiting..."
    exit 1
fi

if [ ! -d "$TARGET_DIR/$HOSTZIP_SUBDIR" ]; then
    echo "Payload structure invalid — missing $HOSTZIP_SUBDIR in $TARGET_DIR"
    rm -rf -- "$TARGET_DIR"
    exit 1
fi

echo "Payload extracted to: $TARGET_DIR/$HOSTZIP_SUBDIR"

mkdir -p "$(dirname "$LOG_FILE")"
if [ ! -f "$LOG_FILE" ]; then
    touch "$LOG_FILE"
    chmod 0644 "$LOG_FILE"
fi

BIN_SRC="$TARGET_DIR/$HOSTZIP_SUBDIR/bin"
if [ -d "$BIN_SRC" ]; then
    for f in "$BIN_SRC"/*; do
        [ -f "$f" ] || continue
        fname=$(basename "$f")
        install -m 0755 "$f" "$BIN_SRC/$fname.new" || { echo "install failed for $fname"; exit 1; }
        mv -f "$BIN_SRC/$fname.new" "$BIN_SRC/$fname"
    done
else
    echo "No bins found at $BIN_SRC"
fi

SYSTEMD_DIR="/etc/systemd/system"
mkdir -p "$SYSTEMD_DIR"
SYSTEMD_SRC_DIR="$TARGET_DIR/$HOSTZIP_SUBDIR/systemd"
if [ -d "$SYSTEMD_SRC_DIR" ]; then
    for unit in "$SYSTEMD_SRC_DIR"/*; do
        [ -e "$unit" ] || continue
        unitname=$(basename "$unit")
        sudo mv -f "$unit" "$SYSTEMD_DIR/$unitname" || { echo "Failed to move $unitname"; exit 1; }
        sudo chmod 0644 "$SYSTEMD_DIR/$unitname"
    done
else
    echo "No systemd units found at $SYSTEMD_SRC_DIR"
fi

sudo systemctl daemon-reload
sudo systemctl enable --now logworker.timer || echo "Failed to enable/start logworker.timer"
sudo systemctl enable --now updater.timer || echo "Failed to enable/start updater.timer"

echo "Installation complete"
exit 0
