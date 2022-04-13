```
   _________       __                         __   ____        __           _____ __             __         
  / ____/ (_)___  / /_  ____  ____ __________/ /  / __ \____ _/ /_____ _   / ___// /____  ____ _/ /__  _____
 / /   / / / __ \/ __ \/ __ \/ __ `/ ___/ __  /  / / / / __ `/ __/ __ `/   \__ \/ __/ _ \/ __ `/ / _ \/ ___/
/ /___/ / / /_/ / /_/ / /_/ / /_/ / /  / /_/ /  / /_/ / /_/ / /_/ /_/ /   ___/ / /_/  __/ /_/ / /  __/ /    
\____/_/_/ .___/_.___/\____/\__,_/_/   \__,_/  /_____/\__,_/\__/\__,_/   /____/\__/\___/\__,_/_/\___/_/     
        /_/                                                                                                 
```                                                                  
Yeah.... I couldn't think of a cool name so I went with "Clipboard Data Stealer". It's short, sweet, says what it does.

### About This Program
This is a basic program that copies clipboard data using the OpenClipboard, GetClipboardData and CloseClipboard Windows APIs. It consistantly monitors for new text copied and sends a GET request to a web server using URLOpenBlockingStreamA Windows API.
There isn't all that much more to say than that other than I am **not** responsible for what you do with this program. It is categorized as malware and should not be used on any production system. Only in a lab environment. If you do bad things with this program, that's on you. I will not be providing a release to prevent abuse of this tool.
With that said, please use this program responsibily <3

<img src='https://raw.githubusercontent.com/Sq00ky/Clipboard-Data-Stealer/main/main.png'></img>

### Complication Instructions
This software can be compiled with Visual Studio Code 2019 as a 32-bit or 64-bit program. Before changing, ensure you set the IP Address or Domain Name in the URL string to a device you control or else data will not be logged. Who needs disk-based artifacts!? It's 2022!

### Known Issues
Currently, non-ascii data (ex. Pictures) will crash the program. I have no idea why. Maybe I'll fix it one day? I'm not exactly sure what's wrong :(
