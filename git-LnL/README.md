# Git Lunch and Learn

[Official Tutorial](https://git-scm.com/docs/gittutorial)

## Introduction

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
    - If you forget the above before pushing your initial commit, do a quick ` git rebase` and `git commit --amend --author="Dan Wolf <dcwolf@gmail.com>" -m "initial commit"`. You'll need to force the push with `git push -u origin master -f`
    - `.gitignore` and `.gitkeep`
    - Remove ignored files from cache (another term for stage or index)
    ```
    foreach ($i in iex 'git ls-files -i --exclude-from=.gitignore') { git rm --cached $i }
    ```
    - If you get the following error:
    ```
    remote: error: File astilectron-eg/5.single_binary_distribution/vendor.go is 199.38 MB; this exceeds GitHub's file size limit of 100.00 MB
    To https://github.com/PaluMacil/misc.git
    ! [remote rejected] master -> master (pre-receive hook declined)
    ```
    ...you will need to remove this file from history because it was too large for Github (100 MB max). If you want apply these changes to ALL branches, you need to use a `-- --all` flag instead of `HEAD`. This is only for when 100% necessary. If other people have somehow pulled your code, they will lose continuity. Make sure the filepath (or dir) is accurate. Note that in my example, I used the path as shown from the error. This command is run from the repository root folder.
    ```
    git filter-branch --index-filter 'git rm -r --cached --ignore-unmatch vendor.go' -f HEAD
    ```
    - Reset (before push) to to set the files back to the state they were in a previous commit (for example, the parent of HEAD). 
    ```
    git reset --hard <SOME-COMMIT, e.g. ~HEAD>
    ```
    - Or, if you only want to add changes to previous commit, you could do a `git reset --soft HEAD~` which will undo the previous commit but leave your changes staged. In other words, your files won't be touched.
## Keep things Organized

 -  File Hygiene
    - Tagging:  `git tag -a v0.2 -m "migration from Bottle to Flask"` This annotates a point in history. As a bonus, the [semantic version](http://semver.org) in the annotation will cause Github to make this a release. If you need more control, consider [creating the release manually](https://help.github.com/articles/creating-releases) on Github. Finally, don't forget to `git push --tags` since they don't get pushed to a remote without explicitly doing so.

## Tips and Tricks: How do I fix it when I … ?
- First things first:
  - Set up how
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
- If you get the following error:
```
remote: error: File astilectron-eg/5.single_binary_distribution/vendor.go is 199.38 MB; this exceeds GitHub's file size limit of 100.00 MB
To https://github.com/PaluMacil/misc.git
! [remote rejected] master -> master (pre-receive hook declined)
```
...you will need to remove this file from history because it was too large for Github (100 MB max). If you want apply these changes to ALL branches, you need to use a `-- --all` flag instead of `HEAD`. This is only for when 100% necessary. If other people have somehow pulled your code, they will lose continuity. Make sure the filepath (or dir) is accurate. Note that in my example, I used the path as shown from the error. This command is run from the repository root folder.
```
git filter-branch --index-filter 'git rm -r --cached --ignore-unmatch vendor.go' -f HEAD
```
- Reset (before push) to to set the files back to the state they were in a previous commit (for example, the parent of HEAD). 
```
git reset --hard <SOME-COMMIT, e.g. ~HEAD>
```
- Or, if you only want to add changes to previous commit, you could do a `git reset --soft HEAD~` which will undo the previous commit but leave your changes staged. In other words, your files won't be touched.

- If you already pushed, use `git revert HEAD` to create an "undo" commit. This will prevent other repos from getting mixed up. **HEAD** is simply the easy (non-SHA) name for your last commit.

- Commit your code *yesterday* if you forgot to commit and want the correct date reflected
```
git commit --date="`date --date='n day ago'`" -am "Commit Message"
```

 - **When forking a Go project** you'll be faced with broken import paths if you use go get and suddenly have a new path to your repo. You have two choices.

    - Create the folder path of `GOPATH\src\user\` manually and `git clone https:\\path.git` into there (and `go get ./...` in there will still grab your deps) or...

    - Use a named remote that points to your fork. You'll need to use the name in pull / push operations after that:
    ```
    go get http://github.com/other-user/repo
    git remote add my-fork http://github.com/my-user/repo

    git pull --rebase my-fork
    git push my-fork
    ```

## Referencing Things

 - **^ and /~** The caret (^) specifies which parent, while tilde (&#126) moves back generations.

   - If you don't know the order of parents, you can see this through `git log` or `git show`. The order shown here is the order to use in this syntax.

   - "HEAD^ means the first parent of the tip of the current branch. Remember that git commits can have more than one parent. HEAD^ is short for HEAD^1, and you can also address HEAD^2 and so on as appropriate. You can get to parents of any commit, not just HEAD. You can also move back through generations: for example, master/~2 means the grandparent of the tip of the master branch, favoring the first parent in cases of ambiguity. These specifiers can be chained arbitrarily , e.g., topic/~3^2." (credit: [Greg Bacon](https://stackoverflow.com/users/123109/greg-bacon))

   ![](./git-tilde-hat.png)

   [credit: Alex Janzik](https://stackoverflow.com/users/22038/alex-janzik)

[&#x2190; Back to Project List](../README.md)
