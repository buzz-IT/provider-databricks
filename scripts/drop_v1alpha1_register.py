#!/usr/bin/env python3
"""Remove hardcoded apis/{cluster,namespaced}/v1alpha1 entries from zz_register.go."""

from __future__ import annotations

import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
FILES = [
    ROOT / "apis/cluster/zz_register.go",
    ROOT / "apis/namespaced/zz_register.go",
]


def main() -> None:
    import_re = re.compile(r'\tvd?1alpha1 "[^"]+/v1alpha1"\n|\tv1alpha1 "[^"]+/v1alpha1"\n')
    scheme_re = re.compile(r"\t\tv1alpha1\.SchemeBuilder\.AddToScheme,\n")
    for path in FILES:
        if not path.exists():
            continue
        text = path.read_text()
        text = import_re.sub("", text)
        text = scheme_re.sub("", text)
        path.write_text(text)


if __name__ == "__main__":
    main()
