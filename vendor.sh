#!/bin/sh

curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_macos -o ./vendor/yt-dlp
xattr -c ./vendor/yt-dlp
chmod u+rx ./vendor/yt-dlp

