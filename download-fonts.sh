#!/bin/bash

# Create fonts directory
mkdir -p fonts

# Download DejaVu font
curl -L "https://github.com/prawnpdf/prawn/blob/master/data/fonts/DejaVuSans.ttf" -o "fonts/DejaVuSans.ttf"

# Download Noto Color Emoji font
curl -L "https://github.com/googlefonts/noto-emoji/blob/main/fonts/NotoColorEmoji.ttf" -o "fonts/NotoColorEmoji-Regular.ttf"

echo "Fonts downloaded successfully!" 