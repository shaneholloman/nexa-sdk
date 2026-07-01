# Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
# SPDX-License-Identifier: BSD-3-Clause

"""Appium session fixture for the on-device bench run."""

import os

import pytest
from appium import webdriver

from utils import options, write_qdc_log


@pytest.fixture(scope="session", autouse=True)
def driver():
    return webdriver.Remote(
        command_executor="http://127.0.0.1:4723/wd/hub", options=options
    )


def pytest_sessionfinish(session, exitstatus):
    xml = getattr(session.config.option, "xmlpath", None) or "results.xml"
    if os.path.exists(xml):
        with open(xml) as f:
            write_qdc_log("results.xml", f.read())
