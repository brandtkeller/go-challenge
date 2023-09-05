# Go Challenge

`Go Challenge` is a Command Line utility for creating quick "challenge" directories to aid those on the quest to improve their software acumen. This project will quickly create the bare-minimum files required to begin solving coding challenges in [Golang](https://go.dev/).

## Note
This is an experiment - I wanted to see if I could improve general workflow for solving coding challenges in a way that removes local development `fluff` as well as enables challengers to utilize TDD.

## Expected Use
I personally pair this with platforms such as [LeetCode](leetcode.com) to serve as a helper utility to aid in solving the problems.

"Why" you might ask? for a few reasons in particular?
- Test Driven Development
    - Before starting on the challenge at hand, Writing tests furthers the understanding of what nuances or exceptions the challengee might be required to face.
- Development Environment
    - Some might argue that being too comfortable with your IDE is a crutch - and they're not wrong. But there is also an argument towards efficiency and using the tools - that you use everyday - to solve challenges just the same as business logic.

## Required Dependencies
- Golang

## Getting Started

Clone repo and Build the CLI
```
git clone https://github.com/brandtkeller/go-challenge.git
cd go-challenge
go build .
```

Create a new project
```
./go-challenge add my-new-challenge
```

*note: this will create a myNewChallenge function - it is best to name the challenge after the intended function to be solved* 
