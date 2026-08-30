# Task-01: Git Exercises

## Objective

The objective of this task was to strengthen my understanding of Git by completing a series of practical exercises involving commits, staging, branches, merging, rebasing, conflict resolution, and Git history manipulation.

## What I Did

I completed all 23 Git exercises provided by the Git Exercises platform. Each exercise focused on a different Git concept and required solving a practical problem using Git commands.


### Exercises Completed

1. **master** — Initialized the exercise and verified the setup.
## " master "
### Purpose
The purpose of this exercise was to initialize the Git exercise repository and verify the initial commit created by the `git start` command.
### Commands Used
git start master
Purpose: --Initializes the master exercise and creates the initial Git setup/commit required by the exercise.

git verify
Purpose: --Checks whether the current state of the repository satisfies the requirements of the exercise.

Result
The exercise was successfully verified.

2. **commit-one-file** — Committed only one selected file.
### Commands Used
git add A.txt
git commit -m "Commit A.txt file"
it add A.txt
-- “Git, prepare A.txt for the next commit.”

git commit -m "Commit A.txt file"
-- “Git, save the prepared changes as a commit with this description.”


3. **commit-one-file-staged** — Unstaged one file and committed the required file.
###Commands used :
git reset A.txt
Purpose: Removes A.txt from the staging area so that it will not be included in the next commit.

git commit -m "Commit B.txt file"
Purpose: Commits the changes that are currently staged, which in this case includes B.txt but not A.txt.


4. **ignore-them** — Used `.gitignore` to ignore unwanted files
### Commands Used
git start ignore-them
Purpose: Starts the ignore-them exercise and creates the required starting files.

ls
Purpose: Shows the files present in the current folder so I can see which files need to be ignored.

nano .gitignore
Purpose: Opens the .gitignore file so I can enter the names of the files that I don't want Git to track.

*.log
Purpose: Tells Git to ignore all files ending with .log.

git status
Purpose: Checks the Git status and confirms that the ignored files are no longer shown as untracked files.

git add .gitignore
Purpose: Stages the .gitignore file so that the ignore rules can be included in the commit.

git commit -m "Ignore unwanted files"
Purpose: Saves the .gitignore changes as a Git commit.

git verify
Purpose: Checks whether the ignore-them exercise has been completed correctly.

Result
The required unwanted files were successfully ignored using .gitignore, and the exercise was successfully verified.


5. **chase-branch** — Worked with branches and merging.
###Commands used 
git merge escaped
Purpose: Merges the changes from the escaped branch into the current branch to bring the required changes together and complete the exercise.


6. **merge-conflict** — Resolved a merge conflict.
###Commands Used :
git echo 2+3=5 > equation.txt
Purpose: Creates a file named equation.txt and writes 2+3=5 into it.

git add equation.txt
Purpose: Stages equation.txt so it can be included in the next commit.

git commit --no-edit
Purpose: Creates the commit using the existing commit message, so there is no need to type a new message.


7. **save-your-work** — Used Git stash to temporarily save changes.

### Commands used
git start save-your-work
Purpose: Start the save-your-work exercise.

git status
Purpose: Check the current changes.

git stash
Purpose: Save my unfinished work temporarily.

cat bug.txt
Purpose: Check the bug.txt file and find the bug.

nano bug.txt
Purpose: Open the file and remove the bug line.

git diff
Purpose: Check the changes I made.

git add bug.txt
Purpose: Add the bug fix to staging.

git commit -m "Fix bug"
Purpose: Commit the bug fix.

git stash pop
Purpose: Bring my saved work back.

git status
Purpose: Check my work after getting it back.

echo "Finally, finished it!" >> bug.txt
Purpose: Add the required final line to bug.txt.

git add bug.txt program.txt
Purpose: Add my finished files to staging.

git commit -m "Finish work"
Purpose: Commit my finished work.

git verify
Purpose: Check whether the exercise is completed.


8. **change-branch-history** — Used rebase to modify branch history.
###Commands used 
git start change-branch-history
Purpose: Start the change-branch-history exercise.

git status
Purpose: Check the current branch.

git rebase hot-bugfix
Purpose: Put my work after the bug-fix commit.

git verify
Purpose: Check whether the exercise is completed.

9. **remove-ignored** — Removed an already-tracked file from Git tracking.

###Commands Used 
git start remove-ignored
Purpose: Start the remove-ignored exercise.

git status
Purpose: Check the repository status.

cat ignored.txt
Purpose: Check the ignored file.

git rm --cached ignored.txt
Purpose: Stop tracking the file but keep it on my computer.

git status
Purpose: Check the change.

git commit -m "Stop tracking ignored file"
Purpose: Commit the change.

git verify
Purpose: Check whether the exercise is completed.


10. **case-sensitive-filename** — Corrected a filename case issue.
###Commands Used 
git start case-sensitive-filename
Purpose: Start the case-sensitive-filename exercise.

git status
Purpose: Check the repository status.

git mv File.txt file.txt
Purpose: Rename File.txt to file.txt.

git status
Purpose: Check that Git detected the rename.

git commit -m "Rename file to lowercase"
Purpose: Commit the file rename.


11. **fix-typo** — Corrected a mistake using commit amendment.
###Commands Used 
git start fix-typo
Purpose: Start the fix-typo exercise.

