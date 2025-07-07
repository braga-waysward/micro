#!/bin/bash
# Setup script for Micro Snippet Manager
# Installs the tool and sets up shell completion

# Check for required commands
for cmd in go make; do
    if ! command -v $cmd &> /dev/null; then
        echo "Error: $cmd is not installed. Please install it first."
        exit 1
    fi
done

# Build and install
echo "=== Building and installing Micro Snippet Manager ==="
make build
sudo make install

# Setup bash completion
echo "=== Setting up bash completion ==="
COMPLETION_SCRIPT="/etc/bash_completion.d/snippetmanager"

sudo bash -c "cat > $COMPLETION_SCRIPT << 'EOF'
#! /bin/bash
_snippetmanager_completion() {
    local cur prev opts
    COMPREPLY=()
    cur=\"\${COMP_WORDS[COMP_CWORD]}\"
    prev=\"\${COMP_WORDS[COMP_CWORD-1]}\"
    opts=\"list add show delete\"

    if [[ \${COMP_CWORD} -eq 1 ]] ; then
        COMPREPLY=( \$(compgen -W \"\${opts}\" -- \${cur}) )
        return 0
    elif [[ \${COMP_CWORD} -eq 2 && \${prev} == 'show' ]] || \
         [[ \${COMP_CWORD} -eq 2 && \${prev} == 'delete' ]]; then
        COMPREPLY=( \$(compgen -W \"\$(snippetmanager list | grep '^- ' | sed 's/^- //')\" -- \${cur}) )
        return 0
    fi
}
complete -F _snippetmanager_completion snippetmanager
EOF"

echo "Bash completion installed to $COMPLETION_SCRIPT"
echo "You may need to restart your shell or run 'source ~/.bashrc'"

echo "=== Installation complete ==="
echo "Usage:"
echo "  snippetmanager list                 # List all snippets"
echo "  snippetmanager add <name>           # Add snippet"
echo "  snippetmanager show <name>          # Show snippet content"
echo "  snippetmanager delete <name>        # Delete snippet"