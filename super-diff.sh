# This file is provided free of charge and may be modified and used by anyone. While I appreciate any improvements through pull requests to https://github.com/skelouse/code-to-gpt, they are not required. I hold no liability for any issues arising from the use of this script.

if test "$#" = 0; then
    (
        added_files=()
        git ls-files --others --exclude-standard |
            while read -r i; do
                git add --intent-to-add "$i"
                added_files+=("$i")
            done
        git diff --no-color

        # Revert the --intent-to-add
        for i in "${added_files[@]}"; do
            git reset "$i"
        done
    )
else
    git diff "$@"
fi