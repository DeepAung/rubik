[![Go Report Card](https://goreportcard.com/badge/github.com/DeepAung/rubik)](https://goreportcard.com/report/github.com/DeepAung/rubik)

# rubik
A Rubik's simulator app in the terminal. using [bubbletea](https://github.com/charmbracelet/bubbletea) as the terminal framework.

![image](https://github.com/DeepAung/rubik/assets/87839907/b9cc01e5-7a8c-4209-aca5-6002e313591c)

## How to use
- install by running `go install github.com/DeepAung/rubik`.
- after that, run `rubik` command to load the app.

## Feature
- A help bar is on the bottom to help you use the keybinding.
- A Rubik in 2d view showing on the left side
- On the right side, you can execute the following action
  - Rotate the cube by entering the notations.
  - Reset the cube.
  - Find the number of times and moves of rotations that make the Rubik revert to its original state
  - Undo and Redo the rotation
