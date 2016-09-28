#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
import sqlite3


OUT_FILE_NAME = "db"


class SQLiteHeaderNotFoundException(Exception):
    pass


def seek_sqlite_header(data):
    HEADER = ['0x53', '0x51', '0x4c', '0x69', '0x74', '0x65', '0x20', '0x66', '0x6f', '0x72', '0x6d', '0x61', '0x74', '0x20', '0x33']

    n = 0
    at = 0
    for i in range(len(data)):
        c = data[i]
        if hex(c) == HEADER[n]:
            if n == 0:
                at = i
            n += 1
        elif n > 0:
            n = 0

        if n == len(HEADER) - 1:
            return at

    raise SQLiteHeaderNotFoundException()


def extract_db():
    IN_FILE_NAME = "sample.clip"

    with open(IN_FILE_NAME, "rb") as fi, open(OUT_FILE_NAME, "wb") as fo:
        data = fi.read()

        try:
            at = seek_sqlite_header(data)
        except SQLiteHeaderNotFoundException:
            print("[ERROR] SQLite header not found.")
            sys.exit(1)

        fo.write(bytes(data[at:]))


def extract_illust():
    ILLUST_FILE = "image.png"

    with sqlite3.connect(OUT_FILE_NAME) as conn:
        cursor = conn.cursor()
        cursor.execute("SELECT ImageData from CanvasPreview")

        with open(ILLUST_FILE, "wb") as f:
            f.write(cursor.fetchone()[0])
            print("OK")


def main():
    extract_db()
    extract_illust()

if __name__ == "__main__":
    main()
