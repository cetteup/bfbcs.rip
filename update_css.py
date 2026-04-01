#!/usr/bin/env python3
"""
Update bfbcs.css to use local asset paths for downloaded files.
Replaces Web Archive URLs with local paths for assets that exist in public/static/.
"""

import re
import os
import sys

# Configuration
CSS_FILE = "public/static/bfbcs.css"
OUTPUT_DIR = "public/static"

# URL pattern to extract from CSS
URL_PATTERN = r'url\(([^)]+)\)'

def extract_urls_from_css(css_file):
    """Extract all URLs from a CSS file."""
    urls = set()

    with open(css_file, 'r', encoding='utf-8') as f:
        content = f.read()

    # Find all url() patterns
    matches = re.findall(URL_PATTERN, content)

    for match in matches:
        # Remove quotes if present
        url = match.strip('\'"')

        # Only process web archive URLs
        if url.startswith('/web/'):
            urls.add(url)

    return sorted(urls)

def get_local_path(archive_url):
    """Convert archive URL to local file path."""
    # Extract the path after the timestamp
    # Format: /web/20140208054952im_/http://files.bfbcs.com/img/bfbcs/logo.png
    match = re.search(r'/web/\d+[a-z]*_/(.*)', archive_url)
    if match:
        original_url = match.group(1)
        # Extract path after domain
        # http://files.bfbcs.com/img/bfbcs/logo.png -> img/bfbcs/logo.png
        path_match = re.search(r'https?://[^/]+/(.*)', original_url)
        if path_match:
            relative_path = path_match.group(1)
            return os.path.join(OUTPUT_DIR, relative_path)

    return None

def update_css_file(css_file):
    """Update CSS file to use local paths for downloaded assets."""
    print(f"Reading CSS file: {css_file}")

    with open(css_file, 'r', encoding='utf-8') as f:
        content = f.read()

    urls = extract_urls_from_css(css_file)
    print(f"Found {len(urls)} unique URLs")

    replacements = 0

    for archive_url in urls:
        local_path = get_local_path(archive_url)

        if local_path and os.path.exists(local_path):
            # Create relative path from CSS file to asset
            # CSS file is at public/static/bfbcs.css
            # Assets are at public/static/img/... etc.
            # So relative path is just img/... etc.
            relative_path = os.path.relpath(local_path, OUTPUT_DIR)

            # Replace the archive URL with local path
            old_url = f'url({archive_url})'
            new_url = f'url(/static/{relative_path})'

            if old_url in content:
                content = content.replace(old_url, new_url)
                replacements += 1
                print(f"  ✓ Updated: {archive_url} -> /static/{relative_path}")
            else:
                print(f"  ⚠ Could not find exact match for: {archive_url}")
        else:
            print(f"  ✗ Asset not found locally: {archive_url}")

    # Write back the updated content
    with open(css_file, 'w', encoding='utf-8') as f:
        f.write(content)

    print(f"\n{'='*60}")
    print(f"CSS update complete!")
    print(f"Replacements made: {replacements}")
    print(f"Total URLs processed: {len(urls)}")

    return replacements

def main():
    """Main function."""
    if not os.path.exists(CSS_FILE):
        print(f"ERROR: CSS file not found: {CSS_FILE}", file=sys.stderr)
        return 1

    replacements = update_css_file(CSS_FILE)

    if replacements == 0:
        print("No assets were found locally to update.")
        return 1

    return 0

if __name__ == "__main__":
    sys.exit(main())

