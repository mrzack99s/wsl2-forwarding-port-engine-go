# wsl2-forwarding-port-engine-go

>	Author: Chatdanai Phakaket <br>
>	Email: zchatdanai@gmail.com 

WSL2-forwarding-port-engine-go is backend of WSL2-forwarding-port cli


## How to install

1. Open WSL2
2. Download the binary with the command 
```
    curl -LO https://github.com/mrzack99s/wsl2-forwarding-port-engine-go/releases/download/v1.0.1-beta/wfp-engine.exe
    curl -LO https://github.com/mrzack99s/wsl2-forwarding-port-engine-go/releases/download/v1.0.1-beta/wfp-engine-autorun.vbs
```
3. Make the wfp-engine binary executable.
```
    chmod +x wfp-engine.exe
```
4. Create directory
```
    mkdir /mnt/c/Users/<window-username>/.wfp-engine
```
5. Change window username in wfp-engine-autorun.vbs
6. Move the binary into PATH.
```
    mv ./wfp-engine.exe /mnt/c/Users/<window-username>/.wfp-engine
```
7. Press key Win+R and paste wfp-engine-autorun.vbs to this below
```
    shell:startup
```

Let's enjoy !!!!


## License

Copyright (c) 2021 - Chatdanai Phakaket

	

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)