nano file.txt
Purpose: Open the file and change wordl to world.

git add file.txt
Purpose: Add the corrected file to staging.

git commit --amend
Purpose: Change the last commit instead of creating a new commit.

In nano, I changed the commit message to:
Add Hello world
Purpose: Correct the typo in the commit message.

git verify
Purpose: Check whether the exercise is completed.


12. **forge-date** — Modified the date of a commit.
###Commands Used 
git commit --amend --no-edit --date="1987-08-03"
-- Changes the date of my last commit without changing its message

13. **fix-old-typo** — Corrected an older commit using interactive rebase.
###Commands Used 
git rebase -i HEAD2
-- Opens the last two commits so I can edit the first commit

-- Mark the first commit with "edit"
-- Stops at that commit so I can fix it

-- Fix the typo in the file
-- Corrects the typo in the file

git add file.txt
-- Adds the corrected file

git rebase --continue
-- Continues the rebase after fixing the file

-- Fix the typo in the commit message
-- Corrects the commit message when Git opens it

14. **commit-lost** — Recovered a lost commit using reflog.
###Commands Used 
git reflog
-- Shows my previous Git actions and commit positions

git reset --hard HEAD@{1}
-- Moves my branch back to the previous commit and removes the changes


15. **split-commit** — Split changes into separate commits.
###Commands Used 
git reset HEAD^
-- Undoes the last commit but keeps the files

git add first.txt
-- Adds first.txt to staging

git commit -m "First.txt"
-- Saves first.txt as a separate commit

git add second.txt
-- Adds second.txt to staging

git commit -m "Second.txt"
-- Saves second.txt as a separate commit

16. **too-many-commits** — Combined commits using interactive rebase.
###Commands Used 
git rebase -i HEAD^^
-- Opens the last two commits so I can combine them

-- squash or fixup the second commit
-- Combines the second commit with the first commit



17. **executable** — Changed file permissions to make a script executable.
###Commands Used 
git update-index --chmod=+x script.sh
-- Gives execute permission to script.sh in Git
-- Tells Git to treat the file as an executable script
-- Useful when the script runs but Git is not detecting the execute permission
-- Updates the file permission in Git without changing its contents

18. **commit-parts** — Staged selected parts of a file using patch mode.
###Commands Used 
git add -p file.txt
-- Shows the changes in file.txt one by one
-- I choose which changes I want to add to the first commit
-- I press "y" for the changes I want to include
-- The selected changes go to staging

git commit -m "First part of changes"
-- Saves the selected changes as the first commit

git commit -am "The rest of the changed"
-- Adds the remaining modified tracked changes and saves them as the second commit


19. **pick-your-features** — Used cherry-pick and squash merge to combine selected features.
###Commands Used 
git start pick-your-features
-- Starts the exercise

git cherry-pick feature-a
-- Picks Feature A commit into my branch

git cherry-pick feature-b
-- Picks Feature B commit into my branch

git cherry-pick feature-c
-- Picks Feature C commit and shows conflict

nano program.txt
-- Opens program.txt to fix the conflict

git add program.txt
-- Marks the conflict as fixed

git cherry-pick --continue
-- Continues the cherry-pick after fixing the conflict

git verify
-- Checks whether the exercise is completed


20. **rebase-complex** — Performed a complex rebase using `--onto`.
###Commands Used 
git start rebase-complex
-- Starts the exercise

git rebase --onto your-master E
-- Rebases only H and I commits onto D


21. **invalid-order** — Reordered commits using interactive rebase.
###Commands Used 
git start invalid-order
-- Starts the exercise

git log -2
-- Shows the last two commits

git rebase -i HEAD~2
-- Opens the last two commits so I can change their order

-- Change the two commit lines in nano
-- Puts the commits in the required order

git verify
-- Checks whether the exercise is completed



22. **find-swearwords** — Located and removed unwanted words from previous commits.

Issues faced -- while doing couldn't find whether the word was in in list.txt or words.txt 
spent so much time figuring out that , but finally could complete it

###Commands Used 
git start find-swearwords
-- Starts the exercise

git log -S "shit" -- words.txt list.txt
--Finds commits that introduced the word shit

git rebase -i 59a1280^
--Opens the commits for interactive rebase
--Change pick to edit for the required commits
--Stops at those commits so I can change the word

grep -n "shit" words.txt list.txt
-- Finds the line containing shit

nano words.txt
-- Opens the file to change shit to flower

nano list.txt
-- Opens the file to change shit to flower

git add .
-- Marks the changed file as resolved

git commit --amend --no-edit
--Updates the current commit without changing its message

git rebase --continue
--Continues the rebase

git rebase --abort
-- Cancels the rebase if I need to restart

git verify
-- Checks whether the exercise is completed


23. **find-bug** — Used Git history and bisect techniques to locate a problematic commit.
###Commands Used 
git bisect start
-- Starts Git bisect to find the commit where the problem started

git bisect bad HEAD
-- Marks the current commit as bad because it has the problem

git bisect good 1.0
-- Marks version 1.0 as good because the problem was not present there

git bisect run sh -c "openssl enc -base64 -A -d < home-screen-text.txt | grep -v jackass"
-- Automatically tests each commit and finds the first commit where jackass was introduced

## Result

All 23 exercises were successfully completed and verified.
