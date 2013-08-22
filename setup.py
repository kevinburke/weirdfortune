#!/usr/bin/env python
from setuptools import setup, find_packages

setup(
    name = "weirdfortune",
    version = "0.1",
    description = "weirder fortune",
    author = "Kevin Burke",
    author_email = "kev@inburke.com",
    packages = find_packages(),
    install_requires = ['clint', 'tweepy']
)

