#!/usr/bin/env python3
"""
Script to validate HTML content matches between two URLs.
Ignores indentation differences and prints lines with differences.

Usage: python3 compare_html.py <url1> <url2>
"""

import sys
from difflib import unified_diff
import urllib.error
import urllib.parse
import urllib.request


def fetch_html(url):
    """Fetch HTML content from a URL."""
    try:
        # Create a request with a User-Agent header to avoid 403 errors from Cloudflare
        headers = {
            'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36'
        }
        request = urllib.request.Request(url, headers=headers)
        with urllib.request.urlopen(request, timeout=10) as response:
            charset = response.headers.get_content_charset() or "utf-8"
            return response.read().decode(charset, errors="replace")
    except (urllib.error.URLError, ValueError) as e:
        print(f"Error fetching {url}: {e}", file=sys.stderr)
        sys.exit(1)


def format_html(html_content):
    """
    Format HTML content using BeautifulSoup's prettify method.
    Returns the formatted HTML with proper indentation.
    """
    try:
        from bs4 import BeautifulSoup
    except ImportError:
        print("Error: beautifulsoup4 is required for --format flag", file=sys.stderr)
        print("Install it with: pip install beautifulsoup4", file=sys.stderr)
        sys.exit(1)

    try:
        soup = BeautifulSoup(html_content, 'html.parser')
        return soup.prettify()
    except Exception as e:
        print(f"Error formatting HTML: {e}", file=sys.stderr)
        sys.exit(1)


def normalize_html(html_content, ignore_text_linebreaks=False):
    """
    Normalize HTML content by stripping indentation from each line.
    Removes leading/trailing whitespace while preserving meaningful content.
    Ignores empty lines and lines containing only whitespace.

    Args:
        html_content: The HTML content to normalize
        ignore_text_linebreaks: If True, joins line breaks within text-only tags
    """
    lines = html_content.split('\n')
    normalized_lines = []

    for line in lines:
        # Strip leading and trailing whitespace
        stripped = line.strip()
        # Only keep lines with actual content
        if stripped:
            normalized_lines.append(stripped)

    # Join lines that are split within text-only tags
    if ignore_text_linebreaks:
        normalized_lines = join_text_only_tags(normalized_lines)

    return normalized_lines


def join_text_only_tags(lines):
    """
    Join lines that are split within text-only HTML tags.
    A line that doesn't start with an opening tag '<' is considered a continuation and is joined with the previous line.
    Also joins closing tags that appear on their own line.
    Preserves spaces when joining to maintain text integrity.

    Characters that should not have a space added when joining (add more as needed):
    - / : For fractions like "230/300"

    For example:
        <div>230
        /300</div>
    becomes:
        <div>230/300</div>

    Also handles:
        <div>text
        </div>
    becomes:
        <div>text</div>
    """
    # Characters that should not have a space added when joining
    NO_SPACE_CHARS = {'/', ',', '.', '!', '?', ':', ';', ')', ']', '}'}

    result = []
    for line in lines:
        if result and (not line.startswith('<') or line.startswith('</')):
            # This line doesn't start with an opening tag, or is a closing tag - join it with previous line
            # Add a space between lines unless:
            # - It's a closing tag, or
            # - Previous line ends with a space, or
            # - Current line starts with a space, or
            # - Current line starts with a no-space character
            if (line.startswith('</') or result[-1].endswith(' ') or
                line.startswith(' ') or line[0] in NO_SPACE_CHARS):
                result[-1] = result[-1] + line
            else:
                result[-1] = result[-1] + ' ' + line
        else:
            result.append(line)
    return result


def compare_html(url1, url2, ignore_text_linebreaks=False, format_html_flag=False):
    """Fetch and compare HTML from two URLs."""
    print(f"Fetching {url1}...", file=sys.stderr)
    html1 = fetch_html(url1)

    print(f"Fetching {url2}...", file=sys.stderr)
    html2 = fetch_html(url2)

    # Format HTML if requested
    if format_html_flag:
        print("Formatting HTML from both URLs...", file=sys.stderr)
        html1 = format_html(html1)
        html2 = format_html(html2)

    # Normalize both HTML contents
    lines1 = normalize_html(html1, ignore_text_linebreaks)
    lines2 = normalize_html(html2, ignore_text_linebreaks)

    # Compare using unified_diff
    diff = list(unified_diff(
        lines1,
        lines2,
        fromfile=url1,
        tofile=url2,
        lineterm='',
        n=1  # Show 1 line of context
    ))

    if diff:
        print("\n=== Differences Found ===\n")
        for line in diff:
            print(line)
        print(f"\n=== Total difference lines: {len(diff)} ===")
        return False
    else:
        print("\n✓ HTML content matches (ignoring indentation)")
        return True


def main():
    import argparse

    parser = argparse.ArgumentParser(
        description="Validate HTML content matches between two URLs, ignoring indentation differences."
    )
    parser.add_argument("url1", help="First URL to compare")
    parser.add_argument("url2", help="Second URL to compare")
    parser.add_argument(
        "--ignore-text-linebreaks",
        action="store_true",
        help="Ignore line breaks within HTML tags containing only text (no nested HTML)"
    )
    parser.add_argument(
        "--format",
        action="store_true",
        help="Format HTML from both URLs using BeautifulSoup before comparison (requires beautifulsoup4)"
    )

    args = parser.parse_args()

    # Validate URLs
    for url in [args.url1, args.url2]:
        try:
            result = urllib.parse.urlparse(url)
            if not all([result.scheme, result.netloc]):
                print(f"Invalid URL: {url}", file=sys.stderr)
                sys.exit(1)
        except Exception as e:
            print(f"Error parsing URL {url}: {e}", file=sys.stderr)
            sys.exit(1)

    success = compare_html(args.url1, args.url2, args.ignore_text_linebreaks, args.format)
    sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()

