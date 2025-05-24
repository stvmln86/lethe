"""
Tests for 'lethe.tools.neat'.
"""

from lethe.tools import neat


def test_body():
    # success
    body = neat.body("\tBody.\n")
    assert body == "Body.\n"


def test_glob():
    # success
    glob = neat.glob("\t*.EXTN\n")
    assert glob == "*.extn"


def test_name():
    # success
    name = neat.name("\tNAME 123 !!!\n")
    assert name == "name-123"


def test_path():
    # success
    path = neat.path("\t/././path\n")
    assert path == "/path"
