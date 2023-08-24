### Notes taken while studying the C programming language.

#### Available modes for opening files in both binary and text modes using the fopen function in C

```text
Text Mode:
    "r": Read mode. Opens the file for reading.
    "w": Write mode. Opens the file for writing. Creates a new file or truncates an existing file to zero length.
    "a": Append mode. Opens the file for writing. Data is added to the end of the file. If the file does not exist, it will be created.
    "r+": Read and write mode. Opens the file for both reading and writing.
    "w+": Read and write mode. Opens the file for reading and writing. Creates a new file or truncates an existing file to zero length.
    "a+": Read and write mode. Opens the file for reading and writing. Data is added to the end of the file. If the file does not exist, it will be created.

Binary Mode:
    "rb": Read mode in binary. Opens the file for reading in binary mode.
    "wb": Write mode in binary. Opens the file for writing in binary mode. Creates a new file or truncates an existing file to zero length.
    "ab": Append mode in binary. Opens the file for writing in binary mode. Data is added to the end of the file. If the file does not exist, it will be created.
    "r+b": Read and write mode in binary. Opens the file for both reading and writing in binary mode.
    "w+b": Read and write mode in binary. Opens the file for reading and writing in binary mode. Creates a new file or truncates an existing file to zero length.
    "a+b": Read and write mode in binary. Opens the file for reading and writing in binary mode. Data is added to the end of the file. If the file does not exist, it will be created.
```
