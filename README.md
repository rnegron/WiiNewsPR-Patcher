# WiiNewsPR-Patcher

A utility to patch Wii News Channel WAD files to redirect news requests from Nintendo's original servers to a custom server (http://wii.rauln.com).

## Usage

```bash
wiinewspr-patcher <path_to_news_wad_file> <output_path>
```

### Example
```bash
wiinewspr-patcher news.wad patched_news.wad
```

## Installation

### Using go install
```bash
go install github.com/rnegron/wiinewspr-patcher@latest
```
After installation, the `wiinewspr-patcher` command should be available.

## Requirements

- Go 1.25.0 or higher (no particular reason, just the version I had installed)
- Wii News Channel WAD file (¯\\_(ツ)_/¯):
```
News Channel (USA) (v7) (Channel).wad
MD5:	7CCB0D36C06BC627ADCE8A0687279940
SHA-1:	B0E32A6C7E2C216CACCFB031667892549CB765CA
```

## Credits

This small utility is built on [wadlib](https://github.com/wii-tools/wadlib).
