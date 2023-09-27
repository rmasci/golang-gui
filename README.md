# golang-gui
Original Video for this is here:
https://www.youtube.com/watch?v=9bDN2rrf-Pw&t=790s

Original Repository for this code:
https://github.com/devhulk/golang-gui

Main item is to show the use of script. The commands are mainly pulled straight out of the original bash script menu.sh

github.com/bitfield/script

This is an awesome tool for running command line apps.

Not saying the original by 'devhulk' code was bad at all, it's fantastic. I just put my spin on it. :)
## Removed Functions:
### bytesToGB
replaced by go package humanize.

### Menu
moved this back to main. Utilized the function in each 'AddItem', this works a lot like a switch statement. When the user selects the menu item, it runs the function.

### setResponse
Instead of a function just for switch, the function of AddItem is used.

## New Function:
### errorHandle
I always put one of these in my apps. 

### Displaybox.
Take a text output from script and show it to the screen. Wait for a keypress and go back to the menu
