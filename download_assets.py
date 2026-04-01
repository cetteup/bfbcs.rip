#!/usr/bin/env python3
"""
Download all assets referenced in bfbcs.css from the Web Archive.
"""

import os
import re
import sys
import time
from urllib.error import URLError
from urllib.parse import unquote
from urllib.request import urlopen, Request

# Configuration
CSS_FILE = "public/static/css/bfbcs.css"
HTML_FILE = "public/views/stats.html"
OUTPUT_DIR = "public/static"
WEB_ARCHIVE_BASE = "https://web.archive.org/"
DOWNLOAD_DELAY = 30  # seconds between downloads to be respectful to Web Archive

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
        url = unquote(url)

        # Only process web archive URLs
        if url.startswith('/web/'):
            urls.add(url)

    return sorted(urls)

def extract_urls_from_html(html_file):
    """Extract all bfbcs.com URLs from HTML file."""
    urls = set()

    with open(html_file, 'r', encoding='utf-8') as f:
        content = f.read()

    # Find all http://files.bfbcs.com/ or files2.bfbcs.com/ URLs
    pattern = r'https?://(?:files|files2)\.bfbcs\.com/[^\'"\s]+'
    matches = re.findall(pattern, content)

    for match in matches:
        match = unquote(match)
        # Convert to archive URL
        archive_url = f"/web/20140209014926im_/{match}"
        urls.add(archive_url)

    return urls

def convert_to_full_url(archive_url):
    """Convert relative Web Archive URL to full URL."""
    return WEB_ARCHIVE_BASE + archive_url.lstrip('/')

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

def download_file(url, local_path):
    """Download a file from URL to local path."""
    try:
        # Create parent directories
        os.makedirs(os.path.dirname(local_path), exist_ok=True)

        # Download file with User-Agent to avoid blocking
        headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36'
        }
        req = Request(url, headers=headers)

        print(f"Downloading: {url}")
        with urlopen(req, timeout=30) as response:
            with open(local_path, 'wb') as out_file:
                out_file.write(response.read())

        print(f"  Saved to: {local_path}")
        return True
    except URLError as e:
        print(f"  ERROR: Failed to download - {e}", file=sys.stderr)
        return False
    except Exception as e:
        print(f"  ERROR: {e}", file=sys.stderr)
        return False

def main():
    """Main function."""
    print("Extracting URLs from CSS file...")
    urls = extract_urls_from_css(CSS_FILE)

    print("Extracting URLs from HTML file...")
    html_urls = extract_urls_from_html(HTML_FILE)
    urls.extend(html_urls)

    # Add rank icons
    print("Adding rank icon URLs...")
    for i in range(1, 51):
        rank = f"r{i:03d}"
        urls.append(f"/web/20140209014926im_/http://files2.bfbcs.com/img/bfbcs/ranks/{rank}.png")
        urls.append(f"/web/20140209014926im_/http://files2.bfbcs.com/img/bfbcs/ranks_big/{rank}.png")

    print(f"Found {len(urls)} unique URLs\n")

    successful = 0
    failed = 0

    for i, archive_url in enumerate(urls):
        full_url = convert_to_full_url(archive_url)
        local_path = get_local_path(archive_url)

        if local_path is None:
            print(f"SKIP: Could not parse URL: {archive_url}", file=sys.stderr)
            failed += 1
            continue

        # Skip if file already exists
        if os.path.exists(local_path):
            print(f"EXISTS: {local_path}")
            successful += 1
            continue

        if download_file(full_url, local_path):
            successful += 1
        else:
            failed += 1

        # Add delay between downloads to be respectful to Web Archive
        if i < len(urls) - 1:
            print(f"Waiting {DOWNLOAD_DELAY}s before next download...")
            time.sleep(DOWNLOAD_DELAY)

    print(f"\n{'='*60}")
    print(f"Download complete!")
    print(f"Successful: {successful}")
    print(f"Failed: {failed}")
    print(f"Total: {successful + failed}")

    return 0 if failed == 0 else 1

if __name__ == "__main__":
    sys.exit(main())

