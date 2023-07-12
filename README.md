# Gopsalm
![Icon Made by Dall e 2 just because](Images/gopher.png)
![Static Badge](https://img.shields.io/badge/Golang-made)
![Static Badge](https://img.shields.io/badge/Golang-made)
![Static Badge](https://img.shields.io/badge/Lipgloss-Charm-MIT)
![Static Badge](https://img.shields.io/badge/MIT-uwu)


Gopsalm is a Cli program made to display random text with lipgloss(charm group) 
everytime you write gopsalm (get some colour) or you open your terminal, it has a focus on
biblical verses but you can remove them or add new ones to make it just for you

![Example](Images/Example.png) 

# :warning: Important

Before installing go psalm make sure to install git, zsh and lipgloss

+ zsh:
https://github.com/ohmyzsh/ohmyzsh/wiki/Installing-ZSH

+ lipgloss:
https://github.com/charmbracelet/lipgloss

If you have p10k go to home and then to .p10k.zsh
and make sure to put typeset -g POWERLEVEL9K_INSTANT_PROMPT
to quiet

```
  typeset -g POWERLEVEL9K_INSTANT_PROMPT=quiet
```

# Installing Gopsalm

To install GoPsalm

+ git clone the repository

```
git clone https://github.com/HanamDavid/GoPsalm.git
```
+ cd into the GoPsalm
```
cd GoPsalm
```
+ make install.sh executable with chmod +x install.sh 

```
chmod +x install.sh
```

+ make any changes to the code that you want

+ Execute install.sh, put your password when it ask you for it  
(If you want to make sure everything is fine look at the code of install.sh) 
and enjoy it!
```
./install.sh
```
# Configuration

+ ** Colours and Display  ** 
to configure how the text is going to display 
go to gopsalm.go and to line 14 to 31 you can make any changes
to how the text is going to display(go to lipgloss git hub for more info)
on your shell, to update it  when you end the configuration re 
run install.sh and your shell 
+ ** Text **

If you want to add more text or remove some of them go to ~/.config/gopsalm/Display/ 
Here you can add any new phrase or versicle writing a new markdown file 
and saving it, alternatively you can go to line 42 in gopsalm.go and 
change the directory where gopsalm gets its text and then change as 
well the path on the installer install.sh

# License

# The MIT License (MIT)

Copyright © 2023 <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the “Software”),
to deal in the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of 
the Software, and to permit persons to whom the Software is furnished to do so, subject
to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED 
TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR 
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

