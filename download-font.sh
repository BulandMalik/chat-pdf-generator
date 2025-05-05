#!/bin/bash

# Create fonts directory
mkdir -p fonts

# Download DejaVu font
curl -L "https://github.com/dejavu-fonts/dejavu-fonts/raw/master/ttf/DejaVuSansCondensed.ttf" -o "fonts/DejaVuSansCondensed.ttf" 