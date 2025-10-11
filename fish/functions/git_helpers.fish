

function gacp
    # git add, commit, and push (with commit message as argument)
    if test (count $argv) -eq 0
        echo "Usage: gacp <commit message>"
        return 1
    end
    git add .
    git commit -m "$argv"
    git push -u origin main
end

