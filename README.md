# Clip
[![CircleCI](https://circleci.com/gh/lycoris0731/clip.svg?style=svg&circle-token=0e33c0cfb7bb1105ff821abbe845483d269145f8)](https://circleci.com/gh/lycoris0731/clip)
A tracking helper for CLIP STUDIO PAINT files with Git  
![Demo](./res/out.gif)  

## Description  
You can track to CLIP STUDIO PAINT files with Git by using this.  

## Equipments
- Go v1.7.1 or later

## Installation
``` sh
$ go get github.com/lycoris0731/clip
```

## Recommendation
You **should** use Git LFS(Large File Storage) to track clip files.  
Because Git managing binary files by whole save files of each commits.  

## Usage
First, you should run below command in Git repository.
``` sh
$ clip init TARGET_FILE_NAME
```
Then, clip creates `.clip` directory and update `post-commit` in Git hooks.  
All images are saved to `.clip`.  
  
See the image at the time of a commit.
``` sh
$ clip show [COMMIT_HASH ...]
```
Also, you can use `HEAD`.  
``` sh
$ clip show HEAD~
```

Create a Gif image about the production process.  
``` sh
$ clip gif [command options]
```

## License
Please see [LICENSE](./LICENSE).
