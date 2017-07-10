# Git Lunch and Learn

[Official Tutorial](https://git-scm.com/docs/gittutorial)

 -	Overview of Git approach as compared to TFS
    - TFS is centralized (SQL Server, TFS Server, Collections, Projects)
    - Easy merges / No "Abandon Ship Procedure"
    - Git is decentralized

      ![](./git-flow.png)
 -	Behind the scenes: What’s happening??
    - Commits are just diffs
      ![](./stages.png)
 -	Tools
    -	Commandline
    - SourceTree
    -	VS Code
    -	Visual Studio
 -	Tips and Tricks: How do I fix it when I … ?
    - First things first:
    ```
    git config --global user.name "Dan Wolf"
    git config --global user.email dcwolf@gmail.com
    git config --global credential.helper wincred
    ```
    - `.gitignore` and `.gitkeep`
    - Remove ignored files from cache (another term for stage or index)
    ```
    foreach ($i in iex 'git ls-files -i --exclude-from=.gitignore') { git rm --cached $i }
    ```
    - Revert (before push)
    ```
    git reset --hard <SOME-COMMIT, e.g. ~HEAD>
    ```
    - Or, if you only want to add changes to previous commit, you could do a `git reset --soft HEAD~` which will undo the previous commit but leave your changes staged.

    - If you already pushed, use `git revert HEAD` to create an "undo" commit. This will prevent other repos from getting mixed up. **HEAD** is simply the easy (non-SHA) name for your last commit.

    - Commit your code *yesterday* if you forgot to commit and want the correct date reflected
    ```
    git commit --date="`date --date='n day ago'`" -am "Commit Message"
    ```

 - **When forking a Go project** you'll be faced with broken import paths if you use go get and suddenly have a new path to your repo. You have two choices.

    - Create the folder path of `GOPATH\src\user\` manually and `git clone https:\\path.git` into there (and go get `./...` in there will still grab your deps) or...

    - Use a named remote that points to your fork. You'll need to use the name in pull / push operations after that:
    ```
    go get http://github.com/other-user/repo
    git remote add my-fork http://github.com/my-user/repo

    git pull --rebase my-fork
    git push my-fork
    ```