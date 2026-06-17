# Copyright 2024-2026 Qualcomm Technologies, Inc. and/or its subsidiaries.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""``geniex.get_compute_unit_list`` shape per runtime."""

from __future__ import annotations

import geniex


def test_compute_unit_list_shape_for_each_runtime(geniex_session):
    for runtime in geniex.get_runtime_list():
        compute_units = geniex.get_compute_unit_list(runtime)
        assert isinstance(compute_units, list)
        for entry in compute_units:
            assert isinstance(entry, tuple) and len(entry) == 2
            compute_unit, label = entry
            assert isinstance(compute_unit, str) and compute_unit
            assert isinstance(label, str)
