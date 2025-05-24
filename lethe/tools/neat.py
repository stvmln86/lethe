"""
The 'lethe.tools.neat' module implements value sanitisation and conversion functions.
"""

import os.path
import pathlib
import string

NAME_CHARS = string.ascii_letters + string.digits + "-_"


def body(body: str) -> str:
    """
    Return a whitespace-trimmed file body string with a trailing newline.
    """

    return body.strip() + "\n"


def glob(glob: str) -> str:
    """
    Return a whitespace-trimmed lowercase glob pattern string.
    """

    return glob.strip().lower()


def name(name: str) -> str:
    """
    Return a whitespace-trimmed lowercase alphanumeric file name string.
    """

    name = "".join(char for char in name.replace(" ", "-") if char in NAME_CHARS)
    return name.strip("-").lower()


def path(path: str | pathlib.Path) -> str:
    """
    Return a whitespace-trimmed normalised file path string.
    """

    path = str(path).strip()
    return os.path.normpath(path)
