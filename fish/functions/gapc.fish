function gapc
    if test (count $argv) -eq 0
        echo "Usage: gapc <commit message>"
        return 1
    end
    git add .
    git commit -m "$argv"
    git push
end
