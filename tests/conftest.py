"""
Global unit testing fixtures for 'tests'.
"""

import pathlib
import pytest

MOCK_FILES = {
    "alpha.extn": "Alpha note.\n",
    "bravo.extn": "Bravo note.\n",
    "charlie.trash": "Charlie note.\n",
    "config.toml": """
        ext_pattern = "*.*"
        hard_delete = false
        lower_names = true
        trim_trails = true
    """,
}


@pytest.fixture(scope="function")
def conf(tmp_path: pathlib.Path) -> str:
    """
    Return a temporary file named "config.toml" containing the MOCK_FILES entry.
    """

    dest = tmp_path.joinpath("config.toml")
    dest.write_text(MOCK_FILES["config.toml"], encoding="utf-8")
    return str(dest)


@pytest.fixture(scope="function")
def dire(tmp_path: pathlib.Path) -> str:
    """
    Return a temporary directory populated with all MOCK_FILES entries.
    """

    for base, body in MOCK_FILES.items():
        dest = tmp_path.joinpath(base)
        dest.write_text(body, encoding="utf-8")

    return str(tmp_path)


@pytest.fixture(scope="function")
def file(tmp_path: pathlib.Path) -> str:
    """
    Return a temporary file named "alpha.extn" containing the MOCK_FILES entry.
    """

    dest = tmp_path.joinpath("alpha.extn")
    dest.write_text(MOCK_FILES["alpha.extn"], encoding="utf-8")
    return str(dest)
