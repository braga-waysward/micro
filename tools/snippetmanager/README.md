# Snippet Manager Documentation

## Overview
A simple tool to manage code snippets for the micro editor. Store and reuse common code fragments easily.

## Installation
1. Ensure Go is installed
2. Run the setup script:
```bash
chmod +x setup.sh
./setup.sh
```

## Basic Commands
| Command | Description | Example |
|---------|-------------|---------|
| `list` | Show all snippets | `snippetmanager list` |
| `add` | Add new snippet | `snippetmanager add python-forloop` |
| `show` | View snippet content | `snippetmanager show python-forloop` |
| `delete` | Remove snippet | `snippetmanager delete python-forloop` |

## How to Use
1. **Adding a snippet**:
   ```bash
   snippetmanager add html-template
   ```
   Then paste your code. Press Enter twice to finish.

2. **Using snippets**:
   - View with `show` command
   - Copy and paste into your editor

## Storage
Snippets are saved in:
```
~/.config/micro/snippets.json
```

## Tips
- Use descriptive names (e.g., "python-flask-route")
- Organise with prefixes (e.g., "js-", "py-")
- The tool works with any programming language

## Uninstallation
```bash
sudo make uninstall
```

##